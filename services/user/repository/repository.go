// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package repository

import (
	"github.com/jinzhu/gorm"
	_user "github.com/travelliu/fund/services/user"
	_userMod "github.com/travelliu/fund/services/user/models"
)

// NewUserRepository New User Repository
func NewUserRepository(db *gorm.DB) _user.Repository {
	u := &user{db: db}
	u.SyncDB()
	return u
}

type user struct {
	db *gorm.DB
}

func (u *user) SetDB(db *gorm.DB) {
	u.db = db
}

func (u *user) SyncDB() {
	u.db.AutoMigrate(&_userMod.Users{})
}
