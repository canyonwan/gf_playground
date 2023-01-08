package backend

import (
	"gf_playground/api/v1/common"
	"github.com/gogf/gf/v2/frame/g"
)

type RoleCreateReq struct {
	g.Meta   `path:"/role" method:"POST" tags:"角色管理" summary:"创建"`
	RoleName string `json:"roleName" v:"required#角色名称不能为空" dc:"角色名称"`
	RoleDesc string `json:"roleDesc" dc:"角色描述"`
}

type RoleCreateRes struct {
	Id int `json:"id" dc:"角色ID"`
}

// RoleUpdateReq 编辑角色
type RoleUpdateReq struct {
	g.Meta   `path:"/role" method:"PUT" tags:"角色管理" summary:"编辑"`
	Id       int    `json:"id" v:"required#角色ID不能为空" dc:"角色ID"`
	RoleName string `json:"roleName" v:"required#角色名称不能为空" dc:"角色名称"`
	RoleDesc string `json:"roleDesc" dc:"角色描述"`
}
type RoleUpdateRes struct {
	Id int `json:"id" dc:"角色ID"`
}

// RoleDeleteReq 删除角色
type RoleDeleteReq struct {
	g.Meta `path:"/role/{id}" method:"DELETE" tags:"角色管理" summary:"删除"`
	Id     int `json:"id" v:"required#角色ID不能为空" dc:"角色ID"`
}
type RoleDeleteRes struct {
	Id int `json:"id" dc:"角色ID"`
}

// RolePageReq 角色分页
type RolePageReq struct {
	g.Meta `path:"/role" method:"GET" tags:"角色管理" summary:"分页"`
	common.PageCommonReq
}
type RolePageRes struct {
	common.PageCommonRes
}
