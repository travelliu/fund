// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package http

import (
	"github.com/gin-gonic/gin"
	_utils "github.com/travelliu/fund/utils"
	_err "github.com/travelliu/fund/utils/errors"
)

// SystemErrorCode 查看系统错误码
func (h *httpHandler) SystemErrorCode(c *gin.Context) {
	allErrorCodes := _err.ErrorMsg
	lang := _utils.GetLang(c)
	errorCode, ok := allErrorCodes[lang]
	if !ok {
		errorCode = allErrorCodes[_err.LangZh]
	}

	_utils.HTTPRequestSuccess(c, errorCode)
}
