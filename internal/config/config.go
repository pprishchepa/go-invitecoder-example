package config

import (
	"context"
	"fmt"
	"os"

	"github.com/hashicorp/go-envparse"
	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	MaxUsersPerCode int `env:"APP_MAXUSERSPERCODE, default=1000"`

	HTTP
	Log
	PgCounters
	PgEmails01
	PgEmails02
	PgEmails03
}

func NewConfig() (Config, error) {
	f, err := os.Open(".env")
	if err != nil && !os.IsNotExist(err) {
		return Config{}, err
	}

	if f != nil {
		envs, err := envparse.Parse(f)
		if err != nil {
			return Config{}, err
		}
		for k, v := range envs {
			if err = os.Setenv(k, v); err != nil {
				return Config{}, err
			}
		}
	}

	var conf Config
	if err := envconfig.Process(context.Background(), &conf); err != nil {
		return Config{}, fmt.Errorf("parse config: %w", err)
	}

	return conf, nil
}
