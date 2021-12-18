package main

import (
	"web/global"
	"web/route"
	"web/utils"
)

// 程序入口
func main() {
	utils.InitMysql()
	global.GlobalLog = utils.Zap()
	global.GlobalLog.Debug("连接成功")
	global.GlobalLog.Warn("连接成功")
	global.GlobalLog.Info("dsdsdsd")
	//global.GlobalLog.Info("连接成功")
	route.Route()
}
