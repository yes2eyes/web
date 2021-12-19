package services

type Services struct {
	ServiceName string `mapstructure: "service_name" json:"service_name" yaml:"service_name"`
	Port        int    `mapstructure: "port" json:"port" yaml:"port"`
	Env         string `mapstructure: "env" json:"env" yaml:"env"`
}
