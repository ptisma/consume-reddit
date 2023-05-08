package initializers

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBHost            string `mapstructure:"POSTGRES_HOST"`
	DBUserName        string `mapstructure:"POSTGRES_USER"`
	DBUserPassword    string `mapstructure:"POSTGRES_PASSWORD"`
	DBName            string `mapstructure:"POSTGRES_DB"`
	DBPort            string `mapstructure:"POSTGRES_PORT"`
	ServerPort        string `mapstructure:"PORT"`
	CachePort         string `mapstructure:"REDIS_PORT"`
	CacheUserName     string `mapstructure:"REDIS_USER"`
	CacheUserPassword string `mapstructure:"REDIS_PASSWORD"`
	CacheDBName       string `mapstructure:"REDIS_DB"`
	CacheHost         string `mapstructure:"CACHE_HOST"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
