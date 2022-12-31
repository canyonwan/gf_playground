// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PositionInfo is the golang structure for table position_info.
type PositionInfo struct {
	Id           int         `json:"id"           description:""`     //
	GoodsName    string      `json:"goodsName"    description:"商品名称"` // 商品名称
	GoodsPicture string      `json:"goodsPicture" description:"商品主图"` // 商品主图
	GoodsId      int         `json:"goodsId"      description:""`     //
	LinkUrl      string      `json:"linkUrl"      description:"跳转地址"` // 跳转地址
	Sort         int         `json:"sort"         description:""`     //
	CreatedAt    *gtime.Time `json:"createdAt"    description:""`     //
	UpdatedAt    *gtime.Time `json:"updatedAt"    description:""`     //
	DeletedAt    *gtime.Time `json:"deletedAt"    description:""`     //
}
