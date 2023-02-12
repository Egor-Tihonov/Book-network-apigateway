package config

import "github.com/spf13/viper"

type Config struct {
	Port        string `mapstructure:"PORT"`
	AuthService string `mapstructure:"AUTH_SVC_URL"`
	BookService string `mapstructure:"BOOKS_SVC_URL"`
	MailService string `mapstructure:"MAIL_SVC_URL"`
	User        string `mapstructure:"USER_SVC_URL"`
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./pkg/config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)

	return
}
