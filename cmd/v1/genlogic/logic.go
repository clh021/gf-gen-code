package genlogic

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gtag"
)

var (
	Logic = cLogic{}
)

type cLogic struct {
	g.Meta `name:"logic" brief:"genereate logic defined go file" eg:"{cLogicEg}" `
}

const (
	cLogicEg = `
gf_gen logic
gf_gen logic -a
gf_gen logic -c
gf_gen logic -cf
`
)

func init() {
	gtag.Sets(g.MapStrStr{
		`cLogicEg`: cLogicEg,
	})
}

type cLogicInput struct {
	g.Meta `name:"logic"  config:"gfcli.logic"`
	All    bool `name:"all" short:"a" brief:"upgrade both version and cli, auto fix codes" orphan:"true"`
	Cli    bool `name:"cli" short:"c" brief:"also upgrade CLI tool" orphan:"true"`
	Fix    bool `name:"fix" short:"f" brief:"auto fix codes(it only make sense if cli is to be upgraded)" orphan:"true"`
}

type cLogicOutput struct{}

func (c cLogic) Index(ctx context.Context, in cLogicInput) (out *cLogicOutput, err error) {
	defer func() {
		if err == nil {
			glog.Print(ctx, `Done! logic defined go file has been generated.`)
		}
	}()
	return
}
