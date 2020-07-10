// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package user

import (
	"context"
	_userMod "github.com/travelliu/fund/services/user/models"
)

//go:generate mockery -name=Repository

// Repository 数据库操作接口
type Repository interface {
	usersRepo
}

type usersRepo interface {
	CreateUser(ctx context.Context, user *_userMod.Users) error                        // 数据库里增加用户
	CheckUserNameExist(ctx context.Context, userName string) (bool, error)             // 检查用户名是否存在
	QueryUserByUserName(ctx context.Context, userName string) (*_userMod.Users, error) // 通过username查询用户
	QueryUserByID(ctx context.Context, id int64) (*_userMod.Users, error)
}
