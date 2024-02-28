package genweb

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gtag"
)

var (
	Web = cWeb{}
)

type cWeb struct {
	g.Meta `name:"web" brief:"genereate web defined go file" eg:"{cWebEg}" `
}

const (
	cWebEg = `
gf_gen web
gf_gen web -a
gf_gen web -c
gf_gen web -cf
`
)

func init() {
	gtag.Sets(g.MapStrStr{
		`cWebEg`: cWebEg,
	})
}

type cWebInput struct {
	g.Meta `name:"web"  config:"gfcli.web"`
	All    bool `name:"all" short:"a" brief:"upgrade both version and cli, auto fix codes" orphan:"true"`
	Cli    bool `name:"cli" short:"c" brief:"also upgrade CLI tool" orphan:"true"`
	Fix    bool `name:"fix" short:"f" brief:"auto fix codes(it only make sense if cli is to be upgraded)" orphan:"true"`
}

type cWebOutput struct{}

func (c cWeb) Index(ctx context.Context, in cWebInput) (out *cWebOutput, err error) {
	defer func() {
		if err == nil {
			glog.Print(ctx, `Done!`)
		}
	}()
	return
}
