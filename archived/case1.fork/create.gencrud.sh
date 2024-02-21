#!/usr/bin/env bash
# leehom Chen clh021@gmail.com

# set -x
rootProjPath="$HOME/Projects/gf/cmd/gf/"
rootProjMain="$rootProjPath/main.go"
gen_basefile="internal/cmd/genctrl/genctrl.go"
# gen_basefile_basestr="func (c CGenCtrl) Ctrl(ctx context.Context, in CGenCtrlInput) (out *CGenCtrlOutput, err error) {"
gen_basefile_basestr="func (c CGenCtrl) Ctrl("
gencrud_file="internal/cmd/gencrud/gencrud.go"
cmd_gen_basefile="internal/cmd/cmd_gen_ctrl.go"
cmd_gencrud_file="internal/cmd/cmd_gen_crud.go"

# 环境检查
# 修改好的代码将提交保存至 `git clone --depth=1 git@github.com:clh021/gf.git`
# 检查是否存在 `~/Projects/gf/cmd/gf/main.go` ，如果不存在则提示用户先 `git clone --depth=1 https://github.com/gogf/gf.git`
if [ ! -f "$rootProjMain" ]; then
    echo "未找到 $rootProjMain 项目，请先 cd ~/Projects/; git clone --depth=1 https://github.com/gogf/gf.git"
    exit 1
fi

# 进入项目
echo "正在进入项目"
cd "$rootProjPath" || exit
# 查找 `internal/cmd/genctrl/genctrl.go`, 是否存在
if [ ! -f $gen_basefile ]; then
    echo "未找到 $gen_basefile 文件"
    exit 1
fi

pwd
# 查找文件中的 `func (c CGenCtrl) Ctrl(ctx context.Context, in CGenCtrlInput) (out *CGenCtrlOutput, err error) {` 所在的行，如果没有找到，直接提示，并退出脚本。如果找到了，就保存包括这一行在内的文件上半部分，到另一个文件`internal/cmd/gencrud/gencrud.go`中
line_number=$(awk -v target="$gen_basefile_basestr" 'index($0, target) == 1 { print NR; exit }' "$gen_basefile")
# line_number=$(grep -n "'$gen_basefile_basestr'" "$gen_basefile" | cut -d ':' -f 1)
if [[ -z $line_number ]]; then
  echo "没有找到目标函数: ${gen_basefile_basestr:0:42}"
  exit 1
fi
file_content=$(head -n "$line_number" "$gen_basefile")
echo "$file_content" > "$gencrud_file"
echo "	return" >> "$gencrud_file"
echo "}" >> "$gencrud_file"
echo "基于参考文件生成了 $gencrud_file."

# TODO: 手动合并了一些传入参数

#，然后，替换所有的 ``
sed -i 's/ctrl/crud/g' "$gencrud_file"
sed -i 's/Ctrl/Crud/g' "$gencrud_file"

cp "$cmd_gen_basefile" "$cmd_gencrud_file"
sed -i 's/ctrl/crud/g' "$cmd_gencrud_file"
sed -i 's/Ctrl/Crud/g' "$cmd_gencrud_file"

echo "下一步需要自行编写 $gencrud_file 了"