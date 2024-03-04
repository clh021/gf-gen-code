package genapi

import (
	"context"
	"log"
	"path/filepath"

	"github.com/clh021/gf-gen-code/service/db"
	"github.com/clh021/gf-gen-code/service/tpl"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gtag"
)

var (
	Api = cApi{}
)

type cApi struct {
	g.Meta `name:"api1" brief:"genereate api defined go file" eg:"{cApiEg}" `
}

//go:generate ./update.depend.sh

const (
	cApiEg = `
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
			glog.Print(ctx, `Done! api defined go file has been generated.`)
		}
	}()
	glog.Print(ctx, in.All)

	dstModuleFolderPath := "api/hello/v1/"
	module := "user2"
	var (
		moduleFilePath = filepath.FromSlash(gfile.Join(dstModuleFolderPath, module+".go"))
	// 	moduleFilePathNew     = filepath.FromSlash(gfile.Join(dstModuleFolderPath, module+"_new.go"))
	// 	ctrlName              = fmt.Sprintf(`Controller%s`, gstr.UcFirst(version))
	// 	interfaceName         = fmt.Sprintf(`%s.I%s%s`, module, gstr.CaseCamel(module), gstr.UcFirst(version))
	// 	newFuncName           = fmt.Sprintf(`New%s`, gstr.UcFirst(version))
	// 	newFuncNameDefinition = fmt.Sprintf(`func %s()`, newFuncName)
	// 	alreadyCreated        bool
	)
	if !gfile.Exists(moduleFilePath+".go") {
		t := tpl.New("/")
		if err := t.Write(moduleFilePath, "gen_templates/api.tpl", g.Map{
			"name":    "test123",
			"Module":  module,
			"version": "jajajajaj",
		}); err != nil {
			glog.Fatal(ctx, err)
		}
		glog.Print(ctx, "--------------------------------")
		glog.Printf(ctx, "generated: %s", moduleFilePath)
	} else {
		glog.Printf(ctx, "already exists: %s", moduleFilePath)
	}

	// cfg, err := cfg.GetByFilePath(ctx, in.Cfg)
	// if err != nil {
	// 	glog.Fatal(ctx, err)
	// 	return
	// }
	// link := cfg.MustGet(ctx, "gfcli.gen.dao.link").String()
	// table := cfg.MustGet(ctx, "gfcli.gen.dao.table").String()
	// glog.Printf(ctx, "link: %s", link)
	// glog.Printf(ctx, "link: %s", link)
	// glog.Printf(ctx, "table: %s", table)
	// // db.TestGetComment()
	// getTableStruct(link, table, ctx)
	return
}

func getTableStruct(link, table string, ctx context.Context) {
	Db, err := db.New(link, table, ctx)
	if err != nil {
		glog.Fatal(ctx, err)
	}
	tables, err := Db.CheckMergeTables()
	if err != nil {
		glog.Fatal(ctx, err)
	}
	log.Println("MergeTables:")
	g.Dump(tables)
	// 进行时，发现 gf_gen 是会处理表注释的，只是 sqlite 没有像 mysql 一样直接支持。
	// 按照本项目的思路
	// 修改 sqlite 对应的代码文件
	// contrib/drivers/sqlite/sqlite_table_fields.go
	// contrib/drivers/sqlite/sqlite_z_unit_core_test.go
	// 像 mysql 对应的代码文件一样即可
	// contrib/drivers/mysql/mysql_table_fields.go
	// contrib/drivers/mysql/mysql_z_unit_core_test.go
	// 尝试直接修改这个部分，也许还能合并进去，及时不能合并进去，也可以先自己用。
	// 就这么干……
	fields, err := Db.Fields(tables[0])
	if err != nil {
		glog.Fatal(ctx, err)
	}
	g.Dump(fields)
}
