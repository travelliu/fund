// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package usecase

import (
	"context"
	_fundMod "github.com/travelliu/fund/services/fund/models"
	"github.com/travelliu/fund/services/fund/utils"
	_utils "github.com/travelliu/fund/utils"
	"github.com/travelliu/fund/utils/trace"
)

// CreateUserFund 用户添加或者修改基金
func (f *fund) CreateUserFund(ctx context.Context, userFund *_fundMod.UserFund) (*_fundMod.UserFund, error) {
	// 检查基金是否存在,没有则获取添加
	var (
		fund  *_fundMod.Fund
		err   error
		uFund *_fundMod.UserFund
	)
	if uFund, err = f.fundRepo.QueryUserFundByUserIDAndCode(ctx, userFund.UserID, userFund.Code); err != nil {
		logger.WithField(string(trace.ContextKeyReqID), trace.GetReqID(ctx)).Errorf("the QueryFundByCode error %s", err)
		return nil, err
	}

	if fund, err = f.fundRepo.QueryFundByCode(ctx, userFund.Code); err != nil {
		logger.WithField(string(trace.ContextKeyReqID), trace.GetReqID(ctx)).Errorf("the QueryFundByCode error %s", err)
		return nil, err
	}
	if fund.ID == 0 {
		if fund, err = utils.GetFundInfo(userFund.Code); err != nil {
			logger.WithField(string(trace.ContextKeyReqID), trace.GetReqID(ctx)).Errorf("the GetFundInfo error %s", err)
			return nil, err
		}
		if err := f.fundRepo.CreateFund(ctx, fund); err != nil {
			logger.WithField(string(trace.ContextKeyReqID), trace.GetReqID(ctx)).Errorf("the CreateFund error %s", err)
			return nil, err
		}
	}

	calcUserFund(userFund, fund)
	if uFund.ID == 0 {
		// 添加基金
		err = f.fundRepo.CreateUserFund(ctx, userFund)
	} else {
		// 修改基金
		userFund.ID = uFund.ID
		userFund.Code = uFund.Code
		userFund.UserID = uFund.UserID
		err = f.fundRepo.EditUserFund(ctx, userFund)
	}

	if err != nil {
		logger.WithField(string(trace.ContextKeyReqID), trace.GetReqID(ctx)).Errorf("the CreateFund error %s", err)
		return nil, err
	}
	return f.fundRepo.QueryUserFundByUserIDAndCode(ctx, userFund.UserID, userFund.Code)
}

// QueryUserFundByUserID 通过code查询基金
func (f *fund) QueryUserFundByUserID(ctx context.Context, userID int64) (*_fundMod.UserFundResponse, error) {
	var (
		userFundResponseList         = []*_fundMod.UserFundResponseList{}
		CostAmount           float64 = 0 // 持仓金额
		CostEquityAmount     float64 = 0 // 净值持仓金额
		CostValuationAmount  float64 = 0 // 估值持仓金额. 估值计算
		TotalEquity          float64 = 0 // 总收益
		TodayValuation       float64 = 0 // 今日收益估值
		TodayEquity          float64 = 0 // 今昨日收益估值
		// YesterdayEarnings    float64 = 0 // 昨日收益
		// TodayEarnings        float64 = 0 // 今日估算收益
		// TotalEarningsPre     float64 = 0 // 总收益率
		// YesterdayEarningsPre float64 = 0 // 昨日收益率
		// TodayEarningsPre     float64 = 0 // 今日估算收益率
	)
	funds, err := f.fundRepo.QueryFundByUserID(ctx, userID)
	if err != nil {
		logger.WithField(string(trace.ContextKeyReqID), trace.GetReqID(ctx)).Errorf("the QueryFundByUserID error %s", err)
		return nil, err
	}
	userFunds, err := f.fundRepo.QueryUserFundByUserID(ctx, userID)
	if err != nil {
		logger.WithField(string(trace.ContextKeyReqID), trace.GetReqID(ctx)).Errorf("the QueryUserFundByUserID error %s", err)
		return nil, err
	}

	for _, uf := range userFunds {
		for _, f := range funds {
			if uf.Code != f.Code {
				continue
			}
			calcUserFund(uf, f)
			userFundResponseList = append(userFundResponseList, &_fundMod.UserFundResponseList{
				FundBase: f.FundBase,
				UserFund: uf,
				CostPre:  0,
			})
		}
		CostAmount = CostAmount + uf.CostAmount
		CostEquityAmount = CostEquityAmount + uf.CostEquityAmount
		CostValuationAmount = CostValuationAmount + uf.CostValuationAmount
		TotalEquity = TotalEquity + uf.TotalEquity
		TodayValuation = TodayValuation + uf.TodayValuation
		TodayEquity = TodayEquity + uf.TodayEquity
	}

	for _, uf := range userFundResponseList {
		costPre := _utils.CalcFloat64(uf.CostAmount/CostAmount*100, 2)
		uf.CostPre = costPre
	}
	return &_fundMod.UserFundResponse{
		List:                userFundResponseList,
		CostAmount:          _utils.CalcFloat64(CostAmount, 2),
		CostEquityAmount:    _utils.CalcFloat64(CostEquityAmount, 2),
		CostValuationAmount: _utils.CalcFloat64(CostValuationAmount, 2),
		TotalEquity:         _utils.CalcFloat64(TotalEquity, 2),
		TotalEquityYield:    _utils.CalcFloat64(TotalEquity/CostAmount*100, 2),
		TodayValuation:      _utils.CalcFloat64(TodayValuation, 2),
		TodayEquity:         _utils.CalcFloat64(TodayEquity, 2),
	}, nil
}

// QueryUserFundByUserIDAndCode 查询用户基金code
func (f *fund) QueryUserFundByUserIDAndCode(ctx context.Context, userID int64, code string) (*_fundMod.UserFund, error) {
	return f.fundRepo.QueryUserFundByUserIDAndCode(ctx, userID, code)
}

func calcUserFund(userFund *_fundMod.UserFund, fund *_fundMod.Fund) {
	netDifference := _utils.CalcFloat64(fund.Equity-userFund.CostPrice, 4)
	valuationDifference := _utils.CalcFloat64(fund.Valuation-userFund.CostPrice, 4)
	// 昨日净值
	sellingPer := float64(1) + float64(userFund.SellingPer)/100                  // 卖出价格. 成本价*(1+百分比)
	PurchasePer := float64(1) - float64(userFund.PurchasePer)/100                // 补仓价格. 最后净值*(1-百分比)
	userFund.SellingPrice = _utils.CalcFloat64(userFund.CostPrice*sellingPer, 4) // 卖出价格. 成本价*(1+百分比)
	userFund.PurchasePrice = _utils.CalcFloat64(fund.Equity*PurchasePer, 4)      // 补仓价格. 最后净值*(1-百分比)

	userFund.CostAmount = _utils.CalcFloat64(userFund.Shares*userFund.CostPrice, 2) // 持仓金额
	userFund.CostEquityAmount = _utils.CalcFloat64(userFund.Shares*fund.Equity, 2)  // 净值持仓金额

	userFund.TotalEquity = _utils.CalcFloat64(netDifference*userFund.Shares, 2)                     // 净值总收益
	userFund.TotalEquityYield = _utils.CalcFloat64(userFund.TotalEquity/userFund.CostAmount*100, 2) // 净值总收益率

	if !CheckFundEquityEqValuation(fund) {
		userFund.CostValuationAmount = _utils.CalcFloat64(userFund.Shares*fund.Valuation, 2)                               // 估值持仓金额
		userFund.TotalValuation = _utils.CalcFloat64(valuationDifference*userFund.Shares, 2)                               // 估值总收益
		userFund.TotalValuationYield = _utils.CalcFloat64(userFund.TotalValuation/userFund.CostAmount*100, 2)              // 估值总收益率
		userFund.TodayValuation = _utils.CalcFloat64(_utils.CalcFloat64(fund.Valuation-fund.Equity, 4)*userFund.Shares, 2) // 今日收益估值
	}

	// 昨日或今日收益
	userFund.TodayEquity = _utils.CalcFloat64(fund.EquityIncrease*userFund.Shares, 2) // 今日收益估值
}

// DeleteUserFundByCode 删除用户基金
func (f *fund) DeleteUserFundByCode(ctx context.Context, userID int64, code string) error {
	return f.fundRepo.DeleteUserFundByCode(ctx, userID, code)
}
