package model

import "github.com/gogf/gf/v2/os/gtime"

// BannerBase banner基本类型
type BannerBase struct {
	Url       string     `json:"url"`
	JumpLink  string     `json:"jumpLink"`
	Sort      int        `json:"sort"`
	CreatedAt gtime.Time `json:"createdAt"`
	UpdatedAt gtime.Time `json:"updatedAt"`
	DeletedAt gtime.Time `json:"deletedAt"`
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
type BannerListOutput struct {
	List []BannerGetOutput `json:"list"`
}
