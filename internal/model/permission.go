package model

import (
	"gf_playground/api/v1/common"
	"gf_playground/internal/model/entity"
)

type PermissionBaseInput struct {
	Name string `json:"name" dc:"权限名称"`
	Path string `json:"path" dc:"权限路径"`
}

type PermissionCreateInput struct {
	PermissionBaseInput
}
type PermissionCreateOutput struct {
	Id uint `json:"id" dc:"权限ID"`
}

type PermissionDeleteOutput struct {
	Id uint `json:"id" dc:"权限ID"`
}

type PermissionUpdateInput struct {
	Id uint `json:"id" dc:"权限ID"`
	PermissionBaseInput
}
type PermissionUpdateOutput struct {
	Id uint `json:"id" dc:"权限ID"`
}

type PermissionPageInput struct {
	common.PageCommonReq
}
type PermissionPageOutput struct {
	common.PageCommonRes
}

// PermissionListInput todo
type PermissionListInput struct{}

type PermissionListOutput struct {
	List []entity.PermissionInfo `json:"list" dc:"权限列表"`
}
