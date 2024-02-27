// Copyright GoFrame gf Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package genlogic

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
	Logic = cLogic{}
)

type cLogic struct {
	g.Meta `name:"logic" brief:"genereate logic defined go file" eg:"{cLogicEg}" `
}

const (
	cLogicEg     = `
gf logic
gf logic -a
gf logic -c
gf logic -cf
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
			mlog.Print()
			mlog.Print(`Done! logic defined go file has been generated.`)
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

func (c cLogic) doUpgradeVersion(ctx context.Context, in cLogicInput) (out *doUpgradeVersionOutput, err error) {
	mlog.Print(`start upgrading version...`)
	return
}

// doUpgradeCLI downloads the new version binary with process.
func (c cLogic) doUpgradeCLI(ctx context.Context) (err error) {
	mlog.Print(`start upgrading cli...`)
	defer func() {
		mlog.Printf(`new version cli binary is successfully installed to "%s"`, gfile.SelfPath())
	}()
	return
}

func (c cLogic) doAutoFixing(ctx context.Context, dirPath string, version string) (err error) {
	mlog.Printf(`auto fixing directory path "%s" from version "%s" ...`, dirPath, version)
	command := fmt.Sprintf(`gf fix -p %s`, dirPath)
	_ = gproc.ShellRun(ctx, command)
	return
}
