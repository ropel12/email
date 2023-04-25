package config

import (
	"github.com/spf13/viper"
)

type NSQConfig struct {
	Host     string `mapstructure:"HOST"`
	Port     string `mapstructure:"PORT"`
	Topic    string `mapstructure:"TOPIC"`
	Channel  string `mapstructure:"CHANNEL"`
	Topic2   string `mapstructure:"TOPIC2"`
	Channel2 string `mapstructure:"CHANNEL2"`
	Topic3   string `mapstructure:"TOPIC2"`
	Channel3 string `mapstructure:"CHANNEL3"`
}

type SenderConfig struct {
	Email     string `mapstructure:"EMAIL"`
	Password  string `mapstructure:"PASSWORD"`
	Phone     string `mapstructure:"PHONE"`
	Name      string `mapstructure:"NAME"`
	Address   string `mapstructure:"ADDRESS"`
	Slogan    string `mapstructure:"SLOGAN"`
	Twitter   string `mapstructure:"TWTR"`
	Instagram string `mapstructure:"IG"`
	Facebook  string `mapstructure:"FB"`
}

type Config struct {
	Sender SenderConfig `mapstructure:"SENDER"`
	NSQ    NSQConfig    `mapstructure:"NSQ"`
}

func NewConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	viper.AutomaticEnv()

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
