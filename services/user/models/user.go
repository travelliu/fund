// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package models

import (
	"github.com/travelliu/fund/utils/databases"
)

// Users 用户定义
type Users struct {
	databases.Model
	// ID       int64  `json:"id" gorm:"PRIMARY_KEY"`                         // 用户iD
	Username string `json:"username",gorm:"type:varchar(30);unique_index"` // 用户名
	Nickname string `json:"nickname",gorm:"type:varchar(30)"`              // 昵称
	Password string `json:"-",gorm:"type:varchar(200)"`                    // 密码
}

// UserPost 创建用户/登录用户
type UserPost struct {
	// ,regex=^[a-zA-Z]*$
	Username string `json:"username" validate:"required,min=3,max=40,username"` // 用户名
	Password string `json:"password" validate:"required"`                       // 密码
}

// TableName 表名
func (u *Users) TableName() string {
	return "users"
}
