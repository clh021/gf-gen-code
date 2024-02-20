package gen

import (
	"context"

	"github.com/clh021/gf-gen-code/cmd/v1/mlog"
	"github.com/gogf/gf/v2/frame/g"
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
	Version bool `short:"v" name:"version"     brief:"Display the program's version information"    orphan:"true"`
	Debug   bool `short:"d" name:"debug"       brief:"Display debug information during scanning"    orphan:"true"`
}

type cOutput struct{}

func (c cGEN) setInputDefault(ctx context.Context, in cInput) (inNew cInput, out *cOutput, err error) {
	// 设置默认值
	inNew = in
	return
}
func (c cGEN) validInput(ctx context.Context, in cInput) (out *cOutput, err error) {
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

	// Set Input Default
	in, _, err = c.setInputDefault(ctx, in)
	if err != nil {
		return
	}

	// Valid Input
	out, err = c.validInput(ctx, in)
	if err != nil {
		return
	}

	// gcmd.CommandFromCtx(ctx).Print()
	return
}
