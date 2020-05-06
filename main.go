package main

import (
	"flag"
	"fmt"
	"log"
	"marriage/calculator"
	"marriage/configuration"
)

const defaultInputFile = "config.yml"
const defaultOutputFile = "generated.yml"

func main() {
	// Generate our config based on the config supplied
	// by the user in the flags

	configFilePtr := flag.String("f", defaultInputFile, "Configuration file to create the game")
	score := flag.Bool("score", false, "Score the game")

	flag.Parse()
	if err := configuration.ValidateConfigPath(*configFilePtr); err != nil {
		check(err)
	}

	if !*score {
		cfg, err := configuration.MarshalConfigFile(*configFilePtr)
		check(err)
		game := configuration.GenerateGameConfig(cfg)
		fmt.Println("Creating game config with", game.Filename)
	} else {
		fmt.Println("Scoring the game")
		calculator.ScoreGame(defaultOutputFile)
	}
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
