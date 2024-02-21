#!/bin/bash

# # 接收从父脚本传入的两个参数
# param1=$1
# param2=$2

# # 输出接收到的参数
# echo "Received parameter 1: $param1"
# echo "Received parameter 2: $param2"


# 假设SQL创建表语句存储在sql_create_table变量中
sql_create_table=$(cat << EOF
CREATE TABLE IF NOT EXISTS `test` (                 -- 书本管理
  `id` bigint NOT NULL,                             -- 编号
  `author_id` bigint DEFAULT '0',                   -- 会员编号
  `class` varchar(64) DEFAULT '',                   -- 应用编号
  `name` varchar(255) DEFAULT '',                   -- 应用名称
  `price` decimal(10,2) DEFAULT '0.00',             -- 年龄
  `cover_fid` bigint DEFAULT '0',                   -- 封面图片|上传图片文件的编号
  `author_fid` varchar(100) NOT NULL DEFAULT '',    -- 作者图片|上传图片文件的编号
  `status` tinyint(1) DEFAULT '1',                  -- 状态|1:正常|0:删除
  `created_at` datetime DEFAULT NULL,               -- 添加时间
  `updated_at` datetime DEFAULT NULL,               -- 更新时间
  `deleted_at` datetime DEFAULT NULL                -- 删除时间
);
EOF
)

# 提取表名
table_name=$(echo "$sql_create_table" | grep -oP '(?<=CREATE TABLE IF NOT EXISTS `)\w+(?=`)')

# 提取列名和注释
column_info=$(echo "$sql_create_table" | grep -Eo '`[^\s]+`\s.*?-- [^;]*')

# 输出结果
echo "Table Name: $table_name"
echo "Column Info:"
echo "$column_info"

# 如果你希望将每一项转换为键值对格式，可以进一步处理：
while read -r line; do
  # 使用sed匹配列名和注释
  column=$(echo "$line" | sed -E 's/^`([^`]+)`\s.*-- (.*)/\1/')
  comment=$(echo "$line" | sed -E 's/^`[^`]+`\s.*-- (.*)/\1/')
  echo "Column: $column, Comment: $comment"
done <<< "$column_info"