// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package errors

import (
	"strings"
)

const (
	// LangZh 中文
	LangZh = "zh"
	// LangEn 英文
	LangEn = "en"
	// LangJp 日文
	// LangJp = "jp"
)

var (
	// ZhMsg 中文错误提示
	defaultZhMsg = map[int]string{
		SUCCESS:            "SUCCESS",
		StatusUnauthorized: "用户未登录",
		ERROR:              "系统异常，请联系管理员",
		ErrInvalidParams:   "请求参数错误. %v %v",
		// 用户错误 11000-11999
		ErrPasswordRule:    "密码不符合规范,最少1个大写字母,最少1个小写字母,最少一个数字,最少一个特殊字符,最少长度8位",
		ErrHashPassword:    "密码加密失败.",
		ErrUserNameExisted: "用户名已经存在,请更换.",
		ErrUserNameIsZero:  "用户、密码不能为空",
		ErrUserNotExisted:  "用户不存在",
		ErrUserPwd:         "用户密码错误",
		// 基金 12000-12999
		ErrUserFundExisted:    "基金已存在",
		ErrUserFundNotExisted: "基金不存在",
		ErrEditUserFundData:   "基金代码,id,用户id不能为空",
		ErrFundCode:           "基金代码不正确",
	}
	// EnMsg 英文错误提示
	defaultEnMsg = map[int]string{
		SUCCESS:            "Succeed",
		ERROR:              "The system is abnormal, please contact the administrator",
		ErrInvalidParams:   "Request parameter error. %v %v",
		StatusUnauthorized: "User is not authenticated",
		// 用户错误 11000-11999
		ErrPasswordRule:    "Password does not meet the specifications, at least 1 uppercase letter, at least 1 lowercase letter, at least one digit, at least one special character, at least 8 digits in length",
		ErrHashPassword:    "Password encryption failed.",
		ErrUserNameExisted: "Username already exists, please replace.",
		ErrUserNameIsZero:  "User and password cannot be empty",
		ErrUserNotExisted:  "User does not exist",
		ErrUserPwd:         "User password is wrong",
		// 基金 12000-12999
		ErrUserFundExisted:    "Fund already exists",
		ErrUserFundNotExisted: "Fund does not exist",
		ErrEditUserFundData:   "Fund code, id, userID cannot be empty",
		ErrFundCode:           "Fund code is incorrect",
	}

	defaultErrorMsg = map[string]map[int]string{
		LangZh: defaultZhMsg,
		LangEn: defaultEnMsg,
	}
	// ErrorMsg Error Msg
	ErrorMsg = map[string]map[int]string{
		LangZh: defaultZhMsg,
		LangEn: defaultEnMsg,
	}
)

// GetErrMsg Get Err Msg
func GetErrMsg(code int, lang string) string {
	return getMsg(code, lang, ERROR)
}

// GetSuccessMsg Get Success Msg
func GetSuccessMsg(code int, lang string) string {
	return getMsg(code, lang, SUCCESS)
}

// getMsg get error information based on Code
func getMsg(code int, lang string, status int) string {
	var msg string
	lang = strings.ToLower(lang)
	if lang == "" {
		lang = LangZh
	}
	errorMsg, ok := ErrorMsg[lang]
	if !ok {
		errorMsg = defaultErrorMsg[LangZh]
	}
	if message, ok := errorMsg[code]; ok {
		msg = message
	} else {
		msg = errorMsg[status]
	}
	return msg

}
