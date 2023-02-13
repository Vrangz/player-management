package config

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
	"github.com/ilyakaznacheev/cleanenv"
)

// Config stores a config specific to this application
type Config struct {
	DB
	Port uint `yaml:"port" env:"PORT" env-default:"8080"`
}

// DB stores a config specific to the db
type DB struct {
	Name     string `yaml:"name" env:"DB_NAME"`
	Host     string `yaml:"host" env:"DB_HOST"`
	User     string `yaml:"user" env:"DB_USER"`
	Password string `yaml:"password" env:"DB_PASSWORD"`
	Port     int    `yaml:"port" env:"DB_PORT"`
}

// Load reads config file in yaml format and also validates the structure
func Load() (cfg Config, err error) {
	if err = cleanenv.ReadConfig("config.yaml", &cfg); err != nil {
		return
	}

	if err = validator.New().Struct(&cfg); err != nil {
		return
	}

	return
}

func (cfg Config) String() string {
	b, _ := json.Marshal(cfg)
	return string(b)
}
