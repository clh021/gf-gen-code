package db

import (
	"embed"
	"fmt"
	"io"
	"os"
	"strings"
)

// 将 同级别目录下的 test.sql 文件，存储为编译后的变量

//go:embed test.sql
var sqlFile embed.FS

func GetTestSql() []string {
	file, err := sqlFile.Open("test.sql")
	if err != nil {
		fmt.Println("Failed to open embedded SQL file:", err)
		os.Exit(1)
	}
	defer file.Close()
	// 读取嵌入的SQL文件内容
	contentBytes, err := io.ReadAll(file)
	if err != nil {
		panic(fmt.Sprintf("Failed to read embedded SQL file: %v", err))
	}

	// 转换为字符串
	return strings.Split(string(contentBytes), ";")
}