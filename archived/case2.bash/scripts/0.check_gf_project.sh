#!/usr/bin/env bash
# leehom Chen clh021@gmail.com

# 读取用户所在目录，检查是否存在以下文件和目录，如果缺少某个目录，则提示，请切换到 gf 项目目录下。
# drwxr-xr-x  3 lee lee 4.0K Feb 18 09:21 api/
# -rw-r--r--  1 lee lee   22 Feb 18 09:21 .gitattributes
# -rw-r--r--  1 lee lee  234 Feb 18 09:21 .gitignore
# -rw-r--r--  1 lee lee 1.8K Feb 18 09:21 go.mod
# -rw-r--r--  1 lee lee  11K Feb 18 09:21 go.sum
# drwxr-xr-x  2 lee lee 4.0K Feb 18 09:21 hack/
# drwxr-xr-x 11 lee lee 4.0K Feb 18 09:21 internal/
# -rw-r--r--  1 lee lee  237 Feb 18 09:21 main.go
# -rw-r--r--  1 lee lee  138 Feb 18 09:21 Makefile
# drwxr-xr-x  7 lee lee 4.0K Feb 18 09:28 manifest/
# -rw-r--r--  1 lee lee  107 Feb 18 14:52 README.MD
# drwxr-xr-x  4 lee lee 4.0K Feb 20 16:25 resource/
# drwxr-xr-x  2 lee lee 4.0K Feb 18 09:21 utility/

# 定义需要检查的文件和目录列表
declare -a REQUIRED_DIRS=("api" "hack" "internal" "manifest" "resource" "utility")
REQUIRED_FILES=(.gitattributes .gitignore go.mod go.sum main.go Makefile README.MD)

# 检查目录是否存在
for dir in "${REQUIRED_DIRS[@]}"; do
  if [ ! -d "$dir" ]; then
    echo "缺失目录：$dir，请确保您在 gf 项目目录下。"
    exit 1
  fi
done

# 检查文件是否存在
for file in "${REQUIRED_FILES[@]}"; do
  if [ ! -f "$file" ]; then
    echo "缺失文件：$file，请确保您在 gf 项目目录下。"
    exit 1
  fi
done
