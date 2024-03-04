#!/usr/bin/env bash
# leehom Chen clh021@gmail.com

# 环境准备
script_dir="$(dirname "$0")"
cd "$(dirname "$script_dir")" || exit 1

# 项目名称
test_proj_name="test_proj"
# 检查项目目录是否已经存在，存在则退出
if [ -d "$test_proj_name" ]; then
  echo "项目目录 $test_proj_name 已经存在，确需重建项目，请删除后重试"
  exit 0
fi

# 初始化项目
gf init "$test_proj_name"

# 建立测试数据库
cp "scripts/resources/db" "$test_proj_name/resource/db"