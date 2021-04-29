package utils

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	Engine   string
	Server   string
	Port     string
	User     string
	Password string
	Database string
}

func GetConfiguration() (Configuration, error) {
	config := Configuration{}
	file, err := os.Open("utils/configuration.json")
	if err != nil {
		return config, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}
