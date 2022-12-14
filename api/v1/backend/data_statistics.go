package backend

import "github.com/gogf/gf/v2/frame/g"

// DataStatisticsGetReq 统计数据 统计图头部数据
type DataStatisticsGetReq struct {
	g.Meta `path:"/dataStatistics" method:"get" tags:"数据统计" summary:"详情"`
}

type DataStatisticsGetRes struct {
	TodayOrderCounts int `json:"todayOrderCounts" desc:"今日订单数"`
	DAU              int `json:"dau" desc:"日活"`
	ConversionRate   int `json:"conversionRate" desc:"转化率"`
}
