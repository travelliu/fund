// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package usecase

import (
	"github.com/stretchr/testify/assert"
	_fundMod "github.com/travelliu/fund/services/fund/models"
	"github.com/travelliu/fund/utils/databases"
	"testing"
	"time"
)

func Test_calcUserFund(t *testing.T) {
	type args struct {
		userFund *_fundMod.UserFund
		fund     *_fundMod.Fund
	}
	tests := []struct {
		name string
		args args
		want *_fundMod.UserFund
	}{
		{
			name: "001618",
			args: args{
				userFund: &_fundMod.UserFund{
					Model: databases.Model{},
					UserFundPost: _fundMod.UserFundPost{
						Code:        "001618",
						CostPrice:   1.1922,
						Shares:      89900.2,
						SellingPer:  30,
						PurchasePer: 5,
					},
					UserID: 1281036673953566720,
				},
				fund: &_fundMod.Fund{
					Model: databases.Model{},
					FundBase: _fundMod.FundBase{
						Code:          "001618",
						Name:          "天弘中证电子ETF联接C",
						Equity:        1.5473,
						EquityPre:     2.79,
						EquityDate:    "2020-07-09",
						Valuation:     1.5475,
						ValuationPre:  2.80,
						ValuationTime: databases.TimeInt64(time.Unix(1594278000, 0)),
					},
				},
			},
			want: &_fundMod.UserFund{
				Model: databases.Model{},
				UserFundPost: _fundMod.UserFundPost{
					Code:        "001618",
					CostPrice:   1.1922,
					Shares:      89900.2,
					SellingPer:  30,
					PurchasePer: 5,
				},
				UserID:              1281036673953566720,
				SellingPrice:        1.5499,
				PurchasePrice:       1.4699,
				CostAmount:          107179.02,
				CostEquityAmount:    139102.58,
				CostValuationAmount: 139120.56,
				TotalEquity:         31923.56,
				TotalEquityYield:    29.79,
				TotalValuation:      31941.54,
				TodayEquity:         3775.81,
				TodayValuation:      17.98,
			},
		},
		{
			name: "001618",
			args: args{
				userFund: &_fundMod.UserFund{
					Model: databases.Model{},
					UserFundPost: _fundMod.UserFundPost{
						Code:        "000746",
						CostPrice:   2.3815,
						Shares:      6193.58,
						SellingPer:  30,
						PurchasePer: 5,
					},
					UserID: 1281036673953566720,
				},
				fund: &_fundMod.Fund{
					Model: databases.Model{},
					FundBase: _fundMod.FundBase{
						Code:          "000746",
						Equity:        3.215,
						EquityPre:     2.95,
						EquityDate:    "2020-07-09",
						Valuation:     3.2247,
						ValuationPre:  0.30,
						ValuationTime: databases.TimeInt64(time.Unix(1594278000, 0)),
					},
				},
			},
			want: &_fundMod.UserFund{
				Model: databases.Model{},
				UserFundPost: _fundMod.UserFundPost{
					Code:        "000746",
					CostPrice:   2.3815,
					Shares:      6193.58,
					SellingPer:  30,
					PurchasePer: 5,
				},
				UserID:              1281036673953566720,
				SellingPrice:        3.096,
				PurchasePrice:       3.0542,
				CostAmount:          14750.01,
				CostEquityAmount:    19912.36,
				CostValuationAmount: 19972.44,
				TodayEquity:         570.43,  // 昨日收益
				TodayValuation:      60.08,   // 今日估值收益
				TotalEquity:         5162.35, // 持有收益
				TotalEquityYield:    35,      // 持有收益率
				TotalValuation:      5222.43, // 估值收益
				TotalValuationYield: 35.41,       // 估值收益
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calcUserFund(tt.args.userFund, tt.args.fund)
			assert.Equal(t, tt.args.userFund, tt.want)
		})
	}
}
