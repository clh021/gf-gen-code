package v1

import (
	"test_proj/internal/model"
	"test_proj/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

// 搜索
type {{ .data.TableName | ToCamel }}SearchReq struct {
	g.Meta  `path:"/{{ .data.TableName | ToSnake }}/search" method:"get" summary:"查询用户列表(搜索分页)" tags:"用户管理"`
	Search  string `json:"search"   v:"length:0,99#搜索用户名|名长度为:{min}到:{max}位" dc:"用户编号"`
	OrderBy string `json:"order_by" v:"length:0,30#排序字段不正确"       d:"id"        dc:"排序字段"`
	Order   string `json:"order"    v:"in:desc,asc#排序方式不正确"       d:"order"     dc:"排序方式"`
	model.CommonPaginationReq
}
type {{ .data.TableName | ToCamel }}SearchRes struct {
	List    []*entity.{{ .data.TableName | ToCamel }} `json:"list"                                              dc:"用户列表"` //
	OrderBy string         `json:"order_by"                            d:"id"        dc:"排序字段不正确"`
	Desc    string         `json:"desc"                                d:"desc"      dc:"排序方式不正确"`
	model.CommonPaginationRes
}

// 创建
type {{ .data.TableName | ToCamel }}CreateReq struct {
	g.Meta     `path:"/{{ .data.TableName | ToSnake }}/create" method:"post" tags:"用户管理" summary:"新建用户"`{{range $cname, $c := .data.Fields }}{{if eq $c.Name "created_at" "updated_at" "deleted_at"}}{{else}}
	{{$c.Name | ToCamel }} {{$c.GType}} `json:{{$c.JsonTag | ToPadAndQuote}}                               dc:"{{$c.Comment}}"`{{end}}{{end}}
}
type {{ .data.TableName | ToCamel }}CreateRes struct {
	NewID uint `json:"new_id"   dc:"ID"`
}

// 查一个
type {{ .data.TableName | ToCamel }}OneReq struct {
	g.Meta `path:"/{{ .data.TableName | ToSnake }}/one/{id}" method:"get" tags:"用户管理" summary:"获取一个"`
	Id     uint `json:"id" in:"path"  v:"min:1#编号不正确"                  dc:"编号"`
}
type {{ .data.TableName | ToCamel }}OneRes struct {
	*entity.{{ .data.TableName | ToCamel }}
}

// 修改
type {{ .data.TableName | ToCamel }}UpdateReq struct {
	g.Meta `path:"/{{ .data.TableName | ToSnake }}/update/{id}" method:"post" tags:"用户管理" summary:"重置密码"`
	Id     uint `json:"id" in:"path"  v:"min:1#编号不正确"                   dc:"编号"`{{range $cname, $c := .data.Fields }}{{if eq $c.Name "id" "created_at" "updated_at" "deleted_at"}}{{else}}
	{{$c.Name | ToCamel }} {{$c.GType}} `json:{{$c.JsonTag | ToPadAndQuote}}                               dc:"{{$c.Comment}}"`{{end}}{{end}}
}
type {{ .data.TableName | ToCamel }}UpdateRes struct {
}

// 删除
type {{ .data.TableName | ToCamel }}DeleteReq struct {
	g.Meta `path:"/{{ .data.TableName | ToSnake }}/delete/{id}" method:"post" tags:"用户管理" summary:"删除指定用户"`
	Id     uint `json:"id" in:"path"        dc:"用户编号"`
}
type {{ .data.TableName | ToCamel }}DeleteRes struct {
}
