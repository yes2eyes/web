package model

import (
	"web/model/initialize/mysqld"
	"web/model/initialize/services"
	"web/model/initialize/tmp"
)

type Server struct {
	Mysqldb mysqld.MysqlConf  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Service services.Services `mapstructure:"service" json:"service"yaml:"service"`
	Tmp     tmp.Tmp           `mapstructure:"tmp" json:"tmp"yaml:"tmp"`
}
