// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package main

import (
	"fmt"
	"github.com/douglarek/zerodown"
	"github.com/travelliu/fund/config"
	"github.com/travelliu/fund/services"
	"github.com/travelliu/fund/utils/databases"
	"github.com/travelliu/fund/utils/logs"
	"log"
)

func main() {
	// init logger
	// init db
	// init repo
	// init services
	// init http
	logger := logs.NewLogger()
	conf, err := config.InitConfig(nil, "", "")
	if err != nil {
		logger.Errorf("the parse conf file error %s", err)
		return
	}
	logger.Infof("conf db %+v", conf.DB)
	db, err := databases.InitDatabase(conf.DB)
	if err != nil {
		logger.Errorf("the InitDatabase error %s", err)
		return
	}
	handler := services.NewService(db)
	addr := genHTTPListener(conf.Server)
	logger.Infof("the server run %s", addr)
	log.Fatalln(zerodown.ListenAndServe(addr, handler))
}

func genHTTPListener(server *config.Server) string {
	var (
		host, port = server.Host, server.Port
	)
	if host == "" {
		host = "0.0.0.0"
	}
	if port == 0 {
		port = 8081
	}
	return fmt.Sprintf("%s:%v", host, port)
}
