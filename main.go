package main

import (
	"flag"
	"log"
	"marriage/configuration"
)

const defaultInputFile = "config.yml"

// ParseFlags will create and parse the CLI flags
// and return the path to be used elsewhere
func ParseFlags() (string, error) {
	// String that contains the configured configuration path
	var configPath string

	// Set up a CLI flag called "-config" to allow users
	// to supply the configuration file
	flag.StringVar(&configPath, "f", defaultInputFile, "path to config file")

	flag.Parse()

	// Validate the path first
	if err := configuration.ValidateConfigPath(configPath); err != nil {
		return "", err
	}

	// Return the configuration path
	return configPath, nil
}

func main() {
	// Generate our config based on the config supplied
	// by the user in the flags
	cfgPath, err := ParseFlags()
	check(err)
	cfg, err := configuration.MarshalConfigFile(cfgPath)
	check(err)
	configuration.GenerateGameConfig(cfg)
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
