package config

import (
	"github.com/tkanos/gonfig"
)

type Configuration struct {
	DBUsername string `json:"db_username"`
	DBPassword string `json:"db_password"`
	DBHost     string `json:"db_host"`
	DBPort     int    `json:"db_port"`
	DBName     string `json:"db_name"`

	AccessTokenSecret  string `json:"access_token_secret"`
	AccessTokenExp     int    `json:"access_token_exp"`
	RefreshTokenSecret string `json:"refresh_token_secret"`
	RefreshTokenExp    int    `json:"refresh_token_exp"`
}

var conf Configuration

func NewConfiguration() *Configuration {
	configuration := Configuration{}

	err := gonfig.GetConf("config/config.json", &configuration)
	if err != nil {
		panic(err)
	}

	return &configuration
}

func NewConfigurationPath(path string) *Configuration {
	configuration := Configuration{}

	err := gonfig.GetConf(path, &configuration)
	if err != nil {
		panic(err)
	}

	return &configuration
}
