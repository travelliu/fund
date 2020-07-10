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
	r := &repo{db: db}
	r.SyncDB()
	return r
}

type repo struct {
	db *gorm.DB
}

func (r *repo) SetDB(db *gorm.DB) {
	r.db = db
}

func (r *repo) SyncDB() {
	r.db.AutoMigrate(&_fundMod.Fund{},
		&_fundMod.UserFund{},
		&_fundMod.FundHistory{},
		&_fundMod.UserFundHistory{},
	)
}
