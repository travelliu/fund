// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package http

import (
	"github.com/gin-gonic/gin"
	_fundMod "github.com/travelliu/fund/services/fund/models"
	_utils "github.com/travelliu/fund/utils"
	"github.com/travelliu/fund/utils/databases"
	_err "github.com/travelliu/fund/utils/errors"
	"github.com/travelliu/fund/utils/trace"
)

// CreateUserFund 用户添加基金
func (h *httpHandler) CreateUserFund(c *gin.Context) {
	var ctx = c.Request.Context()
	userFund, err := getUserPost(c)
	if err != nil {
		_utils.HTTPRequestFailed(c, err, _err.ERROR)
		return
	}
	newUserFund, err := h.fundUC.CreateUserFund(ctx, userFund)
	if err != nil {
		logger.WithField(string(trace.ContextKeyReqID), trace.GetReqID(ctx)).Errorf("the CreateUserFund error %s", err)
		_utils.HTTPRequestFailed(c, err, _err.ERROR)
		return
	}
	_utils.HTTPRequestSuccess(c, newUserFund)
}

// DeleteUserFund 用户删除基金
func (h *httpHandler) DeleteUserFund(c *gin.Context) {
	var (
		ctx    = c.Request.Context()
		userID = c.GetInt64("userID")
		code   = c.GetString("fundCode")
	)
	if err := h.fundUC.DeleteUserFundByCode(ctx, userID, code); err != nil {
		logger.WithField(string(trace.ContextKeyReqID), trace.GetReqID(ctx)).Errorf("the DeleteUserFundByCode error %s", err)
		_utils.HTTPRequestFailed(c, err, _err.ERROR)
		return
	}
	_utils.HTTPRequestSuccess(c, "")
}

func getUserPost(c *gin.Context) (*_fundMod.UserFund, error) {
	var (
		userFundPost _fundMod.UserFundPost
		ctx          = c.Request.Context()
	)
	err := c.ShouldBind(&userFundPost)
	if err != nil {
		logger.WithField(string(trace.ContextKeyReqID), trace.GetReqID(ctx)).Errorf("the BindJSON error %s", err)
		params := _err.ConvertErrorToString(err)

		return nil, _err.New(_err.ErrInvalidParams, "", params)
	}
	if err := validate.Struct(&userFundPost); err != nil {
		logger.WithField(string(trace.ContextKeyReqID), trace.GetReqID(ctx)).Errorf("the Struct error %s", err)
		params := _err.ConvertErrorToString(err)
		return nil, _err.New(_err.ErrInvalidParams, "", params)
	}
	userFund := &_fundMod.UserFund{
		Model:         databases.Model{},
		UserFundPost:  userFundPost,
		UserID:        c.GetInt64("userID"),
		SellingPrice:  0,
		PurchasePrice: 0,
	}
	return userFund, err
}

// GetUserFunds 获取用户全部基金
func (h *httpHandler) GetUserFunds(c *gin.Context) {
	var (
		userID = c.GetInt64("userID")
		ctx    = c.Request.Context()
	)
	result, err := h.fundUC.QueryUserFundByUserID(ctx, userID)
	if err != nil {
		logger.WithField(string(trace.ContextKeyReqID), trace.GetReqID(ctx)).Errorf("the CreateUserFund error %s", err)
		_utils.HTTPRequestFailed(c, err, _err.ERROR)
		return
	}
	_utils.HTTPRequestSuccess(c, result)
}

// GetUserFund 获取用户单个基金
func (h *httpHandler) GetUserFund(c *gin.Context) {

}
