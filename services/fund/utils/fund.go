// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package utils

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	_fundMod "github.com/travelliu/fund/services/fund/models"
	_utils "github.com/travelliu/fund/utils"
	"github.com/travelliu/fund/utils/databases"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

/*
jsonpgz({
	"fundcode":"501016",
	"name":"国泰中证申万证券行业指数",
	"jzrq":"2020-07-08",
	"dwjz":"1.4073",
	"gsz":"1.4086",
	"gszzl":"0.10",
	"gztime":"2020-07-09 13:06"
});

var apidata={ content:"<table class='w782 comm lsjz'><thead><tr><th class='first'>净值日期</th><th>单位净值</th><th>累计净值</th><th>日增长率</th><th>申购状态</th><th>赎回状态</th><th class='tor last'>分红送配</th></tr></thead><tbody><tr><td>2020-07-08</td><td class='tor bold'>2.5482</td><td class='tor bold'>2.6258</td><td class='tor bold red'>0.89%</td><td>开放申购</td><td>开放赎回</td><td class='red unbold'></td></tr></tbody></table>",records:1207,pages:1207,curpage:1};

*/

var (
	fundBaseURL     = "http://fundgz.1234567.com.cn/js/%s.js"
	fundDataBaseURL = "http://fund.eastmoney.com/f10/F10DataApi.aspx?type=lsjz&code=%s&page=1&per=2&sdate=&edate="
)

// ParseFundString Parse Fund
func ParseFundString(s string) (*_fundMod.Fund, error) {
	return parseFundString(s)
}

// parseFundString Parse Fund
func parseFundString(s string) (*_fundMod.Fund, error) {
	s = strings.TrimLeft(s, "jsonpgz(")
	s = strings.TrimRight(s, ");")
	old := &_fundMod.DayFund{}
	if err := json.Unmarshal([]byte(s), old); err != nil {
		return nil, err
	}
	return convertFund(old)
}

func convertFund(old *_fundMod.DayFund) (*_fundMod.Fund, error) {
	equity, err := strconv.ParseFloat(old.Equity, 64)
	if err != nil {
		return nil, err
	}
	valuation, err := strconv.ParseFloat(old.Valuation, 64)
	if err != nil {
		return nil, err
	}
	valuationPre, err := strconv.ParseFloat(old.ValuationPre, 64)
	if err != nil {
		return nil, err

	}
	valuationTime, err := time.ParseInLocation("2006-01-02 15:04", old.ValuationTime, time.Local)
	if err != nil {
		return nil, err

	}
	f := _fundMod.FundBase{
		Code:          strings.TrimSpace(old.Code),
		Name:          strings.TrimSpace(old.Name),
		EquityDate:    strings.TrimSpace(old.EquityData),
		Equity:        equity,
		Valuation:     valuation,
		ValuationPre:  valuationPre,
		ValuationTime: databases.TimeInt64(valuationTime),
	}
	return &_fundMod.Fund{
		FundBase: f,
	}, nil
}

// GetFundInfo Get Fund Info
func GetFundInfo(code string) (*_fundMod.Fund, error) {
	f, err := getFundInfo(code)
	if err != nil {
		return nil, err
	}
	fData, err := getFundData(code)
	if err != nil {
		return nil, err
	}
	f.EquityPre = fData.EquityPre
	f.Equity = fData.Equity
	f.EquityDate = fData.EquityDate
	f.EquityIncrease = fData.EquityIncrease
	return f, nil
}

func getFundInfo(code string) (*_fundMod.Fund, error) {
	url := getFundURL(code)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		// fmt.Println(err)
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		// fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		// fmt.Println(err)
		return nil, err
	}
	return parseFundString(string(body))
}
func getFundURL(code string) string {
	return fmt.Sprintf(fundBaseURL, code)
}
func getFundData(code string) (*_fundMod.Fund, error) {
	url := fmt.Sprintf(fundDataBaseURL, code)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		// fmt.Println(err)
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		// fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		// fmt.Println(err)
		return nil, err
	}
	return convertFundData(string(body))
}

func convertFundData(data string) (*_fundMod.Fund, error) {
	var headings, row []string
	var rows [][]string
	data = strings.TrimLeft(data, "var apidata={ content:\"")
	data = strings.TrimRight(data, "\",records:1208,pages:604,curpage:1};")
	data = fmt.Sprintf("<html><body>%s</body></html>", data)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(data))
	if err != nil {
		return nil, err
	}
	doc.Find("table").Each(func(index int, tableHtml *goquery.Selection) {
		tableHtml.Find("tr").Each(func(indexTr int, rowHtml *goquery.Selection) {
			rowHtml.Find("th").Each(func(indexTh int, tableHeading *goquery.Selection) {
				headings = append(headings, tableHeading.Text())
			})
			rowHtml.Find("td").Each(func(indexTh int, tableCell *goquery.Selection) {
				row = append(row, tableCell.Text())
			})
			if len(row) > 0 {
				rows = append(rows, row)
			}

			row = nil
		})
	})
	if len(rows) < 1 {
		return nil, fmt.Errorf("the parse failed")
	}
	equityPreStr := rows[0][3]
	equityPreStr = strings.TrimRight(equityPreStr, "%")
	equity, err := strconv.ParseFloat(rows[0][1], 64)
	if err != nil {
		return nil, err
	}
	equityPre, err := strconv.ParseFloat(equityPreStr, 64)
	if err != nil {
		return nil, err
	}
	var equityIncrease float64
	if len(rows) == 2 {
		lastEquity, err := strconv.ParseFloat(rows[1][1], 64)
		if err != nil {
			return nil, err
		}
		equityIncrease = _utils.CalcFloat64(equity-lastEquity, 4)
	}
	f := _fundMod.FundBase{
		EquityPre:      equityPre,
		Equity:         equity,
		EquityDate:     rows[0][0],
		EquityIncrease: equityIncrease,
	}
	return &_fundMod.Fund{
		FundBase: f,
	}, nil
}
