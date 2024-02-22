package cfg

import (
	"context"

	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gfile"
)

func GetByFilePath(ctx context.Context, path string) (cfg *gcfg.Config, err error) {
	// 默认配置文件行为
	// 单例对象在创建时会按照文件后缀toml/yaml/yml/json/ini/xml/properties自动检索配置文件。
	// 默认情况下会自动检索配置文件config.toml/yaml/yml/json/ini/xml/properties并缓存，配置文件在外部被修改时将会自动刷新缓存。
	//
	// 当前工作目录
	// USER_PWD/
	// USER_PWD/config
	// USER_PWD/manifest/config
	// 当前可执行文件所在目录
	// BIN_PATH/
	// BIN_PATH/config
	// BIN_PATH/manifest/config
	// 当前main源代码包所在目录
	// CODE_PATH/
	// CODE_PATH/config
	// CODE_PATH/manifest/config
	// GF_DEBUG=true ./main // 有时也许可以帮上忙
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
