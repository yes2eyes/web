package main

import (
	"fmt"
	"web/global"
	"web/route"
	"web/utils"
)

// 程序入口
func main() {
	utils.InitViper()
	utils.ConfMysql()

	global.GlobalLog = utils.Zap()
	aaa := utils.Dsn()
	fmt.Println("aaa", aaa)

	conf := global.GlobalConf.Mysqldb.Conf
	fmt.Println("conf", conf)
	fmt.Println("username", global.GlobalConf.Tmp.Abc)
	fmt.Println("env:", global.GlobalConf.Service.Port)

	route.Route()
}
