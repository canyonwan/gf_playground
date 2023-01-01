// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"gf_playground/internal/model"
)

type (
	IAccount interface {
		Page(ctx context.Context, in model.AccountPageInput) (out *model.AccountPageOutput, err error)
		Create(ctx context.Context, in model.AccountCreateInput) (out *model.AccountCreateOutput, err error)
		Update(ctx context.Context, in model.AccountUpdateInput) (out *model.AccountUpdateOutput, err error)
		Delete(ctx context.Context, in model.AccountDeleteInput) (out *model.AccountDeleteOutput, err error)
	}
)

var (
	localAccount IAccount
)

func Account() IAccount {
	if localAccount == nil {
		panic("implement not found for interface IAccount, forgot register?")
	}
	return localAccount
}

func RegisterAccount(i IAccount) {
	localAccount = i
}