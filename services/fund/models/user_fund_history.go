// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package models

// UserFundHistory 基金历史信息
type UserFundHistory struct {
	UserFund
}

// TableName 表名
func (u *UserFundHistory) TableName() string {
	return "users_funds_history"
}
