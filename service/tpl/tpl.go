package tpl

import (
	"context"

	_ "embed"

	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gview"
)

//go:embed api.tpl
var tplFileApi []byte

type tpl struct {
	_ctx    context.Context
	gv      *gview.View
	tplPath string
}

func New(tplPath string) *tpl {
	_ctx := context.TODO()
	gv := gview.New()
	return &tpl{_ctx, gv, tplPath}
}

func (t *tpl) DefaultFile() string {
	return t.gv.GetDefaultFile()
}

func (t *tpl) Parse(tplName string, params ...map[string]interface{}) (c string, e error) {
	tplContent := string(tplFileApi)
	if gfile.IsFile(tplName) {
		tplContent = gfile.GetContents(tplName)
	}
	return t.gv.ParseContent(t._ctx, tplContent, params...)

}
