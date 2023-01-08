package frontend

import (
	"context"
	"gf_playground/api/v1/frontend"
	"gf_playground/internal/model"
	"gf_playground/internal/service"
)

var (
	BannerFrontend = cBanner{}
)

type cBanner struct{}

// GetList 前端获取
func (c *cBanner) GetList(ctx context.Context, req *frontend.BannerListReq) (res *frontend.BannerListRes, err error) {
	list, err := service.Banner().GetList(ctx)
	if err != nil {
		return nil, err
	}
	return &frontend.BannerListRes{
		BannerListOutput: model.BannerListOutput{List: list.List},
	}, nil
}
