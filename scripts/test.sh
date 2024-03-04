#!/usr/bin/env bash
# leehom Chen clh021@gmail.com

set -x

# 环境准备
script_dir="$(dirname "$0")"
cd "$(dirname "$script_dir")" || exit 1

# 项目名称
test_proj_name=$(cat scripts/proj_name)

# 检查项目目录是否已经存在，存在则退出
if [ ! -d "$test_proj_name" ]; then
  echo "项目目录 $test_proj_name 不存在，请先建立项目"
  exit 0
fi

cd "$test_proj_name" || exit 1

ls -lah

# ./../tmp/gf_gen -h
./../tmp/gf_gen api1 api