package model

import (
	"gf_playground/api/v1/common"
	"gf_playground/internal/model/entity"
)

type AccountBase struct {
	Account      string `json:"account" dc:"帐号"`
	Password     string `json:"password" dc:"密码"`
	IsSuperAdmin int    `json:"isSuperAdmin" dc:"是否超级管理员"`
	RoleIds      string `json:"roleIds" dc:"角色IDs"`
}

// AccountPageInput 分页
type AccountPageInput struct {
	common.PageCommonReq
}
type AccountPageOutput struct {
	common.PageCommonRes
}

// AccountCreateInput 新增
type AccountCreateInput struct {
	AccountBase
}
type AccountCreateOutput struct {
	Id int `json:"id"`
}

// AccountDeleteInput 删除
type AccountDeleteInput struct {
	Id int `json:"id"`
}
type AccountDeleteOutput struct {
	Id int `json:"id"`
}

// AccountUpdateInput 更新
type AccountUpdateInput struct {
	entity.AccountInfo
}
type AccountUpdateOutput struct {
	Id int `json:"id"`
}
