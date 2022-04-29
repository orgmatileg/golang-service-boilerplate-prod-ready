package config

import (
	"fmt"
	"golang_service/pkg/logger"
	"io/ioutil"
	"time"

	str2duration "github.com/xhit/go-str2duration/v2"
	"gopkg.in/yaml.v2"
)

var c *Config

func Get() *Config {
	return c
}

type Config struct {
	Env    string `yaml:"env"`
	AppURL string `yaml:"appURL"`

	PostgreMaster struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		SSLMode  string `yaml:"sslMode"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		DBName   string `yaml:"dbName"`
	} `yaml:"postgres_master"`

	PostgresSlave struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		SSLMode  string `yaml:"sslMode"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		DBName   string `yaml:"dbName"`
	} `yaml:"postgres_slave"`

	JWT struct {
		SecretAccessToken  string `yaml:"secretAccessToken"`
		SecretRefreshToken string `yaml:"secretRefreshToken"`
	} `yaml:"jwt"`

	Auth struct {
		ExpireAccessToken          string `yaml:"expireAccessToken"`
		ExpireRefreshToken         string `yaml:"expireRefreshToken"`
		ExpireAccessTokenDuration  time.Duration
		ExpireRefreshTokenDuration time.Duration
	} `yaml:"auth"`
}

// TODO: Create task to Validate all config should not empty
func ParseConfigFromENV() {
	b, err := ioutil.ReadFile("./config.yml")
	if err != nil {
		panic(fmt.Sprintf("config not found: %s", err.Error()))
	}
	err = yaml.Unmarshal(b, &c)
	if err != nil {
		logger.Errorf("ParseConfigFromENV | could not parse config: %v", err)
	}

	c.Auth.ExpireAccessTokenDuration, err = str2duration.ParseDuration(c.Auth.ExpireAccessToken)
	if err != nil {
		panic(fmt.Sprintf("config auth access token duration string not valid: %s", err.Error()))
	}

	c.Auth.ExpireRefreshTokenDuration, err = str2duration.ParseDuration(c.Auth.ExpireRefreshToken)
	if err != nil {
		panic(fmt.Sprintf("config auth refresh token duration string not valid: %s", err.Error()))
	}

}
