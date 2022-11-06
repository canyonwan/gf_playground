package banner

import (
	"context"
	"gf_playground/internal/dao"
	"gf_playground/internal/model"
	"gf_playground/internal/service"
)

// 全局对象
type sBanner struct{}

// 自动调用的方法
// 在里面进行服务注册
func init() {
	service.RegisterBanner(New())
}

func New() *sBanner {
	return &sBanner{}
}

// Create 为sBanner添加方法
func (s *sBanner) Create(ctx context.Context, in model.BannerCreateInput) (out model.BannerCreateOutput, err error) {
	id, err := dao.Banner.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.BannerCreateOutput{Id: int(id)}, nil
}
