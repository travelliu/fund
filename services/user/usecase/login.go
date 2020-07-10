// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package usecase

import (
	"context"
	_userMod "github.com/travelliu/fund/services/user/models"
	_utils "github.com/travelliu/fund/utils"
	_err "github.com/travelliu/fund/utils/errors"
	"github.com/travelliu/fund/utils/trace"
)

func (u *user) Login(ctx context.Context, user *_userMod.Users) (*_userMod.Users, error) {
	userInfo, err := u.userRepo.QueryUserByUserName(ctx, user.Username)
	if err != nil || userInfo.Username == "" || userInfo.ID == 0 {
		logger.WithField(string(trace.ContextKeyReqID), trace.GetReqID(ctx)).Errorf("the QueryUserByUserName %s error %v",
			user.Username, err)
		return nil, _err.New(_err.ErrUserNotExisted, "")
	}
	if !_utils.ComparePassword(userInfo.Password, user.Password) {
		return nil, _err.New(_err.ErrUserPwd, "")
	}
	return userInfo, nil
}
