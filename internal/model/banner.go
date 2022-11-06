package model

type BannerCreateInput struct {
	Url      string
	JumpLink string
	Sort     int
}

type BannerCreateOutput struct {
	Id int
}
