// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package repository

import (
	"context"
	_fundMod "github.com/travelliu/fund/services/fund/models"
	"github.com/travelliu/fund/utils/trace"
	"strings"
)

// CreateFund 添加基金信息
func (r *repo) CreateOrUpdateFundHistory(ctx context.Context, fundHistory *_fundMod.FundHistory) error {
	fh, err := r.QueryFundHistory(ctx, fundHistory.Code, fundHistory.EquityDate)
	if err != nil {
		logger.WithField(string(trace.ContextKeyReqID), trace.GetReqID(ctx)).Errorf("the QueryFundHistory error %s", err)
		return err
	}
	if fh.ID == 0 {
		err = r.db.Create(fundHistory).Error
	} else {
		if fh.Equity == fundHistory.Equity {
			return nil
		}
		fundHistory.ID = fh.ID
		err = r.db.Model(&_fundMod.FundHistory{}).Update(fundHistory).Error
	}

	return err
}

// QueryFundHistory 查询基金历史一天记录
func (r *repo) QueryFundHistory(ctx context.Context, code, date string) (*_fundMod.FundHistory, error) {
	var (
		fundHistory = &_fundMod.FundHistory{}
	)
	err := r.db.Where("code = ? and equity_date", code, date).Find(fundHistory).Error
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return fundHistory, nil
		}
		return nil, err
	}
	return fundHistory, err
}
