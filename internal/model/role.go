package model

import "gf_playground/api/v1/common"

type RoleBase struct {
	RoleName string `json:"roleName" dc:"角色名称"`
	RoleDesc string `json:"roleDesc" dc:"角色描述"`
}

// RoleCreateInput 创建角色
type RoleCreateInput struct {
	RoleBase
}
type RoleCreateOutput struct {
	Id int `json:"id" dc:"角色ID"`
}

// RoleUpdateInput 编辑角色
type RoleUpdateInput struct {
	Id int `json:"id" dc:"角色ID"`
	RoleBase
}
type RoleUpdateOutput struct {
	Id int `json:"id" dc:"角色ID"`
}

// RoleDeleteInput 删除角色
type RoleDeleteInput struct {
	Id int `json:"id" dc:"角色ID"`
}
type RoleDeleteOutput struct {
	Id int `json:"id" dc:"角色ID"`
}

// RolePageInput 角色分页
type RolePageInput struct {
	common.PageCommonReq
}
type RolePageOutput struct {
	common.PageCommonRes
}
