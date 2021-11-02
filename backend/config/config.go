package config

import (
	"log"

	"github.com/spf13/viper"
)

func Init() {
	viper.SetConfigName("config.toml")
	viper.SetConfigType("toml")
	viper.AddConfigPath("/home/fffzlfk/github/news/backend")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read config failed: %v", err)
	}
	{
		MaxNewsNumofPage = viper.GetInt("number.max_news_num_of_page")
		MaxNewsNumofDB = viper.GetInt64("number.max_news_num_of_db")
		Increasement = viper.GetInt("number.increasement")
	}

	{
		Host = viper.GetString("pgsql.host")
		User = viper.GetString("pgsql.user")
		Password = viper.GetString("pgsql.password")
		DBName = viper.GetString("pgsql.dbname")
	}
}

var (
	MaxNewsNumofPage int
	MaxNewsNumofDB   int64
	Increasement     int
)

var (
	Host     string
	User     string
	Password string
	DBName   string
)
