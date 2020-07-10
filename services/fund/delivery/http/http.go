// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	_fund "github.com/travelliu/fund/services/fund"
	"github.com/travelliu/fund/services/middleware"
	_user "github.com/travelliu/fund/services/user"
	"github.com/travelliu/fund/utils/logs"
	"reflect"
	"strings"
)

var (
	logger   *logrus.Logger
	validate *validator.Validate
)

func init() {
	logger = logs.NewLogger()
	validate = validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		
		if name == "-" {
			return ""
		}
		
		return name
	})
}

type httpHandler struct {
	fundUC _fund.UseCase
}

// NewFundHTTP New Fund Http
func NewFundHTTP(r *gin.RouterGroup, fundUC _fund.UseCase, userUc _user.UseCase) {
	handler := &httpHandler{
		fundUC: fundUC,
	}
	userFund := r.Group("/userFund")
	userFund.Use(middleware.GinAuth(logger, userUc))
	userFund.POST("", handler.CreateUserFund)
	userFund.GET("", handler.GetUserFunds)
	
	userFundDetail := userFund.Group("/:fundCode")
	userFundDetail.Use(UserFundMid(logger, fundUC))
	userFundDetail.DELETE("", handler.DeleteUserFund)
	// userFundDetail.PATCH("", handler.CreateUserFund)
	userFundDetail.GET("", handler.GetUserFund)
	// user.GET("/userInfo", handler.GetUserInfo)
	// user.GET("/logout", handler.Logout)
}
