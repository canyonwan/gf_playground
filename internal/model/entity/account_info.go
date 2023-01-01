// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AccountInfo is the golang structure for table account_info.
type AccountInfo struct {
	Id           int         `json:"id"           description:"id"`       // id
	Account      string      `json:"account"      description:"帐号"`       // 帐号
	Password     string      `json:"password"     description:"密码"`       // 密码
	IsSuperAdmin int         `json:"isSuperAdmin" description:"是否是超级管理员"` // 是否是超级管理员
	RoleIds      string      `json:"roleIds"      description:"角色Ids"`    // 角色Ids
	UserSalt     string      `json:"userSalt"     description:"加密盐"`      // 加密盐
	CreatedAt    *gtime.Time `json:"createdAt"    description:"创建时间"`     // 创建时间
	UpdatedAt    *gtime.Time `json:"updatedAt"    description:"更新时间"`     // 更新时间
	DeletedAt    *gtime.Time `json:"deletedAt"    description:"删除时间"`     // 删除时间
}
