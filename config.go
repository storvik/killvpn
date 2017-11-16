package main

import (
	"encoding/json"
	"io/ioutil"
)

// Read configuration file and return Config struct
func readConfig(cfgPath string) (*Config, error) {
	printInfo("Reading config file.\nUsing: %s\n", cfgPath)
	raw, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		return nil, err
	}

	cfg := new(Config)
	json.Unmarshal(raw, cfg)

	return cfg, nil
}
