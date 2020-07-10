// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	_fund "github.com/travelliu/fund/services/fund"
	_utils "github.com/travelliu/fund/utils"
	_err "github.com/travelliu/fund/utils/errors"
	"github.com/travelliu/fund/utils/trace"
)

// UserFundMid 用户基金中间件
func UserFundMid(logger *logrus.Logger, fundUC _fund.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		userID := c.GetInt64("userID")
		fundCode := c.Param("fundCode")
		if fundCode == "" {
			_utils.HTTPRequestFailed(c, nil, _err.ErrFundCode)
			c.Abort()
			return
		}
		if uFund, err := fundUC.QueryUserFundByUserIDAndCode(ctx, userID, fundCode); err != nil {
			logger.WithField(string(trace.ContextKeyReqID), trace.GetReqID(ctx)).Errorf("the QueryFundByCode error %s", err)
			_utils.HTTPRequestFailed(c, nil, _err.ERROR)
			c.Abort()
			return
		} else if uFund.ID == 0 {
			_utils.HTTPRequestFailed(c, nil, _err.ErrUserFundNotExisted)
			c.Abort()
			return
		}
		c.Set("fundCode", fundCode)
		c.Next()
	}
}
