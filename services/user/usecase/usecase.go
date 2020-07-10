// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package usecase

import (
	"github.com/sirupsen/logrus"
	_user "github.com/travelliu/fund/services/user"
	"github.com/travelliu/fund/utils/logs"
)

var (
	logger *logrus.Logger
)

func init() {
	logger = logs.NewLogger()
}

type user struct {
	userRepo _user.Repository
}

// NewUsersUc New Users Uc
func NewUsersUc(userRepo _user.Repository) _user.UseCase {
	return &user{
		userRepo: userRepo,
	}
}
