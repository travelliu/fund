// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package usecase

import (
	"context"
	_fundMod "github.com/travelliu/fund/services/fund/models"
)

// CreateFund 创建基金
func (f *fund) CreateFund(ctx context.Context, fund *_fundMod.Fund) error {

	return nil
}

// QueryFundByCode 通过code查询基金
func (f *fund) QueryFundByCode(ctx context.Context, code string) (*_fundMod.Fund, error) {

	return nil, nil
}
