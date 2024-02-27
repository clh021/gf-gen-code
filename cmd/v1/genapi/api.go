package genapi

import (
	"context"

	"github.com/clh021/gf-gen-code/cmd/v1/mlog"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gtag"
)

var (
	Api = cApi{}
)

type cApi struct {
	g.Meta `name:"api" brief:"genereate api defined go file" eg:"{cApiEg}" `
}

const (
	cApiEg     = `
gf_gen api
gf_gen api -a
gf_gen api -c
gf_gen api -cf
`
)

func init() {
	gtag.Sets(g.MapStrStr{
		`cApiEg`: cApiEg,
	})
}

type cApiInput struct {
	g.Meta `name:"api"  config:"gfcli.api"`
	All    bool `name:"all" short:"a" brief:"upgrade both version and cli, auto fix codes" orphan:"true"`
	Cli    bool `name:"cli" short:"c" brief:"also upgrade CLI tool" orphan:"true"`
	Fix    bool `name:"fix" short:"f" brief:"auto fix codes(it only make sense if cli is to be upgraded)" orphan:"true"`
}

type cApiOutput struct{}

func (c cApi) Index(ctx context.Context, in cApiInput) (out *cApiOutput, err error) {
	defer func() {
		if err == nil {
			mlog.Print(`Done! api defined go file has been generated.`)
		}
	}()
	mlog.Print(in.All)
	return
}