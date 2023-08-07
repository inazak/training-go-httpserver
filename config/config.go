package config

import (
  "github.com/caarlos0/env/v6"
)

type Config struct {
  Env  string   `env:"TODO_ENV"    envDefault:"dev"`
  Port int      `env:"PORT"        envDefault:"18080"`
  DBPath string `env:"TODO_DBPATH" envDefault:"./todo.db"`
}

func New() (*Config, error) {
  cfg := &Config{}
  if err := env.Parse(cfg); err != nil {
    return nil, err
  }
  return cfg, nil
}

