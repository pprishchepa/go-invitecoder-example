package config

type Log struct {
	Level  string `env:"LOG_LEVEL, default=info"`
	Pretty bool   `env:"LOG_PRETTY, default=false"`
}
