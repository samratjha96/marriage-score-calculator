package main

import (
	"flag"
	"log"
	"marriage/calculator"
	"marriage/configuration"
)

const defaultInputFile = "config.yml"

func main() {
	// Generate our config based on the config supplied
	// by the user in the flags

	configFilePtr := flag.String("f", defaultInputFile, "Configuration file to create the game")
	// score := flag.Bool("score", false, "Score the game")

	flag.Parse()
	if err := configuration.ValidateConfigPath(*configFilePtr); err != nil {
		check(err)
	}
	cfg, err := configuration.MarshalConfigFile(*configFilePtr)
	check(err)
	game := configuration.GenerateGameConfig(cfg)
	calculator.ScoreGame(game)
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
