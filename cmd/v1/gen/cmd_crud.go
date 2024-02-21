package gen

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gtag"
)

const (
	CGenCrudConfig = `gfcli.gen.crud`
	CGenCrudUsage  = `gf gen crud [OPTION]`
	CGenCrudBrief  = `parse api definitions to generate controller/sdk go files`
	CGenCrudEg     = `
gf gen crud
`
	CGenCrudBriefSrcFolder     = `source folder path to be parsed. default: api`
	CGenCrudBriefDstFolder     = `destination folder path storing automatically generated go files. default: internal/controller`
	CGenCrudBriefWatchFile     = `used in file watcher, it re-generates go files only if given file is under srcFolder`
	CGenCrudBriefSdkPath       = `also generate SDK go files for api definitions to specified directory`
	CGenCrudBriefSdkStdVersion = `use standard version prefix for generated sdk request path`
	CGenCrudBriefSdkNoV1       = `do not add version suffix for interface module name if version is v1`
	CGenCrudBriefClear         = `auto delete generated and unimplemented controller go files if api definitions are missing`
	CGenCrudControllerMerge    = `generate all controller files into one go file by name of api definition source go file`
)

const (
	PatternApiDefinition  = `type[\s\(]+(\w+)Req\s+struct\s+{([\s\S]+?)}`
	PatternCrudDefinition = `func\s+\(.+?\)\s+\w+\(.+?\*(\w+)\.(\w+)Req\)\s+\(.+?\*(\w+)\.(\w+)Res,\s+\w+\s+error\)\s+{`
)

const (
	genCrudFileLockSeconds = 10
)

func init() {
	gtag.Sets(g.MapStrStr{
		`CGenCrudConfig`:             CGenCrudConfig,
		`CGenCrudUsage`:              CGenCrudUsage,
		`CGenCrudBrief`:              CGenCrudBrief,
		`CGenCrudEg`:                 CGenCrudEg,
		`CGenCrudBriefSrcFolder`:     CGenCrudBriefSrcFolder,
		`CGenCrudBriefDstFolder`:     CGenCrudBriefDstFolder,
		`CGenCrudBriefWatchFile`:     CGenCrudBriefWatchFile,
		`CGenCrudBriefSdkPath`:       CGenCrudBriefSdkPath,
		`CGenCrudBriefSdkStdVersion`: CGenCrudBriefSdkStdVersion,
		`CGenCrudBriefSdkNoV1`:       CGenCrudBriefSdkNoV1,
		`CGenCrudBriefClear`:         CGenCrudBriefClear,
		`CGenCrudControllerMerge`:    CGenCrudControllerMerge,
	})
}

type (
	CGenCrud      struct{}
	CGenCrudInput struct {
		g.Meta          `name:"crud" config:"{CGenCrudConfig}" usage:"{CGenCrudUsage}" brief:"{CGenCrudBrief}" eg:"{CGenCrudEg}"`
		SrcFolder       string   `short:"s" name:"srcFolder"     brief:"{CGenCrudBriefSrcFolder}" d:"api"`
		DstFolder       string   `short:"d" name:"dstFolder"     brief:"{CGenCrudBriefDstFolder}" d:"internal/controller"`
		DstFileNameCase string   `short:"f" name:"dstFileNameCase" brief:"{CGenServiceBriefFileNameCase}" d:"Snake"`
		WatchFile       string   `short:"w" name:"watchFile"     brief:"{CGenCrudBriefWatchFile}"`
		StPattern       string   `short:"a" name:"stPattern" brief:"{CGenServiceBriefStPattern}" d:"^s([A-Z]\\w+)$"`
		Packages        []string `short:"p" name:"packages" brief:"{CGenServiceBriefPackages}"`
		ImportPrefix    string   `short:"i" name:"importPrefix" brief:"{CGenServiceBriefImportPrefix}"`
		SdkPath         string   `short:"k" name:"sdkPath"       brief:"{CGenCrudBriefSdkPath}"`
		SdkStdVersion   bool     `short:"v" name:"sdkStdVersion" brief:"{CGenCrudBriefSdkStdVersion}" orphan:"true"`
		SdkNoV1         bool     `short:"n" name:"sdkNoV1"       brief:"{CGenCrudBriefSdkNoV1}" orphan:"true"`
		Clear           bool     `short:"c" name:"clear"         brief:"{CGenCrudBriefClear}" orphan:"true"`
		Merge           bool     `short:"m" name:"merge"         brief:"{CGenCrudControllerMerge}" orphan:"true"`
	}
	CGenCrudOutput struct{}
)

func (c CGenCrud) Crud(ctx context.Context, in CGenCrudInput) (out *CGenCrudOutput, err error) {

	// TODO: 分析 sqlite 数据库
	// 得出一个数据库中所有数据表结构和 `符合规则的注释`

	// 1. 根据配置，连接一个 sqlite 数据库
	// 要读取用户当前目录，识别是否是 gf 项目，符合条件才继续运行
	// 读取 gf 项目的 hack 配置文件
	// 2. 得到数据库中的所有表
	// 3. 分析每一个表的表名，表注释
	// 4. 分析每一个表的字段名，字段类型，字段注释
	// 5. 生成 api 接口文件
	// 6. 生成 dao 文件
	// 7. 生成 logic 文件
	// 8. 生成 service 文件
	// 9. 生成 ctrl 文件
	c.genCtrl(ctx, in)
	c.genService(ctx, in)
	return
}
func (c CGenCrud) genCtrl(ctx context.Context, in CGenCrudInput) {
	// var (
	// 	_in = genctrl.CGenCtrlInput{
	// 		SrcFolder:     in.SrcFolder,
	// 		DstFolder:     in.DstFolder,
	// 		WatchFile:     "",
	// 		SdkPath:       "",
	// 		SdkStdVersion: false,
	// 		SdkNoV1:       false,
	// 		Clear:         false,
	// 		Merge:         false,
	// 	}
	// )
	// _, err := genctrl.CGenCtrl{}.Ctrl(ctx, _in)
	// if err != nil {
	// 	panic(err)
	// }
}
func (c CGenCrud) genService(ctx context.Context, in CGenCrudInput) {
	// var (
	// 	_in = genservice.CGenServiceInput{
	// 		SrcFolder:       in.SrcFolder,
	// 		DstFolder:       in.DstFolder,
	// 		DstFileNameCase: "Snake",
	// 		WatchFile:       "",
	// 		StPattern:       "",
	// 		Packages:        nil,
	// 		ImportPrefix:    "",
	// 		Clear:           false,
	// 	}
	// )
	// _, err := genservice.CGenService{}.Service(ctx, _in)
	// if err != nil {
	// 	panic(err)
	// }
}