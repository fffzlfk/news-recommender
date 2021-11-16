package config

import (
	"log"

	"github.com/spf13/viper"
)

type Configurations struct {
	Database DatabaseConfigurations `mapstructure:"pgsql"`
	Numbers  NumbersConfigurations  `mapstructure:"numbers"`
}

type DatabaseConfigurations struct {
	Host     string
	User     string
	Password string
	DBName   string
}

type NumbersConfigurations struct {
	PageSize       int   `mapstructure:"page_size"`
	MaxNewsNumofDB int64 `mapstructure:"max_news_num_of_db"`
}

func Init() {
	viper.SetConfigName("config.toml")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read config failed: %v", err)
	}
}

func GetDatabaseConfigurations() DatabaseConfigurations {
	var c Configurations
	err := viper.Unmarshal(&c)
	if err != nil {
		log.Fatal(err)
	}
	return c.Database
}

func GetPageSize() int {
	Init()
	var c Configurations
	err := viper.Unmarshal(&c)
	if err != nil {
		log.Fatal(err)
	}

	return c.Numbers.PageSize
}

func GetMaxNewsNumofDB() int64 {
	Init()
	var c Configurations
	err := viper.Unmarshal(&c)
	if err != nil {
		log.Fatal(err)
	}
	return int64(c.Numbers.MaxNewsNumofDB)
}
