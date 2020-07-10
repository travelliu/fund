// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package utils

import (
	"github.com/stretchr/testify/assert"
	_fundMod "github.com/travelliu/fund/services/fund/models"
	"github.com/travelliu/fund/utils/databases"
	"testing"
	"time"
)

func TestParseFund(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    *_fundMod.Fund
		wantErr bool
	}{
		{
			name:    "501016",
			args:    args{s: "jsonpgz({\n\t\"fundcode\":\"501016\",\n\t\"name\":\"国泰中证申万证券行业指数\",\n\t\"jzrq\":\"2020-07-08\",\n\t\"dwjz\":\"1.4073\",\n\t\"gsz\":\"1.4086\",\n\t\"gszzl\":\"0.10\",\n\t\"gztime\":\"2020-07-09 13:06\"\n});"},
			wantErr: false,
			want: &_fundMod.Fund{
				Model: databases.Model{},
				FundBase: _fundMod.FundBase{Code: "501016",
					Name:          "国泰中证申万证券行业指数",
					Equity:        1.4073,
					EquityDate:    "2020-07-08",
					Valuation:     1.4086,
					ValuationPre:  0.10,
					ValuationTime: databases.TimeInt64(time.Unix(1594271160, 0)),
				}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseFundString(tt.args.s)
			assert.NoError(t, err)
			assert.Equal(t, got, tt.want)
		})
	}
}

func Test_getFundData(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name    string
		args    args
		want    *_fundMod.Fund
		wantErr bool
	}{
		{
			name:    "001618",
			args:    args{code: "001618"},
			wantErr: false,
			want: &_fundMod.Fund{
				Model: databases.Model{},
				FundBase: _fundMod.FundBase{
					EquityPre: 2.79,
				}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getFundData(tt.args.code)
			assert.NoError(t, err)
			assert.Equal(t, got, tt.want)
		})
	}
}
