package controller

import (
	"context"
	"gf_playground/api/v1/admin"
	"gf_playground/api/v1/common"
	"gf_playground/internal/model"
	"gf_playground/internal/model/entity"
	"gf_playground/internal/service"
)

var (
	Position = cPosition{}
)

type cPosition struct{}

// Page todo
func (sp *cPosition) Page(ctx context.Context, req *admin.PositionPageReq) (res *admin.PositionPageRes, err error) {
	out, err := service.Position().PageGet(ctx, model.PositionPageInput{
		PageCommonReq: common.PageCommonReq{
			Page: req.Page,
			Size: req.Size,
		},
		Sort: req.Sort,
	})
	if err != nil {
		return nil, err
	}
	return &admin.PositionPageRes{
		PageCommonRes: common.PageCommonRes{
			Content: out.Content,
			Page:    out.Page,
			Size:    out.Size,
			Total:   out.Total,
		},
	}, nil

}

func (sp *cPosition) Create(ctx context.Context, req *admin.PositionCreateReq) (res *admin.PositionCreateRes, err error) {
	out, err := service.Position().Create(ctx, model.PositionCreateInput{
		PositionBase: admin.PositionBase{
			GoodsName:    req.GoodsName,
			GoodsPicture: req.GoodsPicture,
			GoodsId:      req.GoodsId,
			LinkUrl:      req.LinkUrl,
			Sort:         req.Sort,
		},
	})
	if err != nil {
		return nil, err
	}
	return &admin.PositionCreateRes{Id: out.Id}, nil

}

func (sp *cPosition) Update(ctx context.Context, req *admin.PositionUpdateReq) (res *admin.PositionUpdateRes, err error) {
	err = service.Position().Update(ctx, model.PositionUpdateInput{
		Id: req.Id,
		PositionInfo: entity.PositionInfo{
			GoodsName:    req.GoodsName,
			GoodsPicture: req.GoodsPicture,
			GoodsId:      req.GoodsId,
			LinkUrl:      req.LinkUrl,
			Sort:         req.Sort,
		},
	})
	return &admin.PositionUpdateRes{Id: req.Id}, err
}

func (sp *cPosition) Delete(ctx context.Context, req *admin.PositionDeleteReq) (res *admin.PositionDeleteRes, err error) {
	err = service.Position().Delete(ctx, req.Id)
	return nil, err
}
