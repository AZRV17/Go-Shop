package config

import "github.com/spf13/viper"

type Config struct {
	HTTP struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	}
	Mongo struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	}
}

func New(configPath string) (*Config, error) {
	config := &Config{}

	viper.SetConfigFile(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	config.HTTP.Host = viper.GetString("http.host")
	config.HTTP.Port = viper.GetString("http.port")
	config.Mongo.User = viper.GetString("mongo.user")
	config.Mongo.Password = viper.GetString("mongo.password")

	return config, nil
}
