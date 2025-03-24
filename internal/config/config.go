package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		Server   `json:"server"`
		Log      `json:"log"`
		PersonDb `json:"person-db"`
	}
	Server struct {
		Host string `env-requires:"true" json:"host"`
		Port string `env-requires:"true" json:"port"`
	}
	Log struct {
		LogFilePath string `env-requires:"true" json:"log-file-path"`
	}
	PersonDb struct {
		Dsn                string `env-requires:"true" env:"DSN"`
		MaxOpenConnections uint   `env-requires:"true" env:"MAX_OPEN_CONNECTIONS"`
	}
)

func MustLoadConfig(configPath string) *Config {
	cfg := &Config{}
	if err := cleanenv.ReadConfig(configPath, cfg); err != nil {
		log.Fatal("Can't parse config: ", err.Error())
	}
	if err := cleanenv.UpdateEnv(cfg); err != nil {
		log.Fatal("Can't update enviroment: ", err.Error())
	}
	return cfg
}
