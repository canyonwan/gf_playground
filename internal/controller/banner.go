package controller

import (
	"context"
	"gf_playground/api/v1/app"
	"gf_playground/internal/model"
	"gf_playground/internal/model/entity"
	"gf_playground/internal/service"
)

// ================================================================================
// 1.将用户传入的json(req)传给service层
// 2.在logic进入service层的依赖注入, 进行业务逻辑的操作
// 3.其中在logic层调用dao层与数据库进行通信获取数据,并返回获取到的数据
// ================================================================================

var (
	Banner = cBanner{}
)

type cBanner struct{}

// Create
// 功能： 接收并解析请求参数
func (c *cBanner) Create(ctx context.Context, req *app.CreateBannerReq) (res *app.CreateBannerRes, err error) {
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
	return &app.CreateBannerRes{Id: out.Id}, nil
}

func (c *cBanner) GetSingle(ctx context.Context, req *app.GetBannerReq) (res *app.GetBannerRes, err error) {
	output, err := service.Banner().GetSingle(ctx, model.BannerGetInput{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return &app.GetBannerRes{Banner: &entity.Banner{
		Id:        output.Id,
		Url:       output.Url,
		JumpLink:  output.JumpLink,
		Sort:      output.Sort,
		CreatedAt: &output.CreatedAt,
		UpdatedAt: &output.UpdatedAt,
		DeletedAt: &output.DeletedAt,
	},
	}, nil
}

func (c *cBanner) DeleteSingle(ctx context.Context, req *app.DeleteBannerReq) (res *app.DeleteBannerRes, err error) {
	err = service.Banner().DeleteSingle(ctx, req.Id)
	return
}
