package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// Config is what's necessary for the app to run
type Config struct {
	AccessSecret    string `required:"true"`
	AccessToken     string `required:"true"`
	ConsumerKey     string `required:"true"`
	ConsumerSecret  string `required:"true"`
	DaysToKeep      int    `default:"30"`
	TweetCount      int    `default:"3200"`
	IncludeRetweets bool   `default:"true"`
	ScreenName      string `required:"true"`
}

// New returns a new Config struct
func New() *Config {
	var c Config
	if err := envconfig.Process("", &c); err != nil {
		log.Fatal(err.Error())
	}

	return &c
}
