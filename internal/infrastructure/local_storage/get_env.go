package localstorage

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	WorkDir string `env:"WORKDIR" required:true`
}

func GetConfig() (*Config, error) {
	godotenv.Load()
	cfg, err := env.ParseAs[Config]()
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
