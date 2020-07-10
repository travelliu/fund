// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package usecase

import (
	"context"
	_userMod "github.com/travelliu/fund/services/user/models"
	_utils "github.com/travelliu/fund/utils"
	_err "github.com/travelliu/fund/utils/errors"
	"github.com/travelliu/fund/utils/trace"
)

func (u *user) CreateUser(ctx context.Context, user *_userMod.Users) error {
	var (
		err error
	)
	user.Password, err = _utils.HashPassword(user.Password)
	if err != nil {
		logger.WithField(string(trace.ContextKeyReqID), trace.GetReqID(ctx)).Errorf("the HashPassword error %s", err)
		return _err.New(_err.ErrHashPassword, "")
	}
	if user.Nickname == "" {
		user.Nickname = user.Username
	}
	// checkUserName
	if exited, err := u.userRepo.CheckUserNameExist(ctx, user.Username); err != nil {
		logger.WithField(string(trace.ContextKeyReqID), trace.GetReqID(ctx)).Errorf("the CheckUserNameExist %s error %s",
			user.Username, err)
		return err
	} else if exited {
		return _err.New(_err.ErrUserNameExisted, "")
	}

	return u.userRepo.CreateUser(ctx, user)
}

func (u *user) QueryUserByID(ctx context.Context, id int64) (*_userMod.Users, error) {
	return u.userRepo.QueryUserByID(ctx, id)
}
