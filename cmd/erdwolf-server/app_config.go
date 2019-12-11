package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

const CONFIG_APPLICATION = "application"

func loadConfig(v interface{}, name string) error {
	if _, err := toml.DecodeFile(fmt.Sprintf("config/%s.toml", name), v); err != nil {
		return err
	}

	return nil
}