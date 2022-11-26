package banner

import (
	"context"
	"gf_playground/internal/dao"
	"gf_playground/internal/model"
	"gf_playground/internal/service"
)

// 全局对象
type sBanner struct{}

// 会在service层引入并自动调用该方法
// 在里面进行服务注册
func init() {
	// 将logic层当前的全局对象地址传给service层的注册方法进行注册
	// 注: 注册方法(RegisterBanner)接收的是个IBanner的类型的接口
	// - IBanner接口:为任何具有Create,GetSingle方法的类型,所有定义了Create,GetSingle这些方法的类型,我们称: 都实现了IBanner接口
	// - 该接口里定义了2个方法,如果New()方法返回的结构体实现了注册方法(RegisterBanner)的接收的IBanner 接口类型
	// - 那么我们说该结构体实现了该接口(IBanner)
	// - 小知识: 接口定义了对象的行为, 具体的行为由对象决定
	// 疑问: 这里的依赖注入, 是谁注入了谁啊
	service.RegisterBanner(&sBanner{})
}

// New()方法其实就是实例化对象
//func New() *sBanner {
//	return &sBanner{}
//}

// Create 为sBanner添加方法
func (s *sBanner) Create(ctx context.Context, in model.BannerCreateInput) (out *model.BannerCreateOutput, err error) {
	id, err := dao.Banner.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return &model.BannerCreateOutput{Id: int(id)}, nil
}

// GetSingle 为sBanner添加方法
// 在logic层调用dao对象获取数据库里的数据并返回
func (s *sBanner) GetSingle(ctx context.Context, in model.BannerGetInput) (out *model.BannerGetOutput, err error) {
	var output = &model.BannerGetOutput{}
	err = dao.Banner.Ctx(ctx).WherePri(in.Id).Scan(output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (s *sBanner) DeleteSingle(ctx context.Context, in model.BannerGetInput) (out *model.BannerCreateOutput, err error) {
	_, _ = dao.Banner.Ctx(ctx).WherePri(in.Id).Delete()
	if err != nil {
		return nil, err
	}
	return nil, nil
}
