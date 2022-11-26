package model

import "github.com/gogf/gf/v2/os/gtime"

// BannerBase banner基本类型
type BannerBase struct {
	Url       string
	JumpLink  string
	Sort      int
	CreatedAt gtime.Time
	UpdatedAt gtime.Time
	DeletedAt gtime.Time
}

// BannerCreateInput 新增banner
type BannerCreateInput struct {
	BannerBase
}

type BannerCreateOutput struct {
	Id int
}

// BannerGetInput 获取单个banner
type BannerGetInput struct {
	Id int
}

type BannerGetOutput struct {
	Id int
	BannerBase
}
