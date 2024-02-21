#!/usr/bin/env bash
# leehom Chen clh021@gmail.com

# 环境准备
# 获取当前脚本所在目录
script_dir="$(dirname "$0")"
scripts_path="$script_dir/scripts"

# 0.检查当前是否处在 gf 项目下
if ! "$scripts_path/0.check_gf_project.sh"; then
  exit 1
fi

# 1.得到当前生成操作的配置
# 使用命令替换捕获子脚本的输出到数组中（每一行作为一个元素）
if ! mapfile -t dao_cfg < <("$scripts_path/1.get_dao_cfg.sh"); then
  exit 1
fi
db_path="${dao_cfg[0]}"
gen_tables="${dao_cfg[1]}"
echo "当前使用数据库:$db_path"
echo "当前生成基础表:$gen_tables"

# 循环处理每个表，获取表详情
for table in $gen_tables; do
  echo "Table: $table"
  # 2.分析表结构
  if ! mapfile -t table_info < <("$scripts_path/2.analyze_one_table.sh" "$table" "$db_path"); then
    exit 1
  fi
  for p in "${table_info[@]}"; do
    echo "  : $p"
  done
done

echo '得出一个数据库中所有数据表结构和 `符合规则的注释`'

echo '1. 根据配置，连接一个 sqlite 数据库'
echo '要读取用户当前目录，识别是否是 gf 项目，符合条件才继续运行'
echo '读取 gf 项目的 hack 配置文件'
echo '2. 得到数据库中的所有表'
echo '3. 分析每一个表的表名，表注释'
echo '4. 分析每一个表的字段名，字段类型，字段注释'
echo '5. 生成 api 接口文件'
echo '6. 生成 dao 文件'
echo '7. 生成 logic 文件'
echo '8. 生成 service 文件'
echo '9. 生成 ctrl 文件'
