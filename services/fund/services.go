// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package fund

import (
	"context"
	_fundMod "github.com/travelliu/fund/services/fund/models"
)

//go:generate mockery -name=UseCase

// UseCase 用户场景
type UseCase interface {
	fundUc
	userFundUc
}

type fundUc interface {
	CreateFund(ctx context.Context, fund *_fundMod.Fund) error
	QueryFundByCode(ctx context.Context, code string) (*_fundMod.Fund, error)
	FundSync(ctx context.Context, workerNum int) error
}

type userFundUc interface {
	CreateUserFund(ctx context.Context, userFund *_fundMod.UserFund) (*_fundMod.UserFund, error) // 添加或者修改
	QueryUserFundByUserID(ctx context.Context, id int64) (*_fundMod.UserFundResponse, error)
	QueryUserFundByCode(ctx context.Context, userID int64, code string) (*_fundMod.UserFund, error)
	DeleteUserFundByCode(ctx context.Context, userID int64, code string) error
}
