// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package databases

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/travelliu/fund/config"
	"strings"
)

func connSQLite(db *config.DB) (*gorm.DB, error) {
	file := db.Dbname
	if !strings.HasSuffix(file, ".db") {
		file = fmt.Sprintf("%s.db", file)
	}
	if db.Dir == "" {
		db.Dir = "./db"
	}
	file = fmt.Sprintf("%s/%s", db.Dir, file)
	logger.Infof("the connect database %s", file)
	return gorm.Open(dbSQLite3, file)
}
