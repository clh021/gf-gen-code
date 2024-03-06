// Copyright GoFrame gf Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package genapi

import (
	"context"
	"path/filepath"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"

	"github.com/gogf/gf/cmd/gf/v2/internal/consts"
	"github.com/gogf/gf/cmd/gf/v2/internal/utility/mlog"
)

func generateApi(ctx context.Context, in CGenDaoInternalInput) {

	var dirPathEntity = gfile.Join(in.Path, in.ApiPath)
	// Model content.
	for i, tableName := range in.TableNames {
		fieldMap, err := in.DB.TableFields(ctx, tableName)
		if err != nil {
			mlog.Fatalf("fetching tables fields failed for table '%s':\n%v", tableName, err)
		}

		var (
			newTableName                    = in.NewTableNames[i]
			entityFilePath                  = filepath.FromSlash(gfile.Join(dirPathEntity, gstr.CaseSnake(newTableName)+".go"))
			structDefinition, appendImports = generateStructDefinition(ctx, generateStructDefinitionInput{
				CGenDaoInternalInput: in,
				TableName:            tableName,
				StructName:           gstr.CaseCamel(newTableName),
				FieldMap:             fieldMap,
				IsDo:                 false,
			})
			_ = generateApiContent( // entityContent
				ctx,
				in,
				newTableName,
				gstr.CaseCamel(newTableName),
				structDefinition,
				appendImports,
			)
		)
		in.generatedFilePaths.EntityFilePaths = append(
			in.generatedFilePaths.EntityFilePaths,
			entityFilePath,
		)
		t := NewTpl()
		if err := t.Write(entityFilePath, "gen_templates/api.go.tpl", g.Map{
			"in":        in,
			"TableName": newTableName,
			// "tableNameCamelCase": gstr.CaseCamel(newTableName),
			// "structDefinition":   structDefinition,
			// "imports":            appendImports,
		}); err != nil {
			mlog.Fatal(ctx, err)
		}
		// 以下模板测试通过
		// {{ .in.Link }}
		// {{range $key, $value := .in.TypeMapping}}
		// - Mapping: {{$key}}, Type: {{$value.Type}}, Import: {{$value.Import}}
		// {{end}}
		// {range $filePath := .in.GeneratedFilePaths.EntityFilePaths}
		// - Generated File Path: {$filePath}
		// {end}
		// {{range $tableName := .in.TableNames}}
		// - Table Name: {{$tableName}}
		// {{end}}
		mlog.Print(ctx, "--------------------------------")
		mlog.Printf("generated: %s", entityFilePath)

		// err = gfile.PutContents(entityFilePath, strings.TrimSpace(entityContent))
		// if err != nil {
		// 	mlog.Fatalf("writing content to '%s' failed: %v", entityFilePath, err)
		// } else {
		// 	utils.GoFmt(entityFilePath)
		// 	mlog.Print("generated:", entityFilePath)
		// }
	}
}

func generateApiContent(
	ctx context.Context, in CGenDaoInternalInput, tableName, tableNameCamelCase, structDefine string, appendImports []string,
) string {
	entityContent := gstr.ReplaceByMap(
		getTemplateFromPathOrDefault(in.TplDaoEntityPath, consts.TemplateGenDaoEntityContent),
		g.MapStrStr{
			tplVarTableName:          tableName,
			tplVarPackageImports:     getImportPartContent(ctx, structDefine, false, appendImports),
			tplVarTableNameCamelCase: tableNameCamelCase,
			tplVarStructDefine:       structDefine,
		},
	)
	entityContent = replaceDefaultVar(in, entityContent)
	return entityContent
}
