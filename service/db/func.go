package db

import (
	"errors"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
)

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