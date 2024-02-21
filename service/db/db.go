package db

import (
	"context"

	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/text/gstr"
)

type DB struct {
	Clink string
	Ctable string
	db   gdb.DB
	ctx context.Context
}

func New(link, table string, ctx context.Context) (db *DB, err error) {
	_db, err := getDB(link)
	db = &DB{
		Clink: link,
		Ctable: table,
		db: _db,
		ctx: ctx,
	}
	return
}

func (d *DB) CheckMergeTables() (tables []string, err error) {
	if d.Ctable != "" {
		tables = gstr.SplitAndTrim(d.Ctable, ",")
	} else {
		tables, err = d.db.Tables(d.ctx)
	}
	return
}

func (d *DB) Fields(tablename string) (fMap map[string]*gdb.TableField, err error) {
	fMap, err = d.db.TableFields(d.ctx, tablename)
	return
}
