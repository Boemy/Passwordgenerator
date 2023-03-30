package readjson

import (
	"encoding/json"
	"io/ioutil"
)

// Define a structure for the JSON file
type Config struct {
	DBLogin []Login `json:"db_login"`
}

// Define a structure for the nested DB_Login in the JSON file
type Login struct {
	DBName string `json:"db_name"`
	DBUser string `json:"db_user"`
	DBPass string `json:"db_pass"`
}

func ReadJson() (Config, error) {
	// Read the JSON file and save the output in the variable "configFile"
	configFile, err := ioutil.ReadFile("./Configuration/Config.json")
	if err != nil {
		return Config{}, err
	}

	// Unmarshal the data in the config file using the previously defined structures
	var cfg Config
	err = json.Unmarshal(configFile, &cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}
