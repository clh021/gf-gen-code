package root

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gtag"
)

var (
	C = cC{}
)

type cC struct {
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

//go:generate goimports -w .
//go install golang.org/x/tools/cmd/goimports@latest

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
	Yes     bool `short:"y" name:"yes"         brief:"all yes for all command without prompt ask"  orphan:"true"`
	Version bool `short:"v" name:"version"     brief:"Display the program's version information"   orphan:"true"`
	Debug   bool `short:"d" name:"debug"       brief:"Display debug information during running"    orphan:"true"`
}

type cOutput struct{}

func (c cC) Index(ctx context.Context, in cInput) (out *cOutput, err error) {
	if in.Debug {
		glog.SetDebug(true)
	}

	// Show Version
	if in.Version {
		_, err = Version.Index(ctx, cVersionInput{})
		return
	}

	// gcmd.CommandFromCtx(ctx).Print()
	return
}
