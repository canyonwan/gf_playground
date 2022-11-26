package controller

import (
	"context"
	v1 "gf_playground/api/v1"
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

func (c *cBanner) GetSingle(ctx context.Context, req *v1.GetBannerReq) (res *v1.GetBannerRes, err error) {
	output, err := service.Banner().GetSingle(ctx, model.BannerGetInput{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return &v1.GetBannerRes{Banner: &entity.Banner{
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

func (c *cBanner) DeleteSingle(ctx context.Context, req *v1.DeleteBannerReq) (res *v1.DeleteBannerRes, err error) {
	_, _ = service.Banner().DeleteSingle(ctx, model.BannerGetInput{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return nil, nil

}
