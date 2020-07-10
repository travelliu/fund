package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_err "github.com/travelliu/fund/utils/errors"
	"github.com/travelliu/fund/utils/trace"
	"net/http"
	"strings"
)

// HTTPResponse is a http response
type HTTPResponse struct {
	RequestID string      `json:"requestID"`
	Code      int         `json:"code" example:"400"`
	Message   string      `json:"message" example:"status bad request"`
	Data      interface{} `json:"data" example:"{}"`
}

// NewHTTPResponse return a http response
func NewHTTPResponse(code int, msg string, data interface{}) *HTTPResponse {
	return &HTTPResponse{
		Code:    code,
		Message: msg,
		Data:    data,
	}
}

// httpResponse return a http response
func httpResponse(requestID string, code int, msg string, data interface{}) *HTTPResponse {
	return &HTTPResponse{
		RequestID: requestID,
		Code:      code,
		Message:   msg,
		Data:      data,
	}
}

// ParseAcceptLanguage Parse Accept Language
func ParseAcceptLanguage(acceptLanguage string) string {

	langQStrs := strings.Split(acceptLanguage, ",")
	for _, langQStr := range langQStrs {
		trimedLangQStr := strings.Trim(langQStr, " ")

		langQ := strings.Split(trimedLangQStr, ";")
		if len(langQ) == 1 {
			return langQ[0]
		}
		qp := strings.Split(langQ[1], "=")
		if len(qp) < 2 {
			continue
		}
		if qp[1] == "1" {
			return langQ[0]
		}
	}
	return ""
}

// GetLang Get Lang
func GetLang(c *gin.Context) string {
	lang := c.GetString("lang")
	if len(lang) == 0 {
		lang = c.GetHeader("lang")
	}
	if len(lang) == 0 {
		lang = ParseAcceptLanguage(c.GetHeader("Accept-Language"))
	}
	return lang
}

// HTTPRequestFailed HTTP Request Failed
func HTTPRequestFailed(c *gin.Context, err error, code int, data ...interface{}) {
	var errCode int
	var errData []interface{}
	errCode, errData = _err.GetCodeAndData(err)
	if errCode == 0 {
		errCode = code
		errData = data
	}
	if errCode == 0 {
		errCode = 500
	}
	msg := _err.GetErrMsg(errCode, GetLang(c))

	if len(errData) > 0 && (errCode != 500 && errCode != 400) && strings.Contains(msg, "%") {
		msg = fmt.Sprintf(msg, errData...)
		msg = strings.TrimSpace(msg)
	}
	// c.Writer.WriteHeader(http.StatusOK)
	c.JSON(http.StatusOK, httpResponse(trace.GetReqID(c.Request.Context()), errCode, msg, ""))

}

// HTTPRequestSuccess HTTP Request Success
func HTTPRequestSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, httpResponse(trace.GetReqID(c.Request.Context()), http.StatusOK,
		_err.GetSuccessMsg(http.StatusOK, GetLang(c)), data))
}
