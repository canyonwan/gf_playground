package backend

import (
	"gf_playground/api/v1/common"
	"github.com/gogf/gf/v2/frame/g"
)

type AccountBase struct {
	Account      string `json:"account" dc:"帐号"`
	Password     string `json:"password" dc:"密码"`
	IsSuperAdmin int    `json:"isSuperAdmin" dc:"是否超级管理员"`
	RoleIds      string `json:"roleIds" dc:"角色IDs"`
}

// AccountPageReq 分页
type AccountPageReq struct {
	g.Meta `path:"/account" method:"get" tags:"帐号管理" summary:"分页"`
	common.PageCommonReq
}
type AccountPageRes struct {
	common.PageCommonRes
}

// AccountCreateReq 新增
type AccountCreateReq struct {
	g.Meta `path:"/account" method:"POST" tags:"帐号管理" summary:"创建" `
	AccountBase
}
type AccountCreateRes struct {
	Id int `json:"id"`
}

// AccountUpdateReq 编辑
type AccountUpdateReq struct {
	g.Meta `path:"/account" method:"put" tags:"帐号管理" summary:"编辑"`
	Id     int `json:"id"`
	AccountBase
}
type AccountUpdateRes struct {
	Id int `json:"id"`
}

// AccountDeleteReq 删除
type AccountDeleteReq struct {
	g.Meta `path:"/account/{id}" in:"path" method:"delete" tags:"帐号管理" summary:"删除" `
	Id     int `json:"id" dc:"管理员Id"`
}
type AccountDeleteRes struct {
	Id int `json:"id" dc:"管理员Id"`
}

// AccountInfoReq 获取帐号信息的接口
type AccountInfoReq struct {
	g.Meta `path:"/account/info/{id}" method:"get" tags:"帐号管理" summary:"帐号详情"`
	Id     int `json:"id" dc:"帐号Id"`
}
type AccountInfoRes struct {
	Id           int    `json:"id" dc:"帐号Id"`
	Account      string `json:"account" dc:"帐号"`
	IsSuperAdmin int    `json:"isSuperAdmin" dc:"是否超级管理员"`
	RoleIds      string `json:"roleIds" dc:"角色Ids"`
}
