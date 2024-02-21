// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package sqlite

import (
	"context"
	"fmt"
	"strings"
	"unicode"

	"github.com/clh021/gf-gen-code/cmd/v1/mlog"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/util/gutil"
)

// TableFields retrieves and returns the fields' information of specified table of current schema.
//
// Also see DriverMysql.TableFields.
func (d *Driver) TableFields(ctx context.Context, table string, schema ...string) (fields map[string]*gdb.TableField, err error) {
	var (
		result     gdb.Result
		link       gdb.Link
		usedSchema = gutil.GetOrDefaultStr(d.GetSchema(), schema...)
	)
	if link, err = d.SlaveLink(usedSchema); err != nil {
		return nil, err
	}
	result, err = d.DoSelect(ctx, link, fmt.Sprintf(`PRAGMA TABLE_INFO(%s)`, d.QuoteWord(table)))
	if err != nil {
		return nil, err
	}
	// Get comment
	schemaRes, err := d.GetValue(ctx, fmt.Sprintf(`SELECT sql FROM sqlite_master WHERE type='table' AND name='%s';`, table))
	if err != nil {
		return nil, err
	}
	mlog.Debug("==schemaRes==:")
	mlog.Debug(schemaRes)
	comments, _, _ := d.getComments(schemaRes.String())
	mlog.Debug("==comments==:")
	mlog.Debug(comments)

	fields = make(map[string]*gdb.TableField)
	for i, m := range result {
		mKey := ""
		if m["pk"].Bool() {
			mKey = "pri"
		}
		_comment, _ := comments.Get(m["name"].String()).(string)
		fields[m["name"].String()] = &gdb.TableField{
			Index:   i,
			Name:    m["name"].String(),
			Type:    m["type"].String(),
			Key:     mKey,
			Default: m["dflt_value"].Val(),
			Null:    !m["notnull"].Bool(),
			Comment: _comment,
		}
	}
	return fields, nil
}

func (d *Driver) getComments(createTableSql string) (fieldCommentMap gmap.Map, tableComment, tableName string) {
	// 按照换行符分割文本
	lines := strings.Split(createTableSql, "\n")

	// 循环输出每一行
	for _, line := range lines {
		// 检查 createTableSql 是否包含 comment
		if strings.Contains(line, "--") {
			if strings.Contains(line, "CREATE TABLE") {
				tableName = d.getLastWord(strings.Split(line, "(")[0])
				tableComment = strings.Split(line, "--")[1]
			} else {
				firstWord := d.getFirstWord(line)
				lastWord := d.getLastWord(line)
				fieldCommentMap.Set(firstWord, lastWord)
			}
		}
	}
	return
}


// 判断字符是否为字母、数字或下划线
func (d *Driver) isWordChar(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_'
}

// 获取一行文本中的第一个完整单词
func (d *Driver) getFirstWord(text string) string {
	fields := strings.FieldsFunc(text, func(r rune) bool { return !d.isWordChar(r) })
	if len(fields) > 0 {
		return fields[0]
	}
	return ""
}

// 获取一行文本中的最后一个完整单词
func (d *Driver) getLastWord(text string) string {
	fields := strings.FieldsFunc(text, func(r rune) bool { return !d.isWordChar(r) })
	if len(fields) > 0 {
		return fields[len(fields)-1]
	}
	return ""
}
