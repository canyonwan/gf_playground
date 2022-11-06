package controller

import (
	"context"
	v1 "gf_playground/api/v1"
	"gf_playground/internal/model"
	"gf_playground/internal/model/entity"
	"gf_playground/internal/service"
	"time"
)

var (
	Banner = cBanner{}
)

type cBanner struct{}

// Create
// 功能： 接收并解析请求参数
func (c *cBanner) Create(ctx context.Context, req *v1.CreateBannerReq) (res *v1.CreateBannerRes, err error) {
	out, err := service.Banner().Create(ctx, model.BannerCreateInput{
		BannerBase: model.BannerBase{
			Url:      req.Url,
			JumpLink: req.JumpLink,
			Sort:     req.Sort,
		},
	})

	if err != nil {
		return res, nil
	}
	return &v1.CreateBannerRes{Id: out.Id}, nil
}

func GetSingle(ctx context.Context, req *v1.GetBannerReq) (res *v1.GetBannerRes, err error) {
	output, err := service.Banner().GetSingle(ctx, model.BannerGetInput{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return &v1.GetBannerRes{Banner: &entity.Banner{
		Id:        output.Id,
		Url:       output.Url,
		JumpLink:  output.JumpLink,
		Sort:      output.Sort,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	},
	}, nil
}
