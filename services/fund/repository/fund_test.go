// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package repository

import (
	"context"
	"github.com/jinzhu/gorm"
	_fundMod "github.com/travelliu/fund/services/fund/models"
	"github.com/travelliu/fund/utils/databases"
	"testing"
)

func Test_user_UpdateFund(t *testing.T) {
	type fields struct {
		db *gorm.DB
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
			name: "test",
			args: args{
				ctx: ctx,
				fund: &_fundMod.Fund{
					Model: databases.Model{
						ID: 1281180256048254976,
					},
					FundBase: _fundMod.FundBase{
						Equity: 10.5473,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repo{
				db: db,
			}
			if err := r.UpdateFund(tt.args.ctx, tt.args.fund); (err != nil) != tt.wantErr {
				t.Errorf("UpdateFund() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
