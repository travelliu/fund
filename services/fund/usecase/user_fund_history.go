// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package usecase

import (
	"context"
	_fundMod "github.com/travelliu/fund/services/fund/models"
	"github.com/travelliu/fund/utils/trace"
)

func (f *fund) replaceUserFundHistory(ctx context.Context, fund *_fundMod.Fund) error {
	var (
		userFunds, err    = f.fundRepo.QueryUserFundByCode(ctx, fund.Code)
		userFundHistories = []*_fundMod.UserFundHistory{}
	)
	if err != nil {
		logger.WithField(string(trace.ContextKeyReqID), trace.GetReqID(ctx)).Errorf("the QueryFundByCode error %s", err)
		return err
	}
	for _, uf := range userFunds {
		calcUserFund(uf, fund)
		uf.ID = 0
		userFundHistories = append(userFundHistories, &_fundMod.UserFundHistory{
			UserFund: *uf,
		})
	}
	if err := f.fundRepo.ReplaceUserFundHistory(ctx, fund.Code, userFundHistories); err != nil {
		logger.WithField(string(trace.ContextKeyReqID), trace.GetReqID(ctx)).Errorf("the QueryFundByCode error %s", err)
		return err
	}
	return nil
}
