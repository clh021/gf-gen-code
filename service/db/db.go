package db

import (
	"context"
	"errors"

	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
)

type DB struct {
	Clink string
	Ctable string
	db   gdb.DB
}

func New(link, table string) (db *DB, err error) {
	_db, err := getDB(link)
	db = &DB{
		Clink: link,
		Ctable: table,
		db: _db,
	}
	return
}

func (d *DB) CheckMergeTables(ctx context.Context) (tables []string, err error) {
	if d.Ctable != "" {
		tables = gstr.SplitAndTrim(d.Ctable, ",")
	} else {
		tables, err = d.db.Tables(ctx)
	}
	return
}

func getDB(link string) (db gdb.DB, err error) {
	var (
		tempGroup = gtime.TimestampNanoStr()
		match, _  = gregex.MatchString(`([a-z]+):(.+)`, link)
	)
	if len(match) == 3 {
		gdb.AddConfigNode(tempGroup, gdb.ConfigNode{
			Type: gstr.Trim(match[1]),
			Link: link,
		})
		db, err = gdb.Instance(tempGroup)
		return
	}
	err = errors.New("link config error")
	return
}
