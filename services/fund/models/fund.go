// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package models

import (
	"github.com/travelliu/fund/utils/databases"
)

// FundBase 基金信息
type FundBase struct {
	Code          string              `json:"code",gorm:"type:varchar(30);unique_index"`
	Name          string              `json:"name"`
	Equity        float64             `json:"equity"`
	EquityPre     float64             `json:"equityPre"`
	EquityDate    string              `json:"equityDate"`
	Valuation     float64             `json:"valuation"`
	ValuationPre  float64             `json:"valuationPre"`
	ValuationTime databases.TimeInt64 `json:"valuationTime" gorm:"type:datetime"`
}

// Fund 基金信息
type Fund struct {
	databases.Model
	FundBase
}

// TableName 表名
func (u *Fund) TableName() string {
	return "funds"
}

// UserFund 用户基金
type UserFund struct {
	databases.Model
	UserFundPost
	UserID              int64   `json:"userID,string"`
	SellingPrice        float64 `json:"sellingPrice"`
	PurchasePrice       float64 `json:"purchasePrice"`
	CostAmount          float64 `json:"costAmount"`          // 持仓金额
	CostEquityAmount    float64 `json:"costEquityAmount"`      // 净值持仓金额
	CostValuationAmount float64 `json:"costValuationAmount"`       // 估值持仓金额
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

// DayFund 天天基金数据
/*
{
	"fundcode":"501016",
	"name":"国泰中证申万证券行业指数",
	"jzrq":"2020-07-08",
	"dwjz":"1.4073",
	"gsz":"1.4086",
	"gszzl":"0.10",
	"gztime":"2020-07-09 13:06"
}
*/
type DayFund struct {
	Code          string `json:"fundcode"` // 基金编码
	Name          string `json:"name"`     // 基金名称
	Equity        string `json:"dwjz"`     // 单位净值
	EquityData    string `json:"jzrq"`     // 净值日期
	Valuation     string `json:"gsz"`      // 估值
	ValuationPre  string `json:"gszzl"`    // 估值百分比
	ValuationTime string `json:"gztime"`   // 估值时间

}
