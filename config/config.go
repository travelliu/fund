// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package config

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/travelliu/fund/utils/logs"
	"strings"
	"sync"
)

var (
	defaultConfFile = "conf"
	conf            *Config
	confLock        sync.Mutex
	logger          = logs.NewLogger()
)

// InitConfig 初始化配置
func InitConfig(args []string, confDir, env string) (*Config, error) {
	viperConf, err := parseViper(confDir, env)
	if err != nil {
		logger.Errorf("the parseViper error %s", err)
		return nil, err
	}
	conf = viperConf
	return conf, nil
}

// GetConf 获取配置信息
func GetConf() *Config {
	return conf
}

// ParseViper Parse Viper
func parseViper(confDir, env string) (*Config, error) {
	conFileName := defaultConfFile
	if env != "" {
		conFileName = fmt.Sprintf("%s-%s", conFileName, env)
		logger.Infof("the use env %s config file %s", env, conFileName)
	}
	conf := &Config{
		Server: &Server{},
		DB:     &DB{},
	}
	viper.SetConfigName(conFileName)
	viper.AddConfigPath(confDir) // call multiple times to add many search paths
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		if strings.Contains(err.Error(), "Not Found ") {
			return conf, nil
		}
		logger.Errorf("viper.ReadInConfig() error %s", err)
		return conf, err
	}
	if err := viper.Unmarshal(conf); err != nil {
		logger.Errorf("viper.Unmarshal error %s", err)
		return conf, err
	}
	return conf, nil
}
