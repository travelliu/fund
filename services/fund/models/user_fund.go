// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package models

import (
	"github.com/travelliu/fund/utils/databases"
)

// UserFund 用户基金
type UserFund struct {
	databases.Model
	UserFundPost
	UserID              int64   `json:"userID,string"`
	SellingPrice        float64 `json:"sellingPrice"`
	PurchasePrice       float64 `json:"purchasePrice"`
	CostAmount          float64 `json:"costAmount"`          // 持仓金额
	CostEquityAmount    float64 `json:"costEquityAmount"`    // 净值持仓金额
	CostValuationAmount float64 `json:"costValuationAmount"` // 估值持仓金额
	TodayEquity         float64 `json:"todayEquity"`         // 今日净值收益
	TodayValuation      float64 `json:"todayValuation"`      // 今日估值收益
	TotalEquity         float64 `json:"totalEquity"`         // 净值总收益
	TotalEquityYield    float64 `json:"totalEquityYield"`    // 净值总收益率
	TotalValuation      float64 `json:"totalValuation"`      // 估值总收益
	TotalValuationYield float64 `json:"totalValuationYield"` // 估值总收益率
}

// TableName 表名
func (u *UserFund) TableName() string {
	return "user_funds"
}

// UserFundResponseList User Fund Response List
type UserFundResponseList struct {
	FundBase
	*UserFund
	CostPre float64 `json:"costPre"` // 持仓占比
}

// UserFundResponse 用户基金
type UserFundResponse struct {
	List                []*UserFundResponseList
	CostAmount          float64 `json:"costAmount"`          // 持仓金额
	CostEquityAmount    float64 `json:"costEquityAmount"`    // 持仓净值
	CostValuationAmount float64 `json:"costValuationAmount"` // 持仓估值
	TotalEquity         float64 `json:"totalEquity"`         // 净值总收益
	TotalEquityYield    float64 `json:"totalEquityYield"`    // 净值总收益率
	TodayEquity         float64 `json:"todayEquity"`         // 今日净值收益
	TodayValuation      float64 `json:"todayValuation"`      // 今日估值收益
}

// UserFundPost 用户申请添加基金数据
type UserFundPost struct {
	Code        string  `json:"code" validate:"required"`
	CostPrice   float64 `json:"costPrice" validate:"required"`
	Shares      float64 `json:"shares" validate:"required"`
	SellingPer  int     `json:"sellingPer" validate:"required,min=1,max=100"`
	PurchasePer int     `json:"purchasePer" validate:"required,min=1,max=100"`
}
