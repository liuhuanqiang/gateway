package config

import (
	"os"

	"github.com/smmit/smmbase/config_reader"
	"github.com/smmit/smmbase/logger"
)

var Cfg Config

type Mysql struct {
	Host     string `json:"Host"`
	Port     string `json:"Port"`
	UserName string `json:"UserName"`
	Password string `json:"Password"`
	Schema   string `json:"schema"`
	MaxConn  int    `json:"MaxConn"`
	MaxIdle  int    `json:"MaxIdle"`
}

type Config struct {
	Mysql Mysql `json:"mysql"`
}

func LoadConf() {
	var err error

	if err = config_reader.ReadConfig(&Cfg); err != nil {
		logger.Error("Config load failed:", err)
		os.Exit(0)
	}

	return
}
