// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	_fund "github.com/travelliu/fund/services/fund"
	_fundMod "github.com/travelliu/fund/services/fund/models"
	"github.com/travelliu/fund/utils/logs"
)

var (
	logger *logrus.Logger
)

func init() {
	logger = logs.NewLogger()
}

// NewFundRepository New Fund Repository
func NewFundRepository(db *gorm.DB) _fund.Repository {
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
	u.db.AutoMigrate(&_fundMod.Fund{},
		&_fundMod.UserFund{})
}
