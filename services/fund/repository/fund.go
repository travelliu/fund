// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package repository

import (
	"context"
	_fundMod "github.com/travelliu/fund/services/fund/models"
	"github.com/travelliu/fund/utils"
	"strings"
)

// CreateFund 添加基金信息
func (u *user) CreateFund(ctx context.Context, fund *_fundMod.Fund) error {
	if fund.ID == 0 {
		fund.ID = utils.GenerateID()
	}
	db := u.db.Create(fund)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

// QueryFundByCode 通过code查询基金
func (u *user) QueryFundByCode(ctx context.Context, code string) (*_fundMod.Fund, error) {
	var (
		fund = &_fundMod.Fund{}
	)
	err := u.db.Where("code = ?", code).Find(fund).Error
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return fund, nil
		}
		return nil, err
	}
	
	return fund, err
}

//  QueryFundByUserID 查询用户拥有的基金信息
func (u *user) QueryFundByUserID(ctx context.Context, userID int64) ([]*_fundMod.Fund, error) {
	var (
		funds = []*_fundMod.Fund{}
	)
	err := u.db.Find(&funds).Joins("JOIN user_funds ON user_funds.code = funds.code AND user_funds.user_id = ?", userID).Error
	
	return funds, err
}

// QueryAllFund 通过code查询基金
func (u *user) QueryAllFund(ctx context.Context) ([]*_fundMod.Fund, error) {
	var (
		funds = []*_fundMod.Fund{}
	)
	err := u.db.Find(&funds).Error
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return funds, nil
		}
		return nil, err
	}
	
	return funds, err
}

// QueryAllFund 通过code查询基金
func (u *user) UpdateFund(ctx context.Context, fund *_fundMod.Fund, ) error {
	return u.db.Model(&_fundMod.Fund{}).Update(fund).Error
}
