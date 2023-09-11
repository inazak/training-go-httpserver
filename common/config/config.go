package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	Env    string `env:"TODO_ENV"    envDefault:"dev"`
	Port   int    `env:"TODO_PORT"   envDefault:"18080"`
	DBPath string `env:"TODO_DBPATH" envDefault:"./todo.db"`
}

func NewConfig() (*Config, error) {
	conf := &Config{}
	if err := env.Parse(conf); err != nil {
		return nil, err
	}
	return conf, nil
}
