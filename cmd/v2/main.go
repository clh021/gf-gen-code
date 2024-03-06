package main

import (
	"github.com/clh021/gf-gen-code/utility/mlog"
	"github.com/gogf/gf/cmd/gf/v2/gfcmd"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	var (
		ctx = gctx.GetInitCtx()
	)
	command, err := gfcmd.GetCommand(ctx)
	if err != nil {
		mlog.Fatalf(`%+v`, err)
	}
	if command == nil {
		panic(gerror.New(`retrieve root command failed for "gf"`))
	}
	command.Run(ctx)
}