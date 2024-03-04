package cmd

import (
	"context"

	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:        "test_proj",
		Usage:       "test_proj",
		Brief:       "start http server",
		Description: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			requiredFiles := []string{"config.yaml", "resource/db/test.db"}
			if SureAllFilesExist(ctx, requiredFiles) {
				runServe(ctx)
			}
			return nil
		},
	}
	Init = gcmd.Command{
		Name:        "init",
		Usage:       "init",
		Brief:       "generate init resource",
		Description: "generate init resource to run server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			autoGenerateFile(ctx, "config.yaml", "resource/template/config.yaml")
			// autoGenerateFile(ctx, "resource/db.sqlite3", "resource/db/test.db")
			return nil
		},
	}
)
