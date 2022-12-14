// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"gf_playground/internal/dao/internal"
)

// internalDataStatisticsDao is internal type for wrapping internal DAO implements.
type internalDataStatisticsDao = *internal.DataStatisticsDao

// dataStatisticsDao is the data access object for table data_statistics.
// You can define custom methods on it to extend its functionality as you wish.
type dataStatisticsDao struct {
	internalDataStatisticsDao
}

var (
	// DataStatistics is globally public accessible object for table data_statistics operations.
	DataStatistics = dataStatisticsDao{
		internal.NewDataStatisticsDao(),
	}
)

// Fill with you ideas below.
