// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package repository

import (
	"context"
	_userMod "github.com/travelliu/fund/services/user/models"
	"github.com/travelliu/fund/utils"
	"strings"
)

// CreateUser 数据库里添加用户
func (u *user) CreateUser(ctx context.Context, user *_userMod.Users) error {
	if user.ID == 0 {
		user.ID = utils.GenerateID()
	}
	db := u.db.Create(user)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (u *user) QueryUserByUserName(ctx context.Context, userName string) (*_userMod.Users, error) {
	var (
		user = &_userMod.Users{}
	)
	err := u.db.Where("username = ?", userName).Find(user).Error
	return user, err
}

// CheckUserNameExist 检查用户名是否已存在. true 已经存在
func (u *user) CheckUserNameExist(ctx context.Context, userName string) (bool, error) {
	var total int = 0
	err := u.db.Where("username = ?", userName).Find(&_userMod.Users{}).Count(&total).Error
	if err != nil && strings.Contains(err.Error(), "error record not found") {
		return false, err
	}
	return total >= 1, nil
}

// QueryUserByID 通过id查询用户信息
func (u *user) QueryUserByID(ctx context.Context, id int64) (*_userMod.Users, error) {
	var (
		user = &_userMod.Users{}
	)
	user.ID = id
	err := u.db.Find(user).Error
	return user, err
}
