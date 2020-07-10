// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	_user "github.com/travelliu/fund/services/user"
	_utils "github.com/travelliu/fund/utils"
	_err "github.com/travelliu/fund/utils/errors"
	_jwt "github.com/travelliu/fund/utils/jwt"
	"github.com/travelliu/fund/utils/trace"
	"net/http"
	"time"
)

// GinLogger 日志记录到文件
func GinLogger(logger *logrus.Logger) gin.HandlerFunc {

	return func(c *gin.Context) {

		// 开始时间
		startTime := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqURI := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()
		clientUserAgent := c.Request.UserAgent()
		// 请求IP
		clientIP := c.ClientIP()
		userID := c.GetInt64("userID")
		// 日志格式
		logger.WithFields(logrus.Fields{
			"statusCode":                  statusCode,
			"latencyTime":                 latencyTime,
			"clientIp":                    clientIP,
			"reqMethod":                   reqMethod,
			"reqUri":                      reqURI,
			string(trace.ContextKeyReqID): trace.GetReqID(c.Request.Context()),
			"userID":                      userID,
			"clientUserAgent":             clientUserAgent,
		}).Info()
	}
}

// GinAuth 检查用户是否已登录
func GinAuth(logger *logrus.Logger, userUc _user.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		token, err := c.Cookie(_jwt.CookieKey)
		if err != nil {
			logger.WithField(string(trace.ContextKeyReqID), trace.GetReqID(ctx)).Errorf("the get Cookie %s error %s",
				_jwt.CookieKey, err)
			_utils.HTTPRequestFailed(c, err, http.StatusUnauthorized)
			c.Abort()
			return

		}
		tokenInfo, err := _jwt.ParseToken(token)
		if err != nil {
			logger.WithField(string(trace.ContextKeyReqID), trace.GetReqID(ctx)).Errorf("the ParseToken error %s",
				_jwt.CookieKey, err)
			_utils.HTTPRequestFailed(c, err, http.StatusUnauthorized)
			c.Abort()
			return
		}

		_, err = userUc.QueryUserByID(ctx, tokenInfo.UserID)
		if err != nil {
			_utils.HTTPRequestFailed(c, err, _err.ErrUserNotExisted)
			c.Abort()
		}
		c.Set("userID", tokenInfo.UserID)
		c.Next()
	}
}
