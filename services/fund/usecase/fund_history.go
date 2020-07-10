// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package usecase

import (
	"context"
	_fundMod "github.com/travelliu/fund/services/fund/models"
)

func (f *fund) addFundHistory(ctx context.Context, fund *_fundMod.Fund) error {
	fundHistory := &_fundMod.FundHistory{
		Fund: *fund,
	}
	return f.fundRepo.CreateOrUpdateFundHistory(ctx, fundHistory)
}
