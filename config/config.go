package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type DBConnectioConfig struct {
	Name     string `json:"name"`
	Addr     string `json:"addr"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type Configuration struct {
	Database DBConnectioConfig
}

func LoadConfig(path string) Configuration {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Config file missed:", err)
	}
	var config Configuration
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatal("Configuration parse error:", err)
	}

	return config

}
