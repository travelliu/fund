// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package services

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	_fund "github.com/travelliu/fund/services/fund"
	_fundHttp "github.com/travelliu/fund/services/fund/delivery/http"
	_fundRepo "github.com/travelliu/fund/services/fund/repository"
	_fundUc "github.com/travelliu/fund/services/fund/usecase"
	"github.com/travelliu/fund/services/middleware"
	_user "github.com/travelliu/fund/services/user"
	_userHttp "github.com/travelliu/fund/services/user/delivery/http"
	_userRepo "github.com/travelliu/fund/services/user/repository"
	_userUC "github.com/travelliu/fund/services/user/usecase"
	"github.com/travelliu/fund/utils/logs"
	"github.com/travelliu/fund/utils/trace"
)

type repo struct {
	userRepo _user.Repository
	fundRepo _fund.Repository
}

type uc struct {
	userUC _user.UseCase
	fundUC _fund.UseCase
}

var (
	logger *logrus.Logger
)

func init() {
	logger = logs.NewLogger()
}

// NewService New Service
func NewService(db *gorm.DB) *gin.Engine {
	repo := initRepo(db)
	uc := initUc(repo)
	route := initHTTP(uc)
	_fundUc.NewCron(context.Background(), uc.fundUC, "", 4)
	return route
}

func initRepo(db *gorm.DB) repo {
	return repo{
		userRepo: _userRepo.NewUserRepository(db),
		fundRepo: _fundRepo.NewFundRepository(db),
	}
}

func initUc(repo repo) uc {
	userUC := _userUC.NewUsersUc(repo.userRepo)
	fundUC := _fundUc.NewFundUc(repo.fundRepo)
	return uc{
		userUC: userUC,
		fundUC: fundUC,
	}
}

func initHTTP(uc uc) *gin.Engine {
	router := gin.New()
	router.Use(middleware.GinLogger(logger), gin.Recovery(), trace.RequestID())
	apiV1 := router.Group("/api/v1")
	// apiV1.Use(trace.RequestID()
	_userHttp.NewUserHTTP(apiV1, uc.userUC)
	_fundHttp.NewFundHTTP(apiV1, uc.fundUC, uc.userUC)
	return router
	
}
