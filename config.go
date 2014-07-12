package hoi

import (
	"encoding/json"
	"os"
)

type Config struct {
	Port string `json:"port"`
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
		Port: "8080",
	}
}
