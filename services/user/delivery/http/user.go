// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package http

import "C"
import (
	"github.com/gin-gonic/gin"
	_userMod "github.com/travelliu/fund/services/user/models"
	_utils "github.com/travelliu/fund/utils"
	_err "github.com/travelliu/fund/utils/errors"
	"github.com/travelliu/fund/utils/trace"
)

// CreateUser 用户注册
func (h *httpHandler) CreateUser(c *gin.Context) {
	var (
		userPost _userMod.UserPost
		ctx      = c.Request.Context()
	)
	err := c.ShouldBind(&userPost)
	if err != nil {
		logger.WithField(string(trace.ContextKeyReqID), trace.GetReqID(ctx)).Errorf("the BindJSON error %s", err)
		params := _err.ConvertErrorToString(err)
		_utils.HTTPRequestFailed(c, nil, _err.ErrInvalidParams, params, "")
		return
	}

	if err := validate.Struct(&userPost); err != nil {
		logger.WithField(string(trace.ContextKeyReqID), trace.GetReqID(ctx)).Errorf("the Struct error %s", err)
		params := _err.ConvertErrorToString(err)
		_utils.HTTPRequestFailed(c, nil, _err.ErrInvalidParams, params, "")
		return
	}

	if !_utils.Password(userPost.Password) {
		_utils.HTTPRequestFailed(c, nil, _err.ErrPasswordRule)
		return
	}
	user := &_userMod.Users{
		Username: userPost.Username,
		Password: userPost.Password,
	}
	if err := h.userUc.CreateUser(ctx, user); err != nil {
		_utils.HTTPRequestFailed(c, err, _err.ERROR)
		return
	}

	logger.WithField(string(trace.ContextKeyReqID), trace.GetReqID(ctx)).Infof("the user input post %+v", userPost)
	_utils.HTTPRequestSuccess(c, user)
}

func (h *httpHandler) GetUserInfo(c *gin.Context) {
	ctx := c.Request.Context()
	userID := c.GetInt64("userID")
	userInfo, err := h.userUc.QueryUserByID(ctx, userID)
	if err != nil {
		_utils.HTTPRequestFailed(c, err, _err.ERROR)
		return
	}
	_utils.HTTPRequestSuccess(c, userInfo)
}
