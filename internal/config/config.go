package config

import (
	"encoding/json"
	"os"

	"github.com/rwndy/go-blog/helper"
)

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	configPath, err := helper.GetConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	readFile, err := os.ReadFile(configPath)

	if err != nil {
		return Config{}, err
	}

	var cfg Config
	if err := json.Unmarshal(readFile, &cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func (cfg *Config) SetUser(name string) error {
	cfg.CurrentUserName = name

	return write(*cfg)
}

func write(cfg Config) error {
	configPath, err := helper.GetConfigFilePath()
	if err != nil {
		return err
	}

	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	err = os.WriteFile(configPath, data, 0600)
	if err != nil {
		return err
	}

	return nil
}
