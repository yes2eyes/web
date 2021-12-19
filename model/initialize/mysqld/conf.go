package mysqld

import "fmt"

type MysqlConf struct {
	Host     string `mapstructure:"host" json: "host" yaml: "host"`
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Database string `mapstructure:"database" json:"database" yaml:"database"`
	Conf     string `mapstructure:"conf" json:"conf" yaml:"conf"`
	//LimitTimeIP   int    `mapstructure:"iplimit-time" json:"iplimitTime" yaml:"iplimit-time"`
}

func (m MysqlConf) ToMysqlConfig() string {
	fmt.Println("conf", m.Conf)
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.Database + "?" + m.Conf
}

//func (m *MysqlConf) Dans() string {
//	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.Database + "?" + m.Conf
//}
