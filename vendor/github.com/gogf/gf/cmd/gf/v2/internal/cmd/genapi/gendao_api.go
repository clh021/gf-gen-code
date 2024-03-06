// Copyright GoFrame gf Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package genapi

import (
	"context"
	"path/filepath"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"

	"github.com/gogf/gf/cmd/gf/v2/internal/consts"
	"github.com/gogf/gf/cmd/gf/v2/internal/utility/mlog"
	"github.com/gogf/gf/cmd/gf/v2/internal/utility/utils"
)


func (c CGenDao) Api(ctx context.Context, in CGenDaoInput) (out *CGenDaoOutput, err error) {
	in.generatedFilePaths = &CGenDaoInternalGeneratedFilePaths{
		DaoFilePaths:         make([]string, 0),
		DaoInternalFilePaths: make([]string, 0),
		DoFilePaths:          make([]string, 0),
		EntityFilePaths:      make([]string, 0),
	}
	if g.Cfg().Available(ctx) {
		v := g.Cfg().MustGet(ctx, CGenDaoConfig)
		if v.IsSlice() {
			for i := 0; i < len(v.Interfaces()); i++ {
				doGenDaoForArray(ctx, i, in)
			}
		} else {
			doGenDaoForArray(ctx, -1, in)
		}
	} else {
		doGenDaoForArray(ctx, -1, in)
	}
	mlog.Print("done!")
	return
}

func generateApi(ctx context.Context, in CGenDaoInternalInput) {
	var (
		dirPathDao         = gfile.Join(in.Path, in.DaoPath)
		dirPathDaoInternal = gfile.Join(dirPathDao, "internal")
	)
	for i := 0; i < len(in.TableNames); i++ {
		generateApiSingle(ctx, generateDaoSingleInput{
			CGenDaoInternalInput: in,
			TableName:            in.TableNames[i],
			NewTableName:         in.NewTableNames[i],
			DirPathDao:           dirPathDao,
			DirPathDaoInternal:   dirPathDaoInternal,
		})
	}
}

// generateApiSingle generates the dao and model content of given table.
func generateApiSingle(ctx context.Context, in generateDaoSingleInput) {
	// Generating table data preparing.
	fieldMap, err := in.DB.TableFields(ctx, in.TableName)
	if err != nil {
		mlog.Fatalf(`fetching tables fields failed for table "%s": %+v`, in.TableName, err)
	}
	var (
		tableNameCamelCase      = gstr.CaseCamel(in.NewTableName)
		tableNameCamelLowerCase = gstr.CaseCamelLower(in.NewTableName)
		tableNameSnakeCase      = gstr.CaseSnake(in.NewTableName)
		importPrefix            = in.ImportPrefix
	)
	if importPrefix == "" {
		importPrefix = utils.GetImportPath(gfile.Join(in.Path, in.DaoPath))
	} else {
		importPrefix = gstr.Join(g.SliceStr{importPrefix, in.DaoPath}, "/")
	}

	fileName := gstr.Trim(tableNameSnakeCase, "-_.")
	if len(fileName) > 5 && fileName[len(fileName)-5:] == "_test" {
		// Add suffix to avoid the table name which contains "_test",
		// which would make the go file a testing file.
		fileName += "_table"
	}

	// dao - index
	generateDaoIndex(generateDaoIndexInput{
		generateDaoSingleInput:  in,
		TableNameCamelCase:      tableNameCamelCase,
		TableNameCamelLowerCase: tableNameCamelLowerCase,
		ImportPrefix:            importPrefix,
		FileName:                fileName,
	})

	// dao - internal
	generateDaoInternal(generateDaoInternalInput{
		generateDaoSingleInput:  in,
		TableNameCamelCase:      tableNameCamelCase,
		TableNameCamelLowerCase: tableNameCamelLowerCase,
		ImportPrefix:            importPrefix,
		FileName:                fileName,
		FieldMap:                fieldMap,
	})
}

func generateApiIndex(in generateDaoIndexInput) {
	path := filepath.FromSlash(gfile.Join(in.DirPathDao, in.FileName+".go"))
	// It should add path to result slice whenever it would generate the path file or not.
	in.generatedFilePaths.DaoFilePaths = append(
		in.generatedFilePaths.DaoFilePaths,
		path,
	)
	if in.OverwriteDao || !gfile.Exists(path) {
		indexContent := gstr.ReplaceByMap(
			getTemplateFromPathOrDefault(in.TplDaoIndexPath, consts.TemplateGenDaoIndexContent),
			g.MapStrStr{
				tplVarImportPrefix:            in.ImportPrefix,
				tplVarTableName:               in.TableName,
				tplVarTableNameCamelCase:      in.TableNameCamelCase,
				tplVarTableNameCamelLowerCase: in.TableNameCamelLowerCase,
			})
		indexContent = replaceDefaultVar(in.CGenDaoInternalInput, indexContent)
		if err := gfile.PutContents(path, strings.TrimSpace(indexContent)); err != nil {
			mlog.Fatalf("writing content to '%s' failed: %v", path, err)
		} else {
			utils.GoFmt(path)
			mlog.Print("generated:", path)
		}
	}
}

func generateApiInternal(in generateDaoInternalInput) {
	path := filepath.FromSlash(gfile.Join(in.DirPathDaoInternal, in.FileName+".go"))
	removeFieldPrefixArray := gstr.SplitAndTrim(in.RemoveFieldPrefix, ",")
	modelContent := gstr.ReplaceByMap(
		getTemplateFromPathOrDefault(in.TplDaoInternalPath, consts.TemplateGenDaoInternalContent),
		g.MapStrStr{
			tplVarImportPrefix:            in.ImportPrefix,
			tplVarTableName:               in.TableName,
			tplVarGroupName:               in.Group,
			tplVarTableNameCamelCase:      in.TableNameCamelCase,
			tplVarTableNameCamelLowerCase: in.TableNameCamelLowerCase,
			tplVarColumnDefine:            gstr.Trim(generateColumnDefinitionForDao(in.FieldMap, removeFieldPrefixArray)),
			tplVarColumnNames:             gstr.Trim(generateColumnNamesForDao(in.FieldMap, removeFieldPrefixArray)),
		})
	modelContent = replaceDefaultVar(in.CGenDaoInternalInput, modelContent)
	in.generatedFilePaths.DaoInternalFilePaths = append(
		in.generatedFilePaths.DaoInternalFilePaths,
		path,
	)
	if err := gfile.PutContents(path, strings.TrimSpace(modelContent)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		utils.GoFmt(path)
		mlog.Print("generated:", path)
	}
}