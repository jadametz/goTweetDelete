package config

import (
	"strings"

	"github.com/kelseyhightower/envconfig"
)

// Config is what's necessary for the app to run
type Config struct {
	AccessSecret     string `required:"true"`
	AccessToken      string `required:"true"`
	ConsumerKey      string `required:"true"`
	ConsumerSecret   string `required:"true"`
	DaysToKeep       int    `default:"30"`
	IgnoreIDs        []int64
	IgnoreSubstrings []string
	IncludeRetweets  bool   `default:"true"`
	ScreenName       string `required:"true"`
}

func (c *Config) ShouldIgnoreId(id int64) bool {
	for _, i := range c.IgnoreIDs {
		if i == id {
			return true
		}
	}
	return false
}

func (c *Config) ShouldIgnoreOnSubstrings(body string) bool {
	for _, s := range c.IgnoreSubstrings {
		if strings.Contains(body, s) {
			return true
		}
	}
	return false
}

// New returns a new Config struct
func New() (*Config, error) {
	var c Config
	if err := envconfig.Process("", &c); err != nil {
		return nil, err
	}

	return &c, nil
}
