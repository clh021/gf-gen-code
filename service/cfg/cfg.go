package cfg

import (
	"context"

	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gfile"
)

func GetByFilePath(ctx context.Context, path string) (cfg *gcfg.Config, err error) {
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
	cfgContent := gfile.GetContents(path)
	adapter, err := gcfg.NewAdapterContent(cfgContent)
	if err != nil {
		return
	}
	cfg = gcfg.NewWithAdapter(adapter)
	// link = config.MustGet(ctx,"gfcli.gen.dao.link").String()
	// table = config.MustGet(ctx,"gfcli.gen.dao.table").String()
	return
}
