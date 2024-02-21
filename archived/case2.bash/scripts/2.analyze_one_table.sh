#!/usr/bin/env bash
# leehom Chen clh021@gmail.com
table=$1
db_path=$2
# 使用pragma命令获取表的列信息
echo sqlite3 "$db_path" ".schema $table"
column_info=$(sqlite3 "$db_path" ".schema $table")
# column_info=$(sqlite3 "$db_path" ".schema $table" | grep -E '^ *CREATE' | sed 's/.*\(\([^ ]*\) [^ (]*\).*/\1/')

# 输出列名
echo "Columns: $column_info"
echo "" # 添加一个空行分隔不同表的信息
