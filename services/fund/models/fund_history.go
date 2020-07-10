// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package models

// FundHistory 基金历史信息
type FundHistory struct {
	Fund
}

// TableName 表名
func (u *FundHistory) TableName() string {
	return "funds_history"
}
