// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package databases

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/travelliu/fund/config"
	"github.com/travelliu/fund/utils/logs"
	"strings"
)

const (
	dbSQLite  = "sqlite"
	dbSQLite3 = "sqlite3"
)

var (
	logger                *logrus.Logger
	errNotSupportDatabase = "the not support database type %s"
)

func init() {
	logger = logs.NewLogger()
}

// InitDatabase 初始化数据库
func InitDatabase(dbConf *config.DB) (*gorm.DB, error) {
	var (
		db  *gorm.DB
		err error
	)
	switch strings.ToLower(dbConf.Type) {
	case dbSQLite, dbSQLite3:
		db, err = connSQLite(dbConf)
	default:
		db, err = nil, fmt.Errorf(errNotSupportDatabase, dbConf.Type)
	}
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	db.SetLogger(nopLogger{})
	return db, nil
}
