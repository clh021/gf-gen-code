#!/usr/bin/env bash
# leehom Chen clh021@gmail.com
cd "$( dirname "${BASH_SOURCE[0]}" )" || exit
set -e
# 下载 gendao.go
rm -f gendao.go
wget -c https://gitee.com/johng/gf/raw/master/cmd/gf/internal/cmd/gendao/gendao.go

# 替换 package 名称
sed -i'' 's/package gendao/package genapi/' gendao.go
sed -i'' 's/gf gen dao/gf_gen api/' gendao.go
sed -i'' 's/CGenDao/CGenApi/g' gendao.go

# 注释掉 gendao.go 中 315行到334行
start_line=315
end_line=335
{
  head -n $((start_line - 1)) gendao.go &&
  printf "%s\n" "// Start of temporarily commented section" &&
  sed -n "$start_line,$end_line p" gendao.go |
    while IFS= read -r line; do
      printf "%s\n" "// $line"
    done &&
  printf "%s\n" "// End of temporarily commented section" &&
  tail -n +$((end_line + 1)) gendao.go
} > gendao.tmp.go && mv gendao.tmp.go gendao.go

# 删除包含 "/internal/" 的行
sed '/\/internal\//d' gendao.go > gendao.tmp
mv gendao.tmp gendao.go

# 格式化代码
# 使用 goimports 格式化代码
goimports -w gendao.go

echo "gendao.go has been formatted with goimports."