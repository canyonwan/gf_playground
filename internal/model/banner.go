package model

// BannerBase banner基本类型
type BannerBase struct {
	Url      string
	JumpLink string
	Sort     int
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
