package xorm

import (
	"fmt"
	"gateway/config"

	"github.com/go-xorm/xorm"
	"github.com/smmit/smmbase/logger"
)

var engine *xorm.Engine

func InitMysql() {
	var err error

	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s"+"?parseTime=true&loc=Local",
		config.Cfg.Mysql.UserName,
		config.Cfg.Mysql.Password,
		config.Cfg.Mysql.Host,
		config.Cfg.Mysql.Port,
		config.Cfg.Mysql.Schema)

	engine, err = xorm.NewEngine("mysql", connStr)
	if err != nil {
		logger.Error(err)
		return
	}

}
