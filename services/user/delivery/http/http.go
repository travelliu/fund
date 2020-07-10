// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/travelliu/fund/services/middleware"
	_user "github.com/travelliu/fund/services/user"
	"github.com/travelliu/fund/utils/logs"
	"reflect"
	"regexp"
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
	validate.RegisterValidation("username", func(fl validator.FieldLevel) bool {
		f, err := regexp.Match("^[a-zA-Z]+([_-]?[a-zA-Z0-9])*$",
			[]byte(fl.Field().String()))

		fmt.Println(fl.Field().String(), f, err)
		return f
	})

}

type httpHandler struct {
	userUc _user.UseCase
}

// NewUserHTTP New User HTTP
func NewUserHTTP(r *gin.RouterGroup, userUc _user.UseCase) {
	handler := &httpHandler{
		userUc: userUc,
	}
	r.POST("/login", handler.Login)
	r.POST("/register", handler.CreateUser)
	user := r.Group("/user")
	user.Use(middleware.GinAuth(logger, userUc))
	user.GET("/userInfo", handler.GetUserInfo)
	user.GET("/logout", handler.Logout)
}
