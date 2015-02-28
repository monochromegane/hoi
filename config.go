package hoi

import (
	"encoding/json"
	"os"
)

type Config struct {
	Port         int          `json:"port"`
	Notification Notification `json:"notification"`
}

type Notification struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Host  string `json:"host"`
	Port  int    `json:"port"`
	Token string `json:"token"`
}

func Load(path string) Config {

	file, err := os.Open(path)
	if err != nil {
		return defaultConfig()
	}

	decoder := json.NewDecoder(file)
	conf := Config{}
	err = decoder.Decode(&conf)
	if err != nil {
		return defaultConfig()
	}
	return conf
}

func defaultConfig() Config {
	return Config{
		Port: 8081,
	}
}
