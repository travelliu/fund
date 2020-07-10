// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package fund

import (
	"context"
	_fundMod "github.com/travelliu/fund/services/fund/models"
)

//go:generate mockery -name=Repository

// Repository 数据库操作
type Repository interface {
	fundRepo
	userFundRepo
}

type fundRepo interface {
	CreateFund(ctx context.Context, fund *_fundMod.Fund) error
	QueryFundByCode(ctx context.Context, code string) (*_fundMod.Fund, error)
	QueryFundByUserID(ctx context.Context, userID int64) ([]*_fundMod.Fund, error)
	QueryAllFund(ctx context.Context) ([]*_fundMod.Fund, error)
	UpdateFund(ctx context.Context, fund *_fundMod.Fund, ) error
}

type userFundRepo interface {
	CreateUserFund(ctx context.Context, fund *_fundMod.UserFund) error
	QueryUserFundByUserID(ctx context.Context, id int64) ([]*_fundMod.UserFund, error)
	QueryUserFundByCode(ctx context.Context, userID int64, code string) (*_fundMod.UserFund, error)
	EditUserFund(ctx context.Context, userFund *_fundMod.UserFund) error
	DeleteUserFundByCode(ctx context.Context, userID int64, code string) error
}
