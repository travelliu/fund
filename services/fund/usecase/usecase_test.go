// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package usecase

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/travelliu/fund/config"
	"github.com/travelliu/fund/utils/databases"
	"github.com/travelliu/fund/utils/trace"
	"os"
	"testing"
)

var (
	db  *gorm.DB
	ctx context.Context
)

func TestMain(m *testing.M) {
	var (
		dbConf = &config.DB{
			Type:   "sqlite",
			Dir:    fmt.Sprintf("%s/%s", os.Getenv("GOPATH"), "src/github.com/travelliu/fund/db"),
			User:   "",
			Pwd:    "",
			Host:   "",
			Port:   0,
			Dbname: "fund",
		}
		err error
	)
	ctx = trace.AttachReqID(context.Background())
	db, err = databases.InitDatabase(dbConf)
	if err != nil {
		return
	}
	m.Run()
}
