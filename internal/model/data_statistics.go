package model

type DataStatisticsGetOutput struct {
	TodayOrderCounts int `json:"todayOrderCounts" desc:"今日订单数"`
	DAU              int `json:"dau" desc:"日活"`
	ConversionRate   int `json:"conversionRate" desc:"转化率"`
}
