package db

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/frame/g"
)

func TestGetComment() {
	sqlStatement := GetTestSql()

	d := NewDriver()
	for _, sqlString := range sqlStatement {
		if len(sqlString) > 0 {
			fieldsComments, tableComment, tableName := d.getComments(sqlString)
			fmt.Println("------------------------------------")
			fmt.Printf("tableName:%s \n", tableName)
			fmt.Printf("tableComment:%s \n", tableComment)
			fmt.Println("fieldsComments:")
			g.Dump(fieldsComments)
		}
	}
}

type Driver struct {
}

func NewDriver() *Driver {
	return &Driver{}
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
