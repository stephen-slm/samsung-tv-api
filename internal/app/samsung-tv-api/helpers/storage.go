package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"
)

type Config struct {
	Token     string    `json:"token"`
	Mac       string    `json:"mac"`
	UpdatedAt time.Time `json:"updated_at"`
}

func getConfigurationFilePath() string {
	dirname, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	return path.Join(dirname, ".samsung-remote-tv.json")
}

func LoadConfiguration() Config {
	file := getConfigurationFilePath()

	var config Config
	configFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 700)
	defer configFile.Close()

	if err != nil {
		fmt.Println(err.Error())
		return Config{}
	}

	jsonParser := json.NewDecoder(configFile)

	decodeErr := jsonParser.Decode(&config)
	if decodeErr != nil {
		return Config{}
	}

	return config
}

func SaveConfiguration(config *Config) error {
	config.UpdatedAt = time.Now()
	file, err := json.MarshalIndent(&config, "", "\t")

	if err != nil {
		return err
	}

	return ioutil.WriteFile(getConfigurationFilePath(), file, 700)
}
