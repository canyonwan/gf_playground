package backend

import (
	"gf_playground/api/v1/common"
	"gf_playground/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type PermissionBase struct {
	Name string `json:"name" dc:"权限名称"`
	Path string `json:"path" dc:"权限路径"`
}

type PermissionCreateReq struct {
	g.Meta `path:"/permission" method:"POST" tags:"权限管理" summary:"创建"`
	PermissionBase
}
type PermissionCreateRes struct {
	Id uint `json:"id" dc:"权限ID"`
}

type PermissionDeleteReq struct {
	g.Meta `path:"/permission/{id}" method:"DELETE" tags:"权限管理" summary:"删除"`
	Id     uint `json:"id" v:"required#id必传" dc:"权限ID"`
}
type PermissionDeleteRes struct {
	Id uint `json:"id" dc:"权限ID"`
}

type PermissionUpdateReq struct {
	g.Meta `path:"/permission" method:"PUT" tags:"权限管理" summary:"更新"`
	Id     uint `json:"id" dc:"权限ID"`
	PermissionBase
}
type PermissionUpdateRes struct {
	Id uint `json:"id" dc:"权限ID"`
}

type PermissionPageReq struct {
	g.Meta `path:"/permission/page" method:"GET" tags:"权限管理" summary:"获取分页"`
	common.PageCommonReq
}
type PermissionPageRes struct {
	common.PageCommonRes
}

type PermissionListReq struct {
	g.Meta `path:"/permission/list" method:"GET" tags:"权限管理" summary:"获取列表"`
}
type PermissionListRes struct {
	List []entity.PermissionInfo `json:"list" dc:"权限列表"`
}
