package position

import (
	"context"
	"gf_playground/api/v1/common"
	"gf_playground/internal/dao"
	"gf_playground/internal/model"
	"gf_playground/internal/model/entity"
	"gf_playground/internal/service"
)

type sPosition struct{}

func init() {
	service.RegisterPosition(&sPosition{})
}

func (sp *sPosition) PageGet(ctx context.Context, in model.PositionPageInput) (out *model.PositionPageOutput, err error) {
	var (
		m = dao.PositionInfo.Ctx(ctx)
	)
	// 1.先将当前查询的page & size 赋值返回
	out = &model.PositionPageOutput{
		PageCommonRes: common.PageCommonRes{
			Page: in.Page,
			Size: in.Size,
		},
	}
	// 2.分页查询
	listModel := m.Page(in.Page, in.Size)
	// 排序
	listModel = listModel.OrderDesc(dao.PositionInfo.Columns().Sort)
	// 2.1 查询当前page&size下的数据列表,单独存起来方便后面使用
	var list []*entity.PositionInfo
	if err := listModel.Scan(&list); err != nil {
		return out, err
	}
	if len(list) == 0 {
		return out, nil
	}
	out.Total, err = m.Count()
	if err != nil {
		return out, err
	}

	// 2.2获取到总数后, 将list列表渲染到结构体Content属性上
	if err := listModel.Scan(&out.Content); err != nil {

		return nil, err
	}
	return
}

func (sp *sPosition) Create(ctx context.Context, in model.PositionCreateInput) (out *model.PositionCreateOutput, err error) {
	id, err := dao.PositionInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return nil, err
	}
	return &model.PositionCreateOutput{Id: int(id)}, nil
}

func (sp *sPosition) Update(ctx context.Context, in model.PositionUpdateInput) (err error) {
	_, err = dao.PositionInfo.Ctx(ctx).Data(in).FieldsEx(dao.PositionInfo.Columns().Id).Where(dao.PositionInfo.Columns().Id, in.Id).Update()
	if err != nil {
		return err
	}
	return err
}
func (sp *sPosition) Delete(ctx context.Context, id int) (err error) {
	_, err = dao.PositionInfo.Ctx(ctx).WherePri(id).Unscoped().Delete()
	if err != nil {
		return err
	}
	return err
}
