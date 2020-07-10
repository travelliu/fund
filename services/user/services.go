// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package user

import (
	"context"
	_userMod "github.com/travelliu/fund/services/user/models"
)

//go:generate mockery -name=UseCase

// UseCase 用户场景
type UseCase interface {
	Login(ctx context.Context, user *_userMod.Users) (*_userMod.Users, error) // 用户登录

	CreateUser(ctx context.Context, user *_userMod.Users) error // 创建用户
	QueryUserByID(ctx context.Context, id int64) (*_userMod.Users, error)
}
