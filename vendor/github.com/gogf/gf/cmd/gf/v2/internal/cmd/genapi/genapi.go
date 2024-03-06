// Copyright GoFrame gf Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package genapi

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"

	"github.com/gogf/gf/cmd/gf/v2/internal/utility/mlog"
)


func (c CGenDao) Api(ctx context.Context, in CGenDaoInput) (out *CGenDaoOutput, err error) {
	in.generatedFilePaths = &CGenDaoInternalGeneratedFilePaths{
		DaoFilePaths:         make([]string, 0),
		DaoInternalFilePaths: make([]string, 0),
		DoFilePaths:          make([]string, 0),
		EntityFilePaths:      make([]string, 0),
		ApiFilePaths:         make([]string, 0),
	}
	if g.Cfg().Available(ctx) {
		v := g.Cfg().MustGet(ctx, CGenDaoConfig)
		if v.IsSlice() {
			for i := 0; i < len(v.Interfaces()); i++ {
				doGenApiForArray(ctx, i, in)
			}
		} else {
			doGenApiForArray(ctx, -1, in)
		}
	} else {
		doGenApiForArray(ctx, -1, in)
	}
	mlog.Print("done!")
	return
}


// doGenApiForArray implements the "gen dao" command for configuration array.
func doGenApiForArray(ctx context.Context, index int, in CGenDaoInput) {
	var (
		err error
		db  gdb.DB
	)
	if index >= 0 {
		err = g.Cfg().MustGet(
			ctx,
			fmt.Sprintf(`%s.%d`, CGenDaoConfig, index),
		).Scan(&in)
		if err != nil {
			mlog.Fatalf(`invalid configuration of "%s": %+v`, CGenDaoConfig, err)
		}
	}
	if dirRealPath := gfile.RealPath(in.Path); dirRealPath == "" {
		mlog.Fatalf(`path "%s" does not exist`, in.Path)
	}
	removePrefixArray := gstr.SplitAndTrim(in.RemovePrefix, ",")

	// It uses user passed database configuration.
	if in.Link != "" {
		var tempGroup = gtime.TimestampNanoStr()
		gdb.AddConfigNode(tempGroup, gdb.ConfigNode{
			Link: in.Link,
		})
		if db, err = gdb.Instance(tempGroup); err != nil {
			mlog.Fatalf(`database initialization failed: %+v`, err)
		}
	} else {
		db = g.DB(in.Group)
	}
	if db == nil {
		mlog.Fatal(`database initialization failed, may be invalid database configuration`)
	}

	var tableNames []string
	if in.Tables != "" {
		tableNames = gstr.SplitAndTrim(in.Tables, ",")
	} else {
		tableNames, err = db.Tables(context.TODO())
		if err != nil {
			mlog.Fatalf("fetching tables failed: %+v", err)
		}
	}
	// Table excluding.
	if in.TablesEx != "" {
		array := garray.NewStrArrayFrom(tableNames)
		for _, v := range gstr.SplitAndTrim(in.TablesEx, ",") {
			array.RemoveValue(v)
		}
		tableNames = array.Slice()
	}

	// merge default typeMapping to input typeMapping.
	if in.TypeMapping == nil {
		in.TypeMapping = defaultTypeMapping
	} else {
		for key, typeMapping := range defaultTypeMapping {
			if _, ok := in.TypeMapping[key]; !ok {
				in.TypeMapping[key] = typeMapping
			}
		}
	}

	// Generating dao & model go files one by one according to given table name.
	newTableNames := make([]string, len(tableNames))
	for i, tableName := range tableNames {
		newTableName := tableName
		for _, v := range removePrefixArray {
			newTableName = gstr.TrimLeftStr(newTableName, v, 1)
		}
		newTableName = in.Prefix + newTableName
		newTableNames[i] = newTableName
	}
	// Api.
	generateApi(ctx, CGenDaoInternalInput{
		CGenDaoInput:  in,
		DB:            db,
		TableNames:    tableNames,
		NewTableNames: newTableNames,
	})

	if in.Clear {
		doClear(ctx, in)
	}
}