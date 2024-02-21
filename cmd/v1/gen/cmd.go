package gen

import (
	"context"

	"github.com/clh021/gf-gen-code/cmd/v1/mlog"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
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

func (c cGEN) setConfig(ctx context.Context, cfgPath string) (link, table string) {
	// 默认配置文件行为
	// 会按照文件后缀toml/yaml/yml/json/ini/xml/properties自动检索配置文件。
	// 当前工作目录,当前可执行文件所在目录,当前main源代码包所在目录
	// ./
	// ./config
	// ./manifest/config
	// ------------------------------------------------------------------
	// 这里将消除默认行为，直接读取用户配置的配置文件路径
	// cfgDir := gfile.Dir(cfgPath)
	// genv.Set("GF_GCFG_PATH", cfgDir) // -- 消除多级目录搜索成功
	// g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetPath(cfgDir) // -- 消除多级搜索失败
	// paths := 	g.Cfg().GetAdapter().(*gcfg.AdapterFile).GetPaths()
	// g.Dump(paths)
	// genv.Set("GF_GCFG_FILE", "configu9g0dsa.gr8ewqtn4m3qgf.dewgrewjgrkewfje.prod") // -- 修改默认配置文件名成功
	// g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName("default.yaml") // -- 修改默认配置文件名成功
	// cfgName := 	g.Cfg().GetAdapter().(*gcfg.AdapterFile).GetFileName()
	// g.Dump(cfgName)
	// 不使用框架自带配置，独立使用传参来获取数据结果
	cfgContent := gfile.GetContents(cfgPath)
	adapter, err := gcfg.NewAdapterContent(cfgContent)
	if err != nil {
		panic(err)
	}
	config := gcfg.NewWithAdapter(adapter)
	link = config.MustGet(ctx,"gfcli.gen.dao.link").String()
	table = config.MustGet(ctx,"gfcli.gen.dao.table").String()
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

	g.Dump(in)
	link, table := c.setConfig(ctx, in.Cfg)
	g.Dump(link)
	g.Dump(table)

	// gcmd.CommandFromCtx(ctx).Print()
	return
}
