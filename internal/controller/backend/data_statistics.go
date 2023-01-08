package backend

import (
	"context"
	"gf_playground/api/v1/backend"
	"gf_playground/internal/service"
)

var DataStatistics = cDataStatistics{}

type cDataStatistics struct{}

func (c *cDataStatistics) GetDataStatistics(ctx context.Context, req *backend.DataStatisticsGetReq) (res *backend.DataStatisticsGetRes, err error) {
	// 将前端传过来的json参数 -> service -> logic进行dao层查询
	statistics, err := service.DataStatistics().DataStatistics(ctx)
	if err != nil {
		return nil, err
	}
	return &backend.DataStatisticsGetRes{
		TodayOrderCounts: statistics.TodayOrderCounts,
		DAU:              statistics.DAU,
		ConversionRate:   statistics.ConversionRate,
	}, nil

}
