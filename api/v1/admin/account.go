package admin

import (
	"gf_playground/api/v1/common"
	"gf_playground/internal/model/entity"
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
	g.Meta `path:"admin/account" method:"get" tags:"帐号管理"  summary:"分页"`
	common.PageCommonReq
}
type AccountPageRes struct {
	common.PageCommonRes
}

// AccountCreateReq 新增
type AccountCreateReq struct {
	g.Meta `path:"admin/account" method:"POST" tags:"帐号管理" summary:"创建管理员" `
	AccountBase
}
type AccountCreateRes struct {
	Id int `json:"id"`
}

// AccountUpdateReq 编辑
type AccountUpdateReq struct {
	g.Meta `path:"admin/account" method:"put" tags:"帐号管理" summary:"编辑管理员" `
	*entity.AccountInfo
}
type AccountUpdateRes struct {
	Id int `json:"id"`
}

// AccountDeleteReq 删除
type AccountDeleteReq struct {
	g.Meta `path:"admin/account/{id}" in:"path" method:"delete" tags:"帐号管理" summary:"删除管理员" `
	Id     int `json:"id" dc:"管理员Id"`
}
type AccountDeleteRes struct {
	Id int `json:"id" dc:"管理员Id"`
}
