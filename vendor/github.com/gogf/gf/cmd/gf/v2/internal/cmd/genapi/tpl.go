package genapi

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gres"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/os/gview"
)

type tpl struct {
	_ctx    context.Context
	Gv      *gview.View
}

func NewTpl() *tpl {
	_ctx := context.TODO()
	gv := gview.New()
	return &tpl{_ctx, gv}
}

func (t *tpl) Parse(tplName string, params ...map[string]interface{}) (c string, e error) {
	glog.Debug(t._ctx, "params : ", params)
	tplContent := string("")
	tplMode := ""
	if gfile.IsFile(tplName) {
		tplMode = "file"
		tplContent = gfile.GetContents(tplName)
	} else if gres.Contains(tplName) {
		tplMode = "gres"
		tplContent = string(gres.GetContent(tplName))
	} else {
		glog.Fatal(t._ctx, "tplName: ", tplName, " not found.")
	}
	glog.Info(t._ctx, "tplName: ", tplName, "(", tplMode, ")")
	return t.Gv.ParseContent(t._ctx, tplContent, params...)
}

func (t *tpl) Write(dstPath, tplName string, params ...map[string]interface{}) error {
	content, err := t.Parse(tplName, params...)
	if err != nil {
		return err
	}
	glog.Debug(t._ctx, content)
	glog.Debug(t._ctx, "write to", dstPath)
	return gfile.PutContents(dstPath, content)
}

func (t *tpl) TempFile() string {
	path := fmt.Sprintf(`%s/%d`, gfile.Temp(), gtime.TimestampNano())
	return fmt.Sprintf(`%s/%s`, path, "t.tpl")
}
