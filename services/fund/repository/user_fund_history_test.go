// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package repository

import (
	"context"
	_fundMod "github.com/travelliu/fund/services/fund/models"
	"testing"
)

func Test_repo_ReplaceUserFundHistory(t *testing.T) {
	type args struct {
		ctx               context.Context
		code              string
		userFundHistories []*_fundMod.UserFundHistory
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				ctx:  ctx,
				code: "001618",
				userFundHistories: []*_fundMod.UserFundHistory{
					&_fundMod.UserFundHistory{
						UserFund: _fundMod.UserFund{
							UserID: 111,
							UserFundPost: _fundMod.UserFundPost{
								Code: "001618",
							},
						},
					},
					&_fundMod.UserFundHistory{
						UserFund: _fundMod.UserFund{
							UserID: 111,
							UserFundPost: _fundMod.UserFundPost{
								Code: "001618",
							},
						},
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
			if err := r.ReplaceUserFundHistory(tt.args.ctx, tt.args.code, tt.args.userFundHistories); (err != nil) != tt.wantErr {
				t.Errorf("ReplaceUserFundHistory() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
