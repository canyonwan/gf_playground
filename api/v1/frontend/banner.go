package frontend

import (
	"gf_playground/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

// CreateBannerReq 新增轮播图
type CreateBannerReq struct {
	g.Meta   `path:"frontend/banner" method:"post" tags:"轮播图" summary:"新增"`
	Url      string `json:"url" v:"required#播图地址必填" dc:"轮播图地址"`
	JumpLink string `json:"jumpLink" v:"required#跳转链接必填" dc:"跳转链接"`
	Sort     int    `json:"sort" dc:"排序"`
}

// CreateBannerRes 新增轮播图
type CreateBannerRes struct {
	Id int `json:"id" dc:"轮播图id"`
}

// GetBannerReq 获取单个轮播图
type GetBannerReq struct {
	g.Meta `path:"frontend/banner/{id}" method:"get" tags:"轮播图" summary:"获取单个"`
	Id     int `json:"id" v:"required#轮播ID必传" dc:"轮播图主键ID"`
}

type GetBannerRes struct {
	*entity.Banner
}

type DeleteBannerReq struct {
	g.Meta `path:"frontend/banner/{id}" method:"delete" tags:"轮播图" summary:"删除单个"`
	Id     int `json:"id" v:"required#轮播ID必传" dc:"轮播图主键ID"`
}

type DeleteBannerRes struct{}
