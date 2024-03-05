package v1

import (
	"test_proj/internal/model"
	"test_proj/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type {{ .TableName }}SearchReq struct {
	g.Meta  `path:"/{{ .TableName }}/search" method:"get" summary:"查询用户列表(搜索分页)" tags:"用户管理"`
	Search  string `json:"search"   v:"length:0,99#搜索用户名|名长度为:{min}到:{max}位" dc:"用户编号"`
	OrderBy string `json:"order_by" v:"length:0,30#排序字段不正确"       d:"id"        dc:"排序字段"`
	Order   string `json:"order"    v:"in:desc,asc#排序方式不正确"       d:"order"     dc:"排序方式"`
	model.CommonPaginationReq
}
type {{ .tableName }}SearchRes struct {
	List    []*entity.{{ .tablename }} `json:"list"                          dc:"用户列表"` //
	OrderBy string         `json:"order_by"        d:"id"        dc:"排序字段不正确"`
	Desc    string         `json:"desc"            d:"desc"      dc:"排序方式不正确"`
	model.CommonPaginationRes
}
type UserCreateReq struct {
	g.Meta     `path:"/user/create" method:"post" tags:"用户管理" summary:"新建用户"`
	Name       string `json:"name"        v:"required|min-length:4"                dc:"用户名"`  //
	Email      string `json:"email"                                                dc:"邮箱"`   //
	Address    string `json:"address"                                              dc:"地址"`   //
	Mark       string `json:"mark"                                                 dc:"备注"`   //
	Permission string `json:"permission"                                           dc:"权限设置"` //
	Password   string `json:"password"    v:"required|ci|same:RePassword"          dc:"用户密码"` //
	RePassword string `json:"repassword"  v:"required"                             dc:"重复密码"` //
}
type UserCreateRes struct {
	NewID uint `json:"new_id"   dc:"ID"`
}
type UserResetPwdReq struct {
	g.Meta `path:"/user/reset_pwd/{id}" method:"post" tags:"用户管理" summary:"重置用户密码"`
	Id     uint `json:"id" in:"path"  v:"min:1#用户的编号不正确"                  dc:"用户编号"`
	// Password   string `json:"password"      v:"required|password|ci|same:RePassword" dc:"用户密码"` //
	Password   string `json:"password"      v:"required|same:RePassword" dc:"用户密码"`             //
	RePassword string `json:"repassword"    v:"required"                             dc:"重复密码"` //
}
type UserResetPwdRes struct {
}
type UserDeleteReq struct {
	g.Meta `path:"/user/delete/{id}" method:"post" tags:"用户管理" summary:"删除指定用户"`
	Id     uint `json:"id" in:"path"        dc:"用户编号"`
}
type UserDeleteRes struct {
}
type UserPermissionsReq struct {
	g.Meta      `path:"/user/permission/{id}" method:"post" tags:"用户管理" summary:"设置用户权限"`
	Id          uint   `json:"id" in:"path"        v:"min:1#用户的编号不正确"  dc:"用户编号"`
	Permissions string `json:"permissions"         v:"required|length:3,500" dc:"用户权限"`
}
type UserPermissionsRes struct {
}
type UserLoginReq struct {
	g.Meta   `path:"/login" method:"post" tags:"用户状态" summary:"登录"`
	UserName string `json:"username"        v:"required|min-length:4"    dc:"用户名"`  //
	Password string `json:"password"        v:"required|ci"              dc:"用户密码"` //
}
type UserLoginRes struct {
	User *model.UserLoginResponse `json:"user" dc:"用户信息"`
}

type UserLogoutReq struct {
	g.Meta `path:"/logout" method:"post" tags:"用户状态" summary:"登出"`
}
type UserLogoutRes struct {
}

type PermissionsAllReq struct {
	g.Meta `path:"/permissions/all" method:"get" tags:"权限记录" summary:"查询所有权限设置项"`
}
type PermissionsAllRes struct {
}
