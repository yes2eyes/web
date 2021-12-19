package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"web/model"
)

var (
	GlobalMy    *gorm.DB
	GlobalLog   *zap.Logger
	GlobalViper *viper.Viper
	GlobalConf  model.Server
	//GlobalMysqlConf  utils.Dsn1()
)
