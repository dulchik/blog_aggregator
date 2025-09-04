package config

import (
	"os"
	"encoding/json"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
		DBURL 					string `json:"db_url"`
		CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	cfg := Config{}

	full_path, err := getConfigFilePath()
	if err != nil {
		return cfg, err
	} 

	dat, err := os.ReadFile(full_path)
	if err != nil {
		return cfg, err
	}
	
	if err = json.Unmarshal(dat, &cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}

func getConfigFilePath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homePath, configFileName), nil
}

func (cfg *Config) SetUser(name string) error {
	cfg.CurrentUserName = name

	fullPath, err := getConfigFilePath()
	if err != nil {
		return err
	}


	byt, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	if err := os.WriteFile(fullPath, byt, 0644); err != nil {
		return err
	}

	return nil
}


