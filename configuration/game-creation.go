package configuration

import (
	"log"
	"marriage/model"
	"os"
)

const defaultOutputFile = "generated.yml"

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// Turn a InitializationConfig struct to a GameConfig struct
func GenerateGameConfig(config *model.InitializationConfig) model.GameConfig {
	game := model.GameConfig{}
	game.Filename = defaultOutputFile
	// Create all the rounds
	for i := 0; i < config.RoundNums; i++ {
		round := model.Round{
			RoundNum: i + 1,
			Players:  PlayerStructGenerate(config.PlayerNames),
		}
		game.Rounds = append(game.Rounds, round)
	}
	if _, err := os.Stat(defaultOutputFile); err == nil {
		log.Fatalf("Game in progress. Remove %s before running the program again", defaultOutputFile)
	}
	game.ToYaml(defaultOutputFile)
	return game
}

// Generate []Player struct from player names
func PlayerStructGenerate(names []string) []model.Player {
	playerArray := make([]model.Player, len(names))
	for i := 0; i < len(names); i++ {
		newPlayer := model.Player{
			Name: names[i],
		}
		playerArray[i] = newPlayer
	}
	return playerArray
}
