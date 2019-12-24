package main

import (
	"gateway/config"
	"gateway/router"
	"gateway/xorm"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//加载配置文件
	config.LoadConf()

	xorm.InitMysql()

	router.StartServer()

}
