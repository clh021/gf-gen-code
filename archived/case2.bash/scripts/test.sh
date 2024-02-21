#!/usr/bin/env bash
db_path=./resource/test.db
table=book
sqlite3 "$db_path" ".headers on"
pragma=$(sqlite3 "$db_path" "PRAGMA table_info($table)")

# -- cid  name         type                   notnull  dflt_value  pk
echo "$pragma"

echo "----------------------------------------"

# 使用awk处理SQLite的输出以形成JSON友好的格式
pragma_result=$(sqlite3 "$db_path" "PRAGMA table_info($table)" | awk -F '|' '{OFS=","; print "{\"cid\":\""$1"\",\"name\":\""$2"\",\"type\":\""$3"\",\"notnull\": \""$4"\",\"dflt_value\":\""$5"\",\"pk\":\""$6"\"}" }')

echo "$pragma_result"

# 用jq添加数组结构并美化输出
json_output=$(echo "["$pragma_result"]" | jq -s '.')

echo "$json_output"