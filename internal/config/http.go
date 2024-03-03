package config

type HTTP struct {
	Port int `env:"HTTP_PORT, default=8081"`
}
