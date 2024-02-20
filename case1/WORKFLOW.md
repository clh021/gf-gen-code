# 记录代码生成器修改记录和分析思路

> 记录后，已备后期改进或升级框架脚手架时查阅，同时也方便其他开发者参与进来

本项目旨在，通过最小代码改动，实现代码生成器的优化和最小维护。

## 程序运行思路分析
代码生成程序是 `gf` 框架中的一个子项目。位于 `gf/cmd/gf/` 目录。
下面(及后面的其它章节)以 `gf/cmd/gf/` 作为讲解的根目录进行分析。

1. 程序启动从入口 `main.go` 开始，使用 `gfcmd/gfcmd.go` 包的 `GetCommand(ctx)` 方法获取命令行参数，并执行相应的命令。
2. `GetCommand(ctx)` 方法中 `root.AddObject(...)` 获取了很多命令对象，其中就包括 `cmd.Gen`，我们需要重点关注的就是 `cmd.Gen` ，对应 `internal/cmd/cmd_gen.go` 文件。
3. 其中我们要改变生成 `ctrl` 的动作，可以看到，其对应 `internal/cmd/cmd_gen_ctrl.go` 文件，并最终链接到 `internal/cmd/genctrl/genctrl.go` 文件。

## 修改步骤
```bash
# 1. 创建文件夹
mkdir -p internal/cmd/gencrud
```