// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// DataStatistics is the golang structure for table data_statistics.go.
type DataStatistics struct {
	Id               int `json:"id"               description:"id"`     // id
	TodayOrderCounts int `json:"todayOrderCounts" description:"今日订单总数"` // 今日订单总数
	DAU              int `json:"dAU"              description:"日活"`     // 日活
	ConversionRate   int `json:"conversionRate"   description:"转化率"`    // 转化率
}
