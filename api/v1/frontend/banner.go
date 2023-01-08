package frontend

import (
	"gf_playground/internal/model"
	"github.com/gogf/gf/v2/frame/g"
)

type BannerListReq struct {
	g.Meta `path:"/banner/list" method:"GET" tags:"App端-轮播图" summary:"列表" `
}
type BannerListRes struct {
	model.BannerListOutput
}
