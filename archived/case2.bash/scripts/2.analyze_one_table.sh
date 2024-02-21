#!/usr/bin/env bash
# leehom Chen clh021@gmail.com
table=$1
db_path=$2
# 使用pragma命令获取表的列信息
# echo sqlite3 "$db_path" ".schema $table"
schema=$(sqlite3 "$db_path" ".schema $table")
# column_info=$(sqlite3 "$db_path" ".schema $table" | grep -E '^ *CREATE' | sed 's/.*\(\([^ ]*\) [^ (]*\).*/\1/')

echo "$schema"


# 用于记录提取到的表名、列名和注释信息
declare -A tables
declare -A columns

# 循环字符串的每一行
while IFS= read -r line; do
  # 检查行是否包含注释
  if [[ $line == *"--"* ]]; then
    # 提取注释信息
    comment=$(echo "$line" | sed 's/.*--\(.*\)/\1/' | sed 's/^[ \t]*//')

    if [[ $line =~ \`(.*)\` ]]; then
      # 提取 ` 字符两侧的内容作为表名或列名
      name="${BASH_REMATCH[1]}"

      if [[ $line == *"CREATE TABLE"* ]]; then
        # 当前行为表名行，提取表名
        tables["$name"]="$comment"
      else
        # 当前行为列名行，提取列名
        columns["$name"]="$comment"
      fi
    fi
  fi
done <<< "$schema"

# 打印提取到的表名、列名和注释信息
for table in "${!tables[@]}"; do
  echo "$table - ${tables[$table]}"
done
for column in "${!columns[@]}"; do
  echo "$column - ${columns[$column]}"
done

pragma=$(sqlite3 "$db_path" "PRAGMA table_info($table)")
# -- cid  name         type                   notnull  dflt_value  pk
# 循环输出每一行，并在行末尾附加这一行的第二列内容
while IFS='|' read -r -a row; do
    # 输出原始行的内容，同时在每个元素后加上列分隔符（除了最后一个元素）
    echo -n "${row[0]}"
    for ((i=1; i<${#row[@]}-1; i++)); do
        echo -n "|${row[i]}"
    done

    # 获取第二列的内容（bash数组索引从0开始，所以第二列是1）
    column2_value="${row[1]}"

    # 在行末尾附加第二列的内容、列分隔符和换行符
    # echo -n "|${column2_value}"
    echo -n "|${columns[$column2_value]}"
    echo  # 添加换行符以区分不同行
done <<< "$pragma"