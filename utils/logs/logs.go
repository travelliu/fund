// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package logs

import (
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"io"
	"os"
	"sync"
)

var (
	log     *logrus.Logger
	logLock sync.Mutex
)

// func init() {
// 	log = initLogRus()
// }

// InitLogRus  set logrus
func initLogRus() *logrus.Logger {
	if log != nil {
		return log
	}
	logLock.Lock()
	defer logLock.Unlock()
	log = logrus.New()
	formatter := new(prefixed.TextFormatter)
	formatter.DisableColors = true
	// formatter.FullTimestamp = true                    // 显示完整时间
	formatter.TimestampFormat = "2006-01-02 15:04:05" // 时间格式
	log.SetOutput(io.MultiWriter(os.Stdout))
	log.SetFormatter(formatter)
	log.SetReportCaller(true)
	filenameHook := NewHook()
	filenameHook.Field = "line"
	log.AddHook(filenameHook)
	log.SetLevel(logrus.DebugLevel)
	return log
}

// NewLogger New Logger
func NewLogger() *logrus.Logger {
	return initLogRus()
}
