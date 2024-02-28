package test

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gtag"
)

var (
	Test = cTest{}
)

type cTest struct {
	g.Meta `name:"test" brief:"genereate test defined go file" eg:"{cTestEg}" `
}

const (
	cTestEg = `
gf_gen test
gf_gen test -a
gf_gen test -c
gf_gen test -cf
`
)

// go:generate pwd # 没一个 go:generate 都是相对于当前文件所属目录运行的

func init() {
	gtag.Sets(g.MapStrStr{
		`cTestEg`: cTestEg,
	})
}

type cTestInput struct {
	g.Meta `name:"test"  config:"gfcli.test"`
	All    bool `name:"all" short:"a" brief:"upgrade both version and cli, auto fix codes" orphan:"true"`
	Cli    bool `name:"cli" short:"c" brief:"also upgrade CLI tool" orphan:"true"`
	Fix    bool `name:"fix" short:"f" brief:"auto fix codes(it only make sense if cli is to be upgraded)" orphan:"true"`
}

type cTestOutput struct{}

func (c cTest) Index(ctx context.Context, in cTestInput) (out *cTestOutput, err error) {
	defer func() {
		if err == nil {
			glog.Print(ctx, `Done!`)
		}
	}()
	return
}
