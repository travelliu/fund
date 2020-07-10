// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package repository

import (
	"context"
	_fundMod "github.com/travelliu/fund/services/fund/models"
	"github.com/travelliu/fund/utils"
	"github.com/travelliu/fund/utils/databases"
	_err "github.com/travelliu/fund/utils/errors"
	"github.com/travelliu/fund/utils/trace"
	"strings"
)

// CreateUserFund 添加用户基金
func (u *user) CreateUserFund(ctx context.Context, userFund *_fundMod.UserFund) error {
	if userFund.ID == 0 {
		userFund.ID = utils.GenerateID()
	}
	db := u.db.Create(userFund)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

// EditUserFund 修改用户基金
func (u *user) EditUserFund(ctx context.Context, userFund *_fundMod.UserFund) error {
	if userFund.ID == 0 || userFund.Code == "" || userFund.UserID == 0 {
		return _err.New(_err.ErrEditUserFundData, "")
	}
	return u.db.Model(&_fundMod.UserFund{}).Update(userFund).Error
}

// QueryUserFundByUserID 查询用户全部基金信息
func (u *user) QueryUserFundByUserID(ctx context.Context, userID int64) ([]*_fundMod.UserFund, error) {
	var (
		userFunds = []*_fundMod.UserFund{}
	)
	err := u.db.Where("user_id = ?", userID).Find(&userFunds).Error
	return userFunds, err
}

// QueryUserFundByCode 通过Code查询用户基金
func (u *user) QueryUserFundByCode(ctx context.Context, userID int64, code string) (*_fundMod.UserFund, error) {
	var (
		userFund = &_fundMod.UserFund{}
	)
	err := u.db.Where("code = ? and user_id = ?", code, userID).Find(userFund).Error
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return userFund, nil
		}
		return nil, err
	}
	
	return userFund, err
}

// DeleteUserFundByUserID 删除用户全部基金
func (u *user) DeleteUserFundByUserID(ctx context.Context, userID int64) error {
	logger.WithField(string(trace.ContextKeyReqID), trace.GetReqID(ctx)).Debugf("the DeleteUserFundByUserID userID %v", userID)
	var (
		userFund = &_fundMod.UserFund{}
	)
	err := u.db.Where("user_id = ?", userID).Delete(userFund).Error
	return err
}

// DeleteUserFundByID 根据id删除用户基金
func (u *user) DeleteUserFundByID(ctx context.Context, userID, id int64) error {
	var (
		userFund = &_fundMod.UserFund{
			Model: databases.Model{
				ID: id,
			},
			UserID: userID,
		}
	)
	err := u.db.Delete(userFund).Error
	return err
}

// DeleteUserFundByCode 根据code删除用户基金
func (u *user) DeleteUserFundByCode(ctx context.Context, userID int64, code string) error {
	var (
		userFund = &_fundMod.UserFund{}
	)
	userFund.Code = code
	userFund.UserID = userID
	err := u.db.Where(userFund).Delete(userFund).Error
	return err
}
