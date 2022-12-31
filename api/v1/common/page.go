package common

import (
	"gf_playground/internal/model/entity"
)

// PageCommonReq 分页
type PageCommonReq struct {
	Page int `json:"page" d:"1" v:"min:0#分页页码不能小于1" dc:"页码,默认第1页"`
	Size int `json:"size" d:"10" v:"max:50#分页数量不能小于50条" dc:"页数,默认10条"`
}

type PageCommonRes struct {
	//Content interface{} `json:"content" dc:"列表"`
	Content []entity.Todo `json:"content" dc:"列表"`
	Page    int           `json:"page" dc:"页码"`
	Size    int           `json:"size" dc:"页数"`
	Total   int           `json:"total" dc:"总数"`
}
