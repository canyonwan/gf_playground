package v1

import "github.com/gogf/gf/v2/frame/g"

// CreateBannerReq 新增轮播图
type CreateBannerReq struct {
	g.Meta   `path:"/banner" method:"post" tags:"轮播图" summary:"创建Banner"`
	Url      string `json:"url" in:"body" v:"required|#轮播图地址必填" dc:"轮播图地址"`
	JumpLink string `json:"jumpLink" in:"body" v:"required|#跳转链接必填" dc:"跳转链接"`
	Sort     int    `json:"sort" in:"body" dc:"排序"`
}

// CreateBannerRes 新增轮播图
type CreateBannerRes struct {
	Id int `json:"id" dc:"轮播图id"`
}
