// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package models

import (
	"github.com/travelliu/fund/utils/databases"
)

// FundBase 基金信息
type FundBase struct {
	Code           string              `json:"code",gorm:"type:varchar(30);unique_index"`
	Name           string              `json:"name"`
	Equity         float64             `json:"equity"`
	EquityPre      float64             `json:"equityPre"`
	EquityIncrease float64             `json:"equityIncrease"`
	EquityDate     string              `json:"equityDate"`
	Valuation      float64             `json:"valuation"`
	ValuationPre   float64             `json:"valuationPre"`
	ValuationTime  databases.TimeInt64 `json:"valuationTime" gorm:"type:datetime"`
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
