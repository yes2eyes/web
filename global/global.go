package global

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GlobalMy  *gorm.DB
	GlobalLog *zap.Logger
)
