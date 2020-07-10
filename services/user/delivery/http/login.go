// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package http

import (
	"github.com/gin-gonic/gin"
	_userMod "github.com/travelliu/fund/services/user/models"
	_utils "github.com/travelliu/fund/utils"
	_err "github.com/travelliu/fund/utils/errors"
	_jwt "github.com/travelliu/fund/utils/jwt"
	"github.com/travelliu/fund/utils/trace"
)

// Login 登录
func (h *httpHandler) Login(c *gin.Context) {
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
	if userPost.Username == "" || userPost.Password == "" {
		_utils.HTTPRequestFailed(c, nil, _err.ErrUserNameIsZero)
		return
	}
	user := &_userMod.Users{
		Username: userPost.Username,
		Password: userPost.Password,
	}
	userInfo, err := h.userUc.Login(ctx, user)
	if err != nil {
		_utils.HTTPRequestFailed(c, err, _err.ERROR)
		return
	}
	token, err := _jwt.GenerateToken(userInfo.ID, userInfo.Username, 24*7)
	if err != nil {
		logger.WithField(string(trace.ContextKeyReqID), trace.GetReqID(ctx)).Errorf("the GenerateToken %s error %v",
			user.Username, err)
		_utils.HTTPRequestFailed(c, err, _err.ERROR)
		return
	}

	c.SetCookie(_jwt.CookieKey, token, 3600*24*7, "/", "", false, true)
	logger.Debugf("the token %s", token)
	_utils.HTTPRequestSuccess(c, userInfo)

}

// Login 退出
func (h *httpHandler) Logout(c *gin.Context) {
	c.SetCookie(_jwt.CookieKey, "", -1, "/", "", false, true)
	_utils.HTTPRequestSuccess(c, "ok")
}
