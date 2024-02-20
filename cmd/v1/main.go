package main

import (
	"context"
	"log"

	"github.com/clh021/gf-gen-code/cmd/v1/gen"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
)
type Command struct {
	*gcmd.Command
}

func GetCommand(ctx context.Context) (*Command, error) {
	root, err := gcmd.NewFromObject(gen.GEN)
	if err != nil {
		panic(err)
	}
	// err = root.AddObject()
	// if err != nil {
	// 	return nil, err
	// }
	command := &Command{
		root,
	}
	return command, nil
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