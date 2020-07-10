// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package usecase

import (
	"context"
	_fund "github.com/travelliu/fund/services/fund"
	_fundMod "github.com/travelliu/fund/services/fund/models"
	_fundRepo "github.com/travelliu/fund/services/fund/repository"
	"github.com/travelliu/fund/utils/databases"
	"testing"
	"time"
)

func TestCheckFundEquityEqValuation(t *testing.T) {
	type args struct {
		fund *_fundMod.Fund
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Same",
			args: args{fund: &_fundMod.Fund{
				FundBase: _fundMod.FundBase{
					Code:          "",
					Name:          "",
					Equity:        0,
					EquityPre:     0,
					EquityDate:    "2020-07-09",
					Valuation:     0,
					ValuationPre:  0,
					ValuationTime: databases.TimeInt64(time.Unix(1594271160, 0)),
				},
			}},
			want: true,
		},
		{
			name: "000834",
			args: args{fund: &_fundMod.Fund{
				FundBase: _fundMod.FundBase{
					Code:          "000834",
					Name:          "",
					Equity:        0,
					EquityPre:     0,
					EquityDate:    "2020-07-09",
					Valuation:     0,
					ValuationPre:  0,
					ValuationTime: databases.TimeInt64(time.Unix(1594324800, 0)),
				},
			}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckFundEquityEqValuation(tt.args.fund); got != tt.want {
				t.Errorf("CheckFundEquityEqValuation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fund_doFundSync(t *testing.T) {
	type fields struct {
		fundRepo _fund.Repository
	}
	type args struct {
		ctx  context.Context
		fund *_fundMod.Fund
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "001618",
			args: args{
				ctx: ctx,
				fund: &_fundMod.Fund{
					Model: databases.Model{
						ID: 1281180256048254976,
						// CreatedAt: databases.TimeInt64(time.Unix()),
						// UpdatedAt: databases.TimeInt64(time.Unix()),
					},
					FundBase: _fundMod.FundBase{
						Code:           "001618",
						Name:           "天弘中证电子ETF联接C",
						Equity:         1.5367,
						EquityPre:      -0.69,
						EquityIncrease: -0.0106,
						EquityDate:     "2020-07-10",
						Valuation:      1.5357,
						ValuationPre:   -0.75,
						ValuationTime:  databases.TimeInt64(time.Unix(1594364400, 0)),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &fund{
				fundRepo: _fundRepo.NewFundRepository(db),
			}
			if err := f.doFundSync(tt.args.ctx, tt.args.fund); (err != nil) != tt.wantErr {
				t.Errorf("doFundSync() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
