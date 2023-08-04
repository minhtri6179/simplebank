package util

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver        string        `mapstructure:"DB_Driver"`
	DBSource        string        `mapstructure:"DB_Source"`
	ServerIPAddress string        `mapstructure:"Server_Address"`
	TokenSymmetric  string        `mapstructure:"TOKEN_ASYMMENTRIC_KEY"`
	AccessTokenTTL  time.Duration `mapstructure:"Access_Token_TTL"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
