package config

import (
	"github.com/spf13/viper"
)

type NSQConfig struct {
	Host      string `mapstructure:"HOST"`
	Port      string `mapstructure:"PORT"`
	Topic     string `mapstructure:"TOPIC"`
	Channel   string `mapstructure:"CHANNEL"`
	Topic2    string `mapstructure:"TOPIC2"`
	Channel2  string `mapstructure:"CHANNEL2"`
	Topic3    string `mapstructure:"TOPIC3"`
	Channel3  string `mapstructure:"CHANNEL3"`
	Topic4    string `mapstructure:"TOPIC4"`
	Channel4  string `mapstructure:"CHANNEL4"`
	Topic5    string `mapstructure:"TOPIC5"`
	Channel5  string `mapstructure:"CHANNEL5"`
	Topic6    string `mapstructure:"TOPIC6"`
	Channel6  string `mapstructure:"CHANNEL6"`
	Topic7    string `mapstructure:"TOPIC7"`
	Channel7  string `mapstructure:"CHANNEL7"`
	Topic8    string `mapstructure:"TOPIC8"`
	Channel8  string `mapstructure:"CHANNEL8"`
	Topic9    string `mapstructure:"TOPIC9"`
	Channel9  string `mapstructure:"CHANNEL9"`
	Topic10   string `mapstructure:"TOPIC10"`
	Channel10 string `mapstructure:"CHANNEL10"`
	Topic11   string `mapstructure:"TOPIC11"`
	Topic12   string `mapstructure:"TOPIC12"`
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
	Sender   SenderConfig `mapstructure:"SENDER"`
	NSQ      NSQConfig    `mapstructure:"NSQ"`
	AuthQuiz string       `mapstructure:"QUIZ"`
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
