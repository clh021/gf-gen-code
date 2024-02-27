// Copyright GoFrame gf Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package genweb

import (
	"context"
	"fmt"

	"github.com/clh021/gf-gen-code/cmd/v1/mlog"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gproc"
	"github.com/gogf/gf/v2/util/gtag"
)

var (
	Web = cWeb{}
)

type cWeb struct {
	g.Meta `name:"web" brief:"genereate web defined go file" eg:"{cWebEg}" `
}

const (
	cWebEg     = `
gf web
gf web -a
gf web -c
gf web -cf
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
			mlog.Print()
			mlog.Print(`Done! web defined go file has been generated.`)
			mlog.Print()
		}
	}()
	return
}

type doUpgradeVersionOutput struct {
	Items []doUpgradeVersionOutputItem
}

type doUpgradeVersionOutputItem struct {
	DirPath string
	Version string
}

func (c cWeb) doUpgradeVersion(ctx context.Context, in cWebInput) (out *doUpgradeVersionOutput, err error) {
	mlog.Print(`start upgrading version...`)
	return
}

// doUpgradeCLI downloads the new version binary with process.
func (c cWeb) doUpgradeCLI(ctx context.Context) (err error) {
	mlog.Print(`start upgrading cli...`)
	defer func() {
		mlog.Printf(`new version cli binary is successfully installed to "%s"`, gfile.SelfPath())
	}()
	return
}

func (c cWeb) doAutoFixing(ctx context.Context, dirPath string, version string) (err error) {
	mlog.Printf(`auto fixing directory path "%s" from version "%s" ...`, dirPath, version)
	command := fmt.Sprintf(`gf fix -p %s`, dirPath)
	_ = gproc.ShellRun(ctx, command)
	return
}
