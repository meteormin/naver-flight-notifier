package config

import "github.com/spf13/viper"

type Config struct {
	Naver Naver
}

type Naver struct {
	URL string
}

type notifier struct {
	Discord Discord
}

type Discord struct {
	WebHookURL string
}

func New(in string) (*Config, error) {
	viper.SetConfigFile(in)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	cfg := Config{}
	viper.Unmarshal(&cfg)

	return &cfg, err
}
