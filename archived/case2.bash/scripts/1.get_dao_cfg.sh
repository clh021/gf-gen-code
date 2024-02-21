#!/usr/bin/env bash
# leehom Chen clh021@gmail.com
# 检查 hack/config.yaml 文件是否存在
if [ ! -f "hack/config.yaml" ]; then
  echo "配置文件 hack/config.yaml 不存在！"
  exit 1
fi

# 检查 yq 是否已安装，如果没有安装，则提示用户安装
if ! command -v yq &> /dev/null; then
  echo "yq 命令不存在，请安装 yq : sudo pacman -S go-yq"
  echo "或者使用: go install github.com/mikefarah/yq/v4@latest"
  exit 1
fi

# 如果文件 hack/config.yaml 存在则读取该配置文件，如果不存在下列配置则退出，存在则输出配置中的 link 和table值。
# gfcli:
#   gen:
#     dao:
#       link: "sqlite::@file(./resource/db.sqlite3)"
#       table: "user,detect_record,detect_item,environment"
# 使用 yq 读取并检查 gfcli.gen.dao.link 和 gfcli.gen.dao.table 是否存在
dao_link=$(yq '.gfcli.gen.dao.link' hack/config.yaml)
dao_table=$(yq '.gfcli.gen.dao.table' hack/config.yaml)

# 检查 link 和 table 是否为空
if [[ -z "$dao_link" ]]; then
  echo "配置文件中缺少 gfcli.gen.dao.link 值！"
  exit 1
fi

gen_tables=""
db_path=$(echo "$dao_link" | sed 's/^.*@file\(\(.*\)\)$/\1/' | tr -d '()')

# 依据 gf 规则，如果 dao_table 为空，则使用所有表名
if [[ -z "$dao_table" ]]; then
  gen_tables=$(sqlite3 "$db_path" ".tables")
else
  # gen_tables=$(echo "$dao_table" | tr ',' '\n' | tr -d ' ' | tr '\n' ',')
  IFS=',' read -ra dao_table_arr <<< "$dao_table"
  gen_tables=${dao_table_arr[*]}
fi

echo "$db_path"
echo "$gen_tables"