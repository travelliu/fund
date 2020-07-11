// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package errors

const (
	// SUCCESS 操作成功
	SUCCESS = 200
	// ERROR 系统内部错误
	ERROR = 500
	// StatusUnauthorized 未认证
	StatusUnauthorized = 401
	// ErrInvalidParams 请求参数错误
	ErrInvalidParams = 10001
)

// 用户错误 11000-11999
const (
	// ErrPasswordRule 密码错误
	ErrPasswordRule = 11001
	// ErrHashPassword 密码加密失败
	ErrHashPassword = 11002
	// ErrUserNameExisted 用户名存在
	ErrUserNameExisted = 11003
	// ErrUserNameIsZero 用户名为空
	ErrUserNameIsZero = 11004
	// ErrUserNotExisted 用户不存在
	ErrUserNotExisted = 11005
	// ErrUserPwd 用户密码失败
	ErrUserPwd = 11006
)

// 基金 12000-12999
const (
	// ErrUserFundExisted 用户基金已存在
	ErrUserFundExisted = 12001
	// ErrUserFundNotExisted 用户基金不存在
	ErrUserFundNotExisted = 12002
	// ErrEditUserFundData 修改基金数据错误
	ErrEditUserFundData = 12003
	// ErrFundCode 基金代码不正确
	ErrFundCode = 12004
)
