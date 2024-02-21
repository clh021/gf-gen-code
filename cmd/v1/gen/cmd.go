package gen

import (
	"context"
	"log"

	"github.com/clh021/gf-gen-code/cmd/v1/mlog"
	"github.com/clh021/gf-gen-code/service/cfg"
	"github.com/clh021/gf-gen-code/service/db"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/gtag"
)

var (
	GEN = cGEN{}
)

type cGEN struct {
	g.Meta `name:"{cName}" ad:"{cAd}"`
}

const (
	cName = `gf_gen`
	cAd   = `
ADDITIONAL
    Use "gf_gen -h"                                 to show help information.
    Use "gf_gen -p ./path/to/code"                  to scan files in the code directory.
`
)

func init() {
	gtag.Sets(g.MapStrStr{
		`cName`: cName,
		`cAd`:   cAd,
	})
}

type cInput struct {
	// 这里不要加太多校验规则，因为参数不通过时，无法友好的将错误提示给用户
	// 也无法进入自主流程，所以不使用校验规则
	g.Meta  `name:"{cName}"`
	Version bool   `short:"v" name:"version"     brief:"Display the program's version information"   orphan:"true"`
	Debug   bool   `short:"d" name:"debug"       brief:"Display debug information during running"    orphan:"true"`
	Cfg     string `short:"c" name:"cfg"         brief:"Config file path"                            d:"./hack/config.yaml"`
}

type cOutput struct{}

func (c cGEN) validInput(ctx context.Context, in cInput) (out *cOutput, err error) {
	if in.Cfg == "" {
		mlog.Fatal(`Please provide the required parameter: cfg. Use the '-c' or '--cfg' option to specify the config file.`)
		return
	}
	pathValid := gfile.IsFile(in.Cfg)
	if !pathValid {
		mlog.Fatalf(`The specified config file "%s" does not exist. Please provide a valid path using the '-c' or '--cfg' option.`, in.Cfg)
		return
	}
	return
}

func (c cGEN) Index(ctx context.Context, in cInput) (out *cOutput, err error) {
	if in.Debug {
		mlog.SetDebug(true)
	}

	// Show Version
	if in.Version {
		_, err = Version.Index(ctx, cVersionInput{})
		return
	}

	// Valid Input
	out, err = c.validInput(ctx, in)
	if err != nil {
		return
	}

	cfg, err := cfg.GetByFilePath(ctx, in.Cfg)
	if err != nil {
		mlog.Fatal(err)
		return
	}
	link := cfg.MustGet(ctx,"gfcli.gen.dao.link").String()
	table := cfg.MustGet(ctx,"gfcli.gen.dao.table").String()
	log.Printf("link: %s", link)
	log.Printf("table: %s", table)
	Db, err := db.New(link, table, ctx)
	if err != nil {
		return nil, err
	}
	tables, err := Db.CheckMergeTables()
	if err != nil {
		return nil, err
	}
	log.Println("MergeTables:")
	g.Dump(tables)
	fields, err := Db.Fields(tables[0])
	if err != nil {
		return nil, err
	}
	g.Dump(fields)

	// gcmd.CommandFromCtx(ctx).Print()
	return
}
