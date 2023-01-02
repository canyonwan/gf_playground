package model

import (
	"gf_playground/api/v1/backend"
	"gf_playground/api/v1/common"
	"gf_playground/internal/model/entity"
)

type PositionPageInput struct {
	common.PageCommonReq
	Sort int `json:"sort" dc:"排序字段"`
}

type PositionPageOutput struct {
	common.PageCommonRes
}

type PositionCreateInput struct {
	backend.PositionBase
}

type PositionCreateOutput struct {
	Id int `json:"id"`
}

type PositionDeleteInput struct {
	Id int `json:"id"`
}
type PositionDeleteOutput struct{}

type PositionUpdateInput struct {
	Id int `json:"id"`
	entity.PositionInfo
}
type PositionUpdateOutput struct{}
