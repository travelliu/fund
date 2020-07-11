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
func (r *repo) CreateUserFund(ctx context.Context, userFund *_fundMod.UserFund) error {
	if userFund.ID == 0 {
		userFund.ID = utils.GenerateID()
	}
	db := r.db.Create(userFund)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

// EditUserFund 修改用户基金
func (r *repo) EditUserFund(ctx context.Context, userFund *_fundMod.UserFund) error {
	if userFund.ID == 0 || userFund.Code == "" || userFund.UserID == 0 {
		return _err.New(_err.ErrEditUserFundData, "")
	}
	return r.db.Model(&_fundMod.UserFund{}).Update(userFund).Error
}

// QueryUserFundByUserID 查询用户全部基金信息
func (r *repo) QueryUserFundByUserID(ctx context.Context, userID int64) ([]*_fundMod.UserFund, error) {
	var (
		userFunds = []*_fundMod.UserFund{}
	)
	err := r.db.Where("user_id = ?", userID).Find(&userFunds).Error
	return userFunds, err
}

// QueryUserFundByUserIDAndCode 通过Code查询用户基金
func (r *repo) QueryUserFundByUserIDAndCode(ctx context.Context, userID int64, code string) (*_fundMod.UserFund, error) {
	var (
		userFund = &_fundMod.UserFund{}
	)
	err := r.db.Where("code = ? and user_id = ?", code, userID).Find(userFund).Error
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return userFund, nil
		}
		return nil, err
	}

	return userFund, err
}

// QueryUserFundByCode 通过Code查询用户基金
func (r *repo) QueryUserFundByCode(ctx context.Context, code string) ([]*_fundMod.UserFund, error) {
	var (
		userFund = []*_fundMod.UserFund{}
	)
	err := r.db.Where("code = ?", code).Find(&userFund).Error
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return userFund, nil
		}
		return nil, err
	}

	return userFund, err
}

// DeleteUserFundByUserID 删除用户全部基金
func (r *repo) DeleteUserFundByUserID(ctx context.Context, userID int64) error {
	logger.WithField(string(trace.ContextKeyReqID), trace.GetReqID(ctx)).Debugf("the DeleteUserFundByUserID userID %v", userID)
	var (
		userFund = &_fundMod.UserFund{}
	)
	err := r.db.Where("user_id = ?", userID).Delete(userFund).Error
	return err
}

// DeleteUserFundByID 根据id删除用户基金
func (r *repo) DeleteUserFundByID(ctx context.Context, userID, id int64) error {
	var (
		userFund = &_fundMod.UserFund{
			Model: databases.Model{
				ID: id,
			},
			UserID: userID,
		}
	)
	err := r.db.Delete(userFund).Error
	return err
}

// DeleteUserFundByCode 根据code删除用户基金
func (r *repo) DeleteUserFundByCode(ctx context.Context, userID int64, code string) error {
	var (
		userFund = &_fundMod.UserFund{}
	)
	userFund.Code = code
	userFund.UserID = userID
	err := r.db.Where(userFund).Delete(userFund).Error
	return err
}

func (r *repo) QueryUserFundAllByUserID(ctx context.Context, userID int64) error {
	var (
		userFund              = &_fundMod.UserFund{}
		userFundResponseLists = []*_fundMod.UserFundResponseList{}
	)
	err := r.db.Table(
		userFund.TableName(),
	).Select(
		"user_funds.*,funds.equity,funds.equity_pre,funds.equity_date,funds.equity_increase,funds.valuation"+
			",funds.valuation_pre,funds.valuation_time").Joins("join funds on user_funds.code=funds.code").Where(
		"user_funds.user_id = ?", userID,
	).Order(
		"equity desc", true,
	).Scan(
		&userFundResponseLists,
	).Error
	if err != nil {
		logger.WithField(string(trace.ContextKeyReqID), trace.GetReqID(ctx)).Errorf("the QueryUserFundAllByUserID error %s", err)
		return err
	}

	return nil
}
