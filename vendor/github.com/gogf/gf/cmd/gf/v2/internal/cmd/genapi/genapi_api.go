// Copyright GoFrame gf Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package genapi

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gview"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"

	"github.com/gogf/gf/cmd/gf/v2/internal/utility/mlog"
	"github.com/gogf/gf/cmd/gf/v2/internal/utility/utils"
)

// 需要的数据:
// - 项目对象
// -   项目名
// - 表对象
// -   首字母大写表名
// -   小写表名
// -   表备注
// - 字段对象
// -   首字母大写字段名
// -   字段 golang 类型
// -   字段 json 名(小写字段名)
// -   字段备注
type ApiData struct {
	ProjName     string               // 项目名，英文 import 使用路径
	TableName    string               // 表名，引文 首字母大写
	TableComment string               // 表备注
	Fields       map[string]*ApiField // 字段对象
}
type ApiField struct {
	gdb.TableField
	GType   string
	JsonTag string
}

func generateApi(ctx context.Context, in CGenDaoInternalInput) {
	for i, tableName := range in.TableNames {
		newTableName := in.NewTableNames[i]
		schemaRes, err := in.DB.GetValue(ctx, fmt.Sprintf(`SELECT sql FROM sqlite_master WHERE type='table' AND name='%s';`, tableName))
		if err != nil {
			mlog.Fatalf("fetching tables schemaRes failed for table '%s':\n%v", tableName, err)
		}
		_, tableComment, _ := getTableComments(ctx, schemaRes.String())
		data := ApiData{
			ProjName:     getProjName(),
			TableName:    tableName,
			TableComment: gstr.Trim(tableComment),
		}
		fieldMap, err := in.DB.TableFields(ctx, tableName)
		if err != nil {
			mlog.Fatalf("fetching tables fields failed for table '%s':\n%v", tableName, err)
		}

		data.Fields = getApiFields(ctx, fieldMap, generateStructDefinitionInput{
			CGenDaoInternalInput: in,
			TableName:            tableName,
			StructName:           gstr.CaseCamel(newTableName),
			FieldMap:             fieldMap,
			IsDo:                 false,
		})
		entityFilePath := filepath.FromSlash(gfile.Join(in.Path, in.ApiPath, gstr.CaseSnake(newTableName)+".go"))
		tpl := NewTpl()
		tpl.Gv.BindFuncMap(gview.FuncMap{
			"ToSnake":       gstr.CaseSnake,
			"ToCamel":       gstr.CaseCamel,
			"ToPadAndQuote": strToEqLen,
		})
		if err := tpl.Write(entityFilePath, "gen_templates/api.go.tpl", g.Map{
			"data": data,
		}); err != nil {
			mlog.Fatalf("writing content to '%s' failed: %v", entityFilePath, err)
		} else {
			mlog.Print(ctx, "--------------------------------")
			utils.GoFmt(entityFilePath)
			mlog.Printf("generated: %s", entityFilePath)
		}
	}
}

// 保持返回的字符串长度始终是 20，如果长度不够，就在后面加 空格
func strToEqLen(s string) string {
	return fmt.Sprintf("\"%s\"%s", s, gstr.Repeat(" ", 20-len(s)))
}

func getApiFields(ctx context.Context, gField map[string]*gdb.TableField, in generateStructDefinitionInput) (apiFields map[string]*ApiField) {
	apiFields = make(map[string]*ApiField)
	for k, v := range gField {
		apiFields[k] = &ApiField{
			TableField: *v,
			GType:      getGTypeByFType(ctx, v.Type, in),
			JsonTag:    gstr.CaseSnakeFirstUpper(v.Name),
		}
	}
	return
}

func getGTypeByFType(ctx context.Context, ltype string, in generateStructDefinitionInput) (gtype string) {
	localTypeName, err := in.DB.CheckLocalTypeForField(ctx, ltype, nil)
	if err != nil {
		panic(err)
	}
	switch localTypeName {
	case gdb.LocalTypeDate, gdb.LocalTypeDatetime:
		if in.StdTime {
			gtype = "time.Time"
		} else {
			gtype = "*gtime.Time"
		}

	case gdb.LocalTypeInt64Bytes:
		gtype = "int64"

	case gdb.LocalTypeUint64Bytes:
		gtype = "uint64"

	// Special type handle.
	case gdb.LocalTypeJson, gdb.LocalTypeJsonb:
		if in.GJsonSupport {
			gtype = "*gjson.Json"
		} else {
			gtype = "string"
		}
	}
	if gtype == "" {
		gtype = "string"
	}
	return
}

func getTableComments(ctx context.Context, createTableSql string) (fieldCommentMap gmap.Map, tableComment, tableName string) {
	// 按照换行符分割文本
	lines := strings.Split(createTableSql, "\n")

	// 循环输出每一行
	for _, line := range lines {
		// 检查 createTableSql 是否包含 comment
		if strings.Contains(line, "--") {
			if strings.Contains(line, "CREATE TABLE") {
				tableName = getLastWord(strings.Split(line, "(")[0])
				tableComment = strings.Split(line, "--")[1]
			} else {
				firstWord := getFirstWord(line)
				lastWord := getLastWord(line)
				fieldCommentMap.Set(firstWord, lastWord)
			}
		}
	}
	return
}

// 判断字符是否为字母、数字或下划线
func isWordChar(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_'
}

// 获取一行文本中的第一个完整单词
func getFirstWord(text string) string {
	fields := strings.FieldsFunc(text, func(r rune) bool { return !isWordChar(r) })
	if len(fields) > 0 {
		return fields[0]
	}
	return ""
}

// 获取一行文本中的最后一个完整单词
func getLastWord(text string) string {
	fields := strings.FieldsFunc(text, func(r rune) bool { return !isWordChar(r) })
	if len(fields) > 0 {
		return fields[len(fields)-1]
	}
	return ""
}

func getProjName() (projName string) {
	// goModPath := gfile.Join(in.Path, "go.mod")
	goModPath := "go.mod"
	if gfile.Exists(goModPath) {
		match, _ := gregex.MatchString(`^module\s+(.+)\s*`, gfile.GetContents(goModPath))
		projName = gstr.Trim(match[1])
	}
	return
}