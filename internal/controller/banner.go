package controller

import (
	"context"
	v1 "gf_playground/api/v1"
	"gf_playground/internal/model"
	"gf_playground/internal/service"
)

var (
	Banner = cBanner{}
)

type cBanner struct{}

// Create
// 功能： 接收并解析请求参数
func (c *cBanner) Create(ctx context.Context, req *v1.CreateBannerReq) (res *v1.CreateBannerRes, err error) {
	out, err := service.Banner().Create(ctx, model.BannerCreateInput{
		Url:      req.Url,
		JumpLink: req.JumpLink,
		Sort:     req.Sort,
	})
	if err != nil {
		return res, nil
	}
	return &v1.CreateBannerRes{Id: out.Id}, nil
}
