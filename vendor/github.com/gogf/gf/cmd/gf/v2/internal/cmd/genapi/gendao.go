// Copyright GoFrame gf Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package genapi

import (
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gtag"
)

const (
	CGenDaoConfig = `gfcli.gen.dao`
	CGenDaoUsage  = `gf gen api [OPTION]`
	CGenDaoBrief  = `automatically generate go files for dao/do/entity`
	CGenDaoEg     = `
gf gen api
gf gen api -l "mysql:root:12345678@tcp(127.0.0.1:3306)/test"
gf gen api -p ./model -g user-center -t user,user_detail,user_login
gf gen api -r user_
`

	CGenDaoAd = `
CONFIGURATION SUPPORT
    Options are also supported by configuration file.
    It's suggested using configuration file instead of command line arguments making producing.
    The configuration node name is "gfcli.gen.dao", which also supports multiple databases, for example(config.yaml):
	gfcli:
	  gen:
		dao:
		- link:     "mysql:root:12345678@tcp(127.0.0.1:3306)/test"
		  tables:   "order,products"
		  jsonCase: "CamelLower"
		- link:   "mysql:root:12345678@tcp(127.0.0.1:3306)/primary"
		  path:   "./my-app"
		  prefix: "primary_"
		  tables: "user, userDetail"
		  typeMapping:
			decimal:
			  type:   decimal.Decimal
			  import: github.com/shopspring/decimal
			numeric:
			  type: string
`
	CGenDaoBriefPath              = `directory path for generated files`
	CGenDaoBriefLink              = `database configuration, the same as the ORM configuration of GoFrame`
	CGenDaoBriefTables            = `generate models only for given tables, multiple table names separated with ','`
	CGenDaoBriefTablesEx          = `generate models excluding given tables, multiple table names separated with ','`
	CGenDaoBriefPrefix            = `add prefix for all table of specified link/database tables`
	CGenDaoBriefRemovePrefix      = `remove specified prefix of the table, multiple prefix separated with ','`
	CGenDaoBriefRemoveFieldPrefix = `remove specified prefix of the field, multiple prefix separated with ','`
	CGenDaoBriefStdTime           = `use time.Time from stdlib instead of gtime.Time for generated time/date fields of tables`
	CGenDaoBriefWithTime          = `add created time for auto produced go files`
	CGenDaoBriefGJsonSupport      = `use gJsonSupport to use *gjson.Json instead of string for generated json fields of tables`
	CGenDaoBriefImportPrefix      = `custom import prefix for generated go files`
	CGenDaoBriefDaoPath           = `directory path for storing generated dao files under path`
	CGenDaoBriefDoPath            = `directory path for storing generated do files under path`
	CGenDaoBriefEntityPath        = `directory path for storing generated entity files under path`
	CGenDaoBriefApiPath           = `directory path for storing generated api files under path`
	CGenDaoBriefOverwriteDao      = `overwrite all dao files both inside/outside internal folder`
	CGenDaoBriefModelFile         = `custom file name for storing generated model content`
	CGenDaoBriefModelFileForDao   = `custom file name generating model for DAO operations like Where/Data. It's empty in default`
	CGenDaoBriefDescriptionTag    = `add comment to description tag for each field`
	CGenDaoBriefNoJsonTag         = `no json tag will be added for each field`
	CGenDaoBriefNoModelComment    = `no model comment will be added for each field`
	CGenDaoBriefClear             = `delete all generated go files that do not exist in database`
	CGenDaoBriefTypeMapping       = `custom local type mapping for generated struct attributes relevant to fields of table`
	CGenDaoBriefGroup             = `
specifying the configuration group name of database for generated ORM instance,
it's not necessary and the default value is "default"
`
	CGenDaoBriefJsonCase = `
generated json tag case for model struct, cases are as follows:
| Case            | Example            |
|---------------- |--------------------|
| Camel           | AnyKindOfString    |
| CamelLower      | anyKindOfString    | default
| Snake           | any_kind_of_string |
| SnakeScreaming  | ANY_KIND_OF_STRING |
| SnakeFirstUpper | rgb_code_md5       |
| Kebab           | any-kind-of-string |
| KebabScreaming  | ANY-KIND-OF-STRING |
`
	CGenDaoBriefTplDaoIndexPath    = `template file path for dao index file`
	CGenDaoBriefTplDaoInternalPath = `template file path for dao internal file`
	CGenDaoBriefTplDaoDoPathPath   = `template file path for dao do file`
	CGenDaoBriefTplDaoEntityPath   = `template file path for dao entity file`
	CGenDaoBriefTplDaoApiPath   = `template file path for dao entity file`

	tplVarTableName               = `{TplTableName}`
	tplVarTableNameCamelCase      = `{TplTableNameCamelCase}`
	tplVarTableNameCamelLowerCase = `{TplTableNameCamelLowerCase}`
	tplVarPackageImports          = `{TplPackageImports}`
	tplVarImportPrefix            = `{TplImportPrefix}`
	tplVarStructDefine            = `{TplStructDefine}`
	tplVarColumnDefine            = `{TplColumnDefine}`
	tplVarColumnNames             = `{TplColumnNames}`
	tplVarGroupName               = `{TplGroupName}`
	tplVarDatetimeStr             = `{TplDatetimeStr}`
	tplVarCreatedAtDatetimeStr    = `{TplCreatedAtDatetimeStr}`
)

var (
	createdAt          = gtime.Now()
	defaultTypeMapping = map[DBFieldTypeName]CustomAttributeType{
		"decimal": {
			Type: "float64",
		},
		"money": {
			Type: "float64",
		},
		"numeric": {
			Type: "float64",
		},
		"smallmoney": {
			Type: "float64",
		},
	}
)

func init() {
	gtag.Sets(g.MapStrStr{
		`CGenDaoiConfig`:                  CGenDaoConfig,
		`CGenDaoiUsage`:                   CGenDaoUsage,
		`CGenDaoiBrief`:                   CGenDaoBrief,
		`CGenDaoiEg`:                      CGenDaoEg,
		`CGenDaoiAd`:                      CGenDaoAd,
		`CGenDaoiBriefPath`:               CGenDaoBriefPath,
		`CGenDaoiBriefLink`:               CGenDaoBriefLink,
		`CGenDaoiBriefTables`:             CGenDaoBriefTables,
		`CGenDaoiBriefTablesEx`:           CGenDaoBriefTablesEx,
		`CGenDaoiBriefPrefix`:             CGenDaoBriefPrefix,
		`CGenDaoiBriefRemovePrefix`:       CGenDaoBriefRemovePrefix,
		`CGenDaoiBriefRemoveFieldPrefix`:  CGenDaoBriefRemoveFieldPrefix,
		`CGenDaoiBriefStdTime`:            CGenDaoBriefStdTime,
		`CGenDaoiBriefWithTime`:           CGenDaoBriefWithTime,
		`CGenDaoiBriefDaoPath`:            CGenDaoBriefDaoPath,
		`CGenDaoiBriefDoPath`:             CGenDaoBriefDoPath,
		`CGenDaoiBriefEntityPath`:         CGenDaoBriefEntityPath,
		`CGenDaoiBriefApiPath`:            CGenDaoBriefApiPath,
		`CGenDaoiBriefGJsonSupport`:       CGenDaoBriefGJsonSupport,
		`CGenDaoiBriefImportPrefix`:       CGenDaoBriefImportPrefix,
		`CGenDaoiBriefOverwriteDao`:       CGenDaoBriefOverwriteDao,
		`CGenDaoiBriefModelFile`:          CGenDaoBriefModelFile,
		`CGenDaoiBriefModelFileForDao`:    CGenDaoBriefModelFileForDao,
		`CGenDaoiBriefDescriptionTag`:     CGenDaoBriefDescriptionTag,
		`CGenDaoiBriefNoJsonTag`:          CGenDaoBriefNoJsonTag,
		`CGenDaoiBriefNoModelComment`:     CGenDaoBriefNoModelComment,
		`CGenDaoiBriefClear`:              CGenDaoBriefClear,
		`CGenDaoiBriefTypeMapping`:        CGenDaoBriefTypeMapping,
		`CGenDaoiBriefGroup`:              CGenDaoBriefGroup,
		`CGenDaoiBriefJsonCase`:           CGenDaoBriefJsonCase,
		`CGenDaoiBriefTplDaoIndexPath`:    CGenDaoBriefTplDaoIndexPath,
		`CGenDaoiBriefTplDaoInternalPath`: CGenDaoBriefTplDaoInternalPath,
		`CGenDaoiBriefTplDaoDoPathPath`:   CGenDaoBriefTplDaoDoPathPath,
		`CGenDaoiBriefTplDaoEntityPath`:   CGenDaoBriefTplDaoEntityPath,
		`CGenDaoiBriefTplDaoApiPath`:      CGenDaoBriefTplDaoApiPath,
	})
}

type (
	CGenDao      struct{}
	CGenDaoInput struct {
		g.Meta             `name:"api" config:"{CGenDaoiConfig}" usage:"{CGenDaoiUsage}" brief:"{CGenDaoBrief}" eg:"{CGenDaoiEg}" ad:"{CGenDaoAd}"`
		Path               string `name:"path"                short:"p"  brief:"{CGenDaoBriefPath}" d:"api"`
		Link               string `name:"link"                short:"l"  brief:"{CGenDaoBriefLink}"`
		Tables             string `name:"tables"              short:"t"  brief:"{CGenDaoBriefTables}"`
		TablesEx           string `name:"tablesEx"            short:"x"  brief:"{CGenDaoBriefTablesEx}"`
		Group              string `name:"group"               short:"g"  brief:"{CGenDaoBriefGroup}" d:"default"`
		Prefix             string `name:"prefix"              short:"f"  brief:"{CGenDaoBriefPrefix}"`
		RemovePrefix       string `name:"removePrefix"        short:"r"  brief:"{CGenDaoBriefRemovePrefix}"`
		RemoveFieldPrefix  string `name:"removeFieldPrefix"   short:"rf" brief:"{CGenDaoBriefRemoveFieldPrefix}"`
		JsonCase           string `name:"jsonCase"            short:"j"  brief:"{CGenDaoBriefJsonCase}" d:"CamelLower"`
		ImportPrefix       string `name:"importPrefix"        short:"i"  brief:"{CGenDaoBriefImportPrefix}"`
		DaoPath            string `name:"daoPath"             short:"d"  brief:"{CGenDaoBriefDaoPath}" d:"dao"`
		DoPath             string `name:"doPath"              short:"o"  brief:"{CGenDaoBriefDoPath}" d:"model/do"`
		EntityPath         string `name:"entityPath"          short:"e"  brief:"{CGenDaoBriefEntityPath}" d:"model/entity"`
		ApiPath            string `name:"apiPath"             short:"api" brief:"{CGenDaoiBriefApiPath}" d:"hello/v1"`
		TplDaoIndexPath    string `name:"tplDaoIndexPath"     short:"t1" brief:"{CGenDaoBriefTplDaoIndexPath}"`
		TplDaoInternalPath string `name:"tplDaoInternalPath"  short:"t2" brief:"{CGenDaoBriefTplDaoInternalPath}"`
		TplDaoDoPath       string `name:"tplDaoDoPath"        short:"t3" brief:"{CGenDaoBriefTplDaoDoPathPath}"`
		TplDaoEntityPath   string `name:"tplDaoEntityPath"    short:"t4" brief:"{CGenDaoBriefTplDaoEntityPath}"`
		TplDaoApiPath      string `name:"tplDaoApiPath"       short:"t5" brief:"{CGenDaoiBriefTplDaoApiPath}"`
		StdTime            bool   `name:"stdTime"             short:"s"  brief:"{CGenDaoBriefStdTime}" orphan:"true"`
		WithTime           bool   `name:"withTime"            short:"w"  brief:"{CGenDaoBriefWithTime}" orphan:"true"`
		GJsonSupport       bool   `name:"gJsonSupport"        short:"n"  brief:"{CGenDaoBriefGJsonSupport}" orphan:"true"`
		OverwriteDao       bool   `name:"overwriteDao"        short:"v"  brief:"{CGenDaoBriefOverwriteDao}" orphan:"true"`
		DescriptionTag     bool   `name:"descriptionTag"      short:"c"  brief:"{CGenDaoBriefDescriptionTag}" orphan:"true"`
		NoJsonTag          bool   `name:"noJsonTag"           short:"k"  brief:"{CGenDaoBriefNoJsonTag}" orphan:"true"`
		NoModelComment     bool   `name:"noModelComment"      short:"m"  brief:"{CGenDaoBriefNoModelComment}" orphan:"true"`
		Clear              bool   `name:"clear"               short:"a"  brief:"{CGenDaoBriefClear}" orphan:"true"`

		TypeMapping        map[DBFieldTypeName]CustomAttributeType `name:"typeMapping" short:"y" brief:"{CGenDaoBriefTypeMapping}" orphan:"true"`
		generatedFilePaths *CGenDaoInternalGeneratedFilePaths
	}
	CGenDaoOutput struct{}

	CGenDaoInternalInput struct {
		CGenDaoInput
		DB            gdb.DB
		TableNames    []string
		NewTableNames []string
	}

	CGenDaoInternalGeneratedFilePaths struct {
		DaoFilePaths         []string
		DaoInternalFilePaths []string
		DoFilePaths          []string
		EntityFilePaths      []string
		ApiFilePaths         []string
	}

	DBFieldTypeName     = string
	CustomAttributeType struct {
		Type   string `brief:"custom attribute type name"`
		Import string `brief:"custom import for this type"`
	}
)