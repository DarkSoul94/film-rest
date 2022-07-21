package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Release  bool
	AppName  string
	HTTPport string
	LogDir   string
	LogFile  string
	DbPath   string
}

// InitConfig - load config from config.yml
func InitConfig() Config {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	conf := Config{
		Release:  viper.GetBool("release"),
		AppName:  viper.GetString("name"),
		HTTPport: viper.GetString("http_port"),
		LogDir:   viper.GetString("log.dir"),
		LogFile:  viper.GetString("log.file"),
		DbPath:   viper.GetString("db.path"),
	}

	return conf
}
