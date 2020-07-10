// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package repository

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/travelliu/fund/config"
	_fundMod "github.com/travelliu/fund/services/fund/models"
	"github.com/travelliu/fund/utils/databases"
	"github.com/travelliu/fund/utils/trace"
	"os"
	"testing"
)

var (
	db  *gorm.DB
	ctx context.Context
)

func TestMain(m *testing.M) {
	var (
		dbConf = &config.DB{
			Type:   "sqlite",
			Dir:    fmt.Sprintf("%s/%s", os.Getenv("GOPATH"), "src/github.com/travelliu/fund/db"),
			User:   "",
			Pwd:    "",
			Host:   "",
			Port:   0,
			Dbname: "fund",
		}
		err error
	)
	ctx = trace.AttachReqID(context.Background())
	db, err = databases.InitDatabase(dbConf)
	if err != nil {
		return
	}
	m.Run()
}

func Test_user_CreateUserFund(t *testing.T) {
	type args struct {
		ctx      context.Context
		userFund *_fundMod.UserFund
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "CreateUserFund",
			args: args{
				ctx: ctx,
				userFund: &_fundMod.UserFund{
					UserID: 1281036673953566720,
					UserFundPost: _fundMod.UserFundPost{
						Code: "001618",
					},
				},
			},
		},
		{
			name: "CreateUserFund",
			args: args{
				ctx: ctx,
				userFund: &_fundMod.UserFund{
					UserID: 1281036673953566720,
					UserFundPost: _fundMod.UserFundPost{
						Code: "001617",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &user{
				db: db,
			}
			if err := u.CreateUserFund(tt.args.ctx, tt.args.userFund); (err != nil) != tt.wantErr {
				t.Errorf("CreateUserFund() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_user_DeleteUserFundByCode(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID int64
		code   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "DeleteUserFundByCode",
			args: args{
				ctx:    ctx,
				userID: 1281036673953566720,
				code:   "001617",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &user{
				db: db,
			}
			if err := u.DeleteUserFundByCode(tt.args.ctx, tt.args.userID, tt.args.code); (err != nil) != tt.wantErr {
				t.Errorf("DeleteUserFundByCode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_user_DeleteUserFundByUserID(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "DeleteUserFundByUserID",
			args: args{
				ctx:    ctx,
				userID: 1281036673953566720,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &user{
				db: db,
			}
			if err := u.DeleteUserFundByUserID(tt.args.ctx, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("DeleteUserFundByUserID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}


func Test_user_QueryAllFund(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    []*_fundMod.Fund
		wantErr bool
	}{
		{
			name:"test",
			args: args{ctx: ctx},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &user{
				db: db,
			}
			got, err := u.QueryAllFund(tt.args.ctx)
			assert.NoError(t, err)
			for _, g := range got {
				fmt.Printf("%+v\n", g)
			}
		})
	}
}
