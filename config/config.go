package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	URL                string `json:"url"`
	MongoConnectionURL string `json:"mongo_connection_url"`
	MongoDatabase      string `json:"mongo_database"`
	MongoCollection string `json:"mongo_collection"`
}

var cachedConfig Config

func readConfig() error {
	file, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatal("Error reading config file: ", err)
		return err
	}

	err = json.Unmarshal(file, &cachedConfig)
	if err != nil {
		log.Fatal("Error parsing JSON: ", err)
		return err
	}

	return nil
}

func GetConfig() (Config, error) {
	if cachedConfig == (Config{}) {
		err := readConfig()
		if err != nil {
			return Config{}, err
		}
	}
	return cachedConfig, nil
}

func GetURL() (string, error) {
	config, err := GetConfig()
	if err != nil {
		return "", err
	}
	return config.URL, nil
}

func GetMongoURL() (string, error) {
	config, err := GetConfig()
	if err != nil {
		return "", err
	}
	return config.MongoConnectionURL, nil
}

func GetMongoDatabase() (string, error) {
	config, err := GetConfig()
	if err != nil {
		return "", err
	}
	return config.MongoDatabase, nil
}

func GetMongoCollection() (string, error) {
	config, err := GetConfig()
	if err != nil {
		return "", err
	}
	return config.MongoCollection, nil
}

