package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c *Config) SetUser(user string) {
	c.CurrentUserName = user
	write(*c)
}

func Read() (Config, error) {
	fileConfig, err := getConfigFilePath()
	var config Config

	if err != nil {
		return Config{}, err
	}

	jsonFile, err := os.Open(fileConfig)
	if err != nil {
		return Config{}, err
	}
	// Defer the close until the main function finishes
	defer jsonFile.Close()

	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		return config, err
	}

	errjson := json.Unmarshal(jsonData, &config)
	if errjson != nil {
		return config, errjson
	}

	return config, nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return homeDir + "/" + configFileName, nil
}

func write(cfg Config) error {
	jsonData, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		fmt.Printf("Error marshalling data: %v\n", err)
		return err
	}

	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return err
	}
	fmt.Println("Successfully wrote JSON data to output.json")
	return nil
}
