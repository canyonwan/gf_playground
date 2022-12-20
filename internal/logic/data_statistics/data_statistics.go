package data_statistics

import (
	"context"
	"gf_playground/internal/dao"
	"gf_playground/internal/model"
	"gf_playground/internal/service"
	"gf_playground/utility"
	"github.com/gogf/gf/v2/os/gtime"
)

// 创建一个单例
type sDataStatistics struct{}

// 注册服务
func init() {
	service.RegisterDataStatistics(&sDataStatistics{})
}

//func New() *sDataStatistics {
//	return &sDataStatistics{}
//}

// DataStatistics Get 获取数据统计
func (s *sDataStatistics) DataStatistics(ctx context.Context) (out *model.DataStatisticsGetOutput, err error) {
	return &model.DataStatisticsGetOutput{
		TodayOrderCounts: queryTodayOrderCounts(ctx),
		DAU:              utility.RandInt(1000),
		ConversionRate:   utility.RandInt(50),
	}, nil
}

// 私有方法: 查询今天的订单数
func queryTodayOrderCounts(ctx context.Context) (count int) {
	// 本来是要写订单数的(OrderInfo表),但我还没有写, 所以这里先查商品表的个数吧
	count, err := dao.Goods.Ctx(ctx).WhereBetween(dao.Goods.Columns().CreateAt, gtime.Now().StartOfDay(), gtime.Now().EndOfDay()).Count(dao.Goods.Columns().Id)
	if err != nil {
		return -1
	}

	return count
}
