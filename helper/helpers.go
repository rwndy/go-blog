package helper

import (
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

func GetConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	configPath := filepath.Join(homeDir, configFileName)

	return configPath, nil
}
