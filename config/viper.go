package config

import (
	"log"

	"github.com/spf13/viper"
)

type EnvVars struct {
	PORT       string `mapstructure:"PORT"`
	DB_HOST    string `mapstructure:"DB_HOST"`
	DB_USER    string `mapstructure:"DB_USER"`
	DB_PASS    string `mapstructure:"DB_PASSWORD"`
	DB_NAME    string `mapstructure:"DB_NAME"`
	DB_PORT    string `mapstructure:"DB_PORT"`
	DB_SSLMODE string `mapstructure:"DB_SSLMODE"`
	WS_PORT    string `mapstructure:"WS_PORT"`
}

func LoadConfig() (config EnvVars, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal(err)
	}

	return
}
