// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package errors

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

type errors struct {
	code    int
	message string
	data    []interface{}
}

type errorsInterface interface {
	Code() int
	Error() string
	Data() []interface{}
}

// func NewV2(code int, messages ...interface{}) error {
// 	format := GetErrMsg(code, "")
// 	return Wrap(code, nil, format, messages...)
// }

// New New error
func New(code int, format string, messages ...interface{}) error {
	return Wrap(code, nil, format, messages...)
}

// Wrap Wrap error
func Wrap(code int, err error, format string, messages ...interface{}) error {
	if err != nil {
		format = fmt.Sprintf("%s : %s", format, err.Error())
	}

	message := fmt.Sprintf(format, messages...)

	return &errors{code, message, messages}
}

// Code error Code
func (e *errors) Code() int {
	return e.code
}

// Error error message
func (e *errors) Error() string {
	return e.message
}

// Data error data
func (e *errors) Data() []interface{} {
	return e.data
}

// GetCode 找到对应错误原因代码
func GetCode(err error) int {
	type coder interface {
		Code() int
	}

	if err == nil {
		return ERROR
	}
	if c, ok := err.(coder); ok {
		return c.Code()
	}
	return ERROR

}

// GetCodeAndData Get Code And Data
func GetCodeAndData(err error) (int, []interface{}) {
	if c, ok := err.(errorsInterface); ok {
		return c.Code(), c.Data()
	}
	return 0, nil
}

func testError(err error, lang string, code int, data ...interface{}) {
	var errCode int
	var errData []interface{}
	errCode, errData = GetCodeAndData(err)
	if errCode == 0 {
		errCode = code
		errData = data
	}
	if errCode == 0 {
		errCode = 500
	}
	msg := GetErrMsg(errCode, lang)

	if len(errData) > 0 && (errCode != 500 && errCode != 400) {
		msg = fmt.Sprintf(msg, errData...)
	}
}

// ConvertErrorToString Convert Error To String
func ConvertErrorToString(err error) string {
	if err == nil {
		return ""
	}
	switch err := err.(type) {
	case *json.SyntaxError:
		return err.Error()
	case *json.UnmarshalTypeError:
		return err.Field
	case validator.ValidationErrors:
		var col []string
		for _, v := range err {
			msg := fmt.Sprintf("%s %v %v %v ", v.Field(), v.ActualTag(), v.Param(), v.Value())
			msg = strings.TrimSpace(msg)
			col = append(col, msg)
		}
		return strings.Join(col, ",")
	default:
		return err.Error()
	}
}
