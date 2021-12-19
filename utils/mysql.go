package utils

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
	"web/global"
)

type InitDBService struct{}

type InitDB struct {
	Host     string `json:"host"`                        // 服务器地址
	Port     string `json:"port"`                        // 数据库连接端口
	UserName string `json:"userName" binding:"required"` // 数据库用户名
	Password string `json:"password"`                    // 数据库密码
	DBName   string `json:"dbName" binding:"required"`   // 数据库名
}

func Dsn() string {
	//var configs string
	config := global.GlobalConf.Mysqldb.Username + ":" + global.GlobalConf.Mysqldb.Password +
		"@tcp(" + global.GlobalConf.Mysqldb.Host + ":" + global.GlobalConf.Mysqldb.Port + ")/" +
		global.GlobalConf.Mysqldb.Database + "?" + global.GlobalConf.Mysqldb.Conf
	return config
}

func ConfMysql() {

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       Dsn(),
		DefaultStringSize:         256,  // string 类型字段的默认长度
		DisableDatetimePrecision:  true, // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true, // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true, // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 延迟关闭数据库连接

	sqlDB, err := db.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(100)

	// SetMaxOpenConns 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)

	defer sqlDB.Close()
}
