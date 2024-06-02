package constants

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type UserConfig struct {
	CONTRACT                string  `yaml:"CONTRACT"`
	INPUTDATA               string  `yaml:"INPUT_DATA"`
	VALUE                   int64   `yaml:"VALUE"`
	GASTARGET               float64 `yaml:"GAS_TARGET"`
	RPC                     string  `yaml:"RPC"`
	SLEEPBETWEENACCOUNTWORK []int   `yaml:"SLEEP_BETWEEN_ACCOUNT_WORK"`
	WORKERSCOUNT            int     `yaml:"WORKERS_COUNT"`
	SCAN                    string  `yaml:"SCAN"`
}

func ReadSettings(filepath string) UserConfig {
	data, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal("Error reading config file:", err)
		return UserConfig{}
	}

	var config UserConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		log.Fatal("Error decoding YAML:", err)
		return UserConfig{}
	}

	return config
}

var SETTINGS = ReadSettings("configs/general.yaml")

const LOGO = `
░██████╗░█████╗░░█████╗░███████╗██╗  ██████╗░███╗░░██╗░█████╗░
██╔════╝██╔══██╗██╔══██╗██╔════╝██║  ██╔══██╗████╗░██║██╔══██╗
╚█████╗░██║░░██║██║░░╚═╝█████╗░░██║  ██║░░██║██╔██╗██║███████║
░╚═══██╗██║░░██║██║░░██╗██╔══╝░░██║  ██║░░██║██║╚████║██╔══██║
██████╔╝╚█████╔╝╚█████╔╝██║░░░░░██║  ██████╔╝██║░╚███║██║░░██║
╚═════╝░░╚════╝░░╚════╝░╚═╝░░░░░╚═╝  ╚═════╝░╚═╝░░╚══╝╚═╝░░╚═╝
`
