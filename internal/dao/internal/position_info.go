// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PositionInfoDao is the data access object for table position_info.
type PositionInfoDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns PositionInfoColumns // columns contains all the column names of Table for convenient usage.
}

// PositionInfoColumns defines and stores column names for table position_info.
type PositionInfoColumns struct {
	Id           string //
	GoodsName    string // 商品名称
	GoodsPicture string // 商品主图
	GoodsId      string //
	LinkUrl      string // 跳转地址
	Sort         string //
	CreatedAt    string //
	UpdateedAt   string //
	DeletedAt    string //
}

// positionInfoColumns holds the columns for table position_info.
var positionInfoColumns = PositionInfoColumns{
	Id:           "id",
	GoodsName:    "goods_name",
	GoodsPicture: "goods_picture",
	GoodsId:      "goods_id",
	LinkUrl:      "link_url",
	Sort:         "sort",
	CreatedAt:    "created_at",
	UpdateedAt:   "updateed_at",
	DeletedAt:    "deleted_at",
}

// NewPositionInfoDao creates and returns a new DAO object for table data access.
func NewPositionInfoDao() *PositionInfoDao {
	return &PositionInfoDao{
		group:   "default",
		table:   "position_info",
		columns: positionInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PositionInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PositionInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PositionInfoDao) Columns() PositionInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PositionInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PositionInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PositionInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
