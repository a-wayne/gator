package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sync"
)

const configFileName = ".gatorconfig.json"

var configMu sync.Mutex

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUsername string `json:"current_user_name"`
}

func Read() (Config, error) {
	configMu.Lock()
	defer configMu.Unlock()

	configFilePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	jsonFile, err := os.Open(configFilePath)
	if err != nil {
		return Config{}, fmt.Errorf("could not open config file: %s", configFileName)
	}
	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		return Config{}, errors.New("could not decode configuration file")
	}

	return config, nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", errors.New("could not get home dir")
	}

	return homeDir + "/" + configFileName, nil
}

func write(cfg Config) error {
	configMu.Lock()
	defer configMu.Unlock()

	configFilePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	jsonFile, err := os.Create(configFilePath)
	if err != nil {
		return fmt.Errorf("could not open config file: %s", configFileName)
	}
	defer jsonFile.Close()

	encoder := json.NewEncoder(jsonFile)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}

	return nil
}

func (cfg *Config) SetUser(name string) error {
	cfg.CurrentUsername = name
	err := write(*cfg)
	return err
}
