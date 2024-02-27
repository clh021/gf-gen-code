package main

import (
	"context"
	"log"

	"github.com/clh021/gf-gen-code/cmd/v1/genapi"
	"github.com/clh021/gf-gen-code/cmd/v1/genlogic"
	"github.com/clh021/gf-gen-code/cmd/v1/genweb"
	"github.com/clh021/gf-gen-code/cmd/v1/root"
	"github.com/clh021/gf-gen-code/cmd/v1/test"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
)

type Command struct {
	*gcmd.Command
}

func GetCommand(ctx context.Context) (*Command, error) {
	root, err := gcmd.NewFromObject(root.C)
	if err != nil {
		panic(err)
	}
	if err = root.AddObject(
		genapi.Api,
		genlogic.Logic,
		genweb.Web,
		test.Test,
	); err != nil {
		return nil, err
	}
	return &Command{root}, nil
}

func main() {
	var (
		ctx = gctx.GetInitCtx()
	)
	command, err := GetCommand(ctx)
	if err != nil {
		log.Fatalf(`%+v`, err)
	}
	if command == nil {
		panic(gerror.New(`retrieve root command failed`))
	}
	command.Run(ctx)
	// command.RunWithError(ctx)
}
