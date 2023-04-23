package AppConfig

import (
	"encoding/json"
	"log"
	"os"
	"reflect"
)

var config AppConfig

type AppConfig struct {
	PORT string `config:"PORT"`
	TEST string `config:"test"`
}

func GetAppConfig() *AppConfig {
	if !reflect.ValueOf(config).IsZero() {
		return &config
	}
	file, err := os.Open("config.json")
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Panic(err)
	}
	return &config
}
