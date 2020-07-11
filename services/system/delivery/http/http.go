// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

var (
	// logger   *logrus.Logger
	validate *validator.Validate
)

func init() {
	// logger = logs.NewLogger()
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
	// fundUC _fund.UseCase
}

// NewFundHTTP New Fund Http
func NewSystemHTTP(r *gin.RouterGroup) {
	handler := &httpHandler{
		// fundUC: fundUC,
	}
	system := r.Group("/system")
	system.GET("errorCode", handler.SystemErrorCode)
}
