package backend

import (
	"gf_playground/api/v1/common"
	"github.com/gogf/gf/v2/frame/g"
)

type PositionBase struct {
	GoodsName    string `json:"goods_name" v:"required#商品名称不能为空" dc:"商品名称"`
	GoodsPicture string `json:"goods_picture" v:"required#商品图片不能为空" dc:"商品图片"`
	GoodsId      int    `json:"goods_id" v:"required#商品id不能为空" dc:"商品id"`
	LinkUrl      string `json:"link_url" dc:"跳转地址"`
	Sort         int    `json:"sort" dc:"排序"`
}

// PositionPageReq 分页
type PositionPageReq struct {
	g.Meta `path:"/position" method:"GET" tags:"手工位管理"  summary:"分页"`
	Page   int `json:"page" dc:"页码" d:"1"`
	Size   int `json:"size" dc:"页数" d:"10"`
	Sort   int `json:"sort" dc:"排序字段"`
}
type PositionPageRes struct {
	common.PageCommonRes
}

// PositionCreateReq 新增
type PositionCreateReq struct {
	g.Meta `path:"/position" method:"POST" tags:"手工位管理" summary:"新增"`
	PositionBase
}
type PositionCreateRes struct {
	Id int `json:"id"`
}

// PositionDeleteReq 删除
type PositionDeleteReq struct {
	g.Meta `path:"/position/{id}" in:"path" method:"delete" tags:"手工位管理" summary:"删除"`
	Id     int `json:"id" v:"required#id不能为空" dc:"id"`
}
type PositionDeleteRes struct{}

// PositionUpdateReq 更新
type PositionUpdateReq struct {
	g.Meta `path:"/position" method:"PUT" tags:"手工位管理" summary:"更新"`
	Id     int `json:"id" v:"required#id不能为空" dc:"id"`
	PositionBase
}

type PositionUpdateRes struct {
	Id int `json:"id"`
}
