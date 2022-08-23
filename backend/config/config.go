package config

import (
	"log"

	"github.com/spf13/viper"
)

type configurations struct {
	Numbers numbersConfigurations `mapstructure:"numbers"`
}

type numbersConfigurations struct {
	PageSize       int   `mapstructure:"page_size"`
	MaxNewsNumofDB int64 `mapstructure:"max_news_num_of_db"`
}

var conf configurations

func Init() {
	viper.SetConfigName("config.toml")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read config failed: %v", err)
	}
	err = viper.Unmarshal(&conf)
	if err != nil {
		log.Fatal(err)
	}
}

func GetPageSize() int {
	return conf.Numbers.PageSize
}

func GetMaxNewsNumofDB() int64 {
	return conf.Numbers.MaxNewsNumofDB
}
