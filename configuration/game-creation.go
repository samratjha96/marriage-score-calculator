package configuration

import (
	"io/ioutil"
	"log"
	"marriage/model"

	"gopkg.in/yaml.v2"
)

const defaultOutputFile = "generated.yml"

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// Turn a InitializationConfig struct to a GameConfig struct
func GenerateGameConfig(config *model.InitializationConfig) {
	game := model.GameConfig{}
	// Create all the rounds
	for i := 0; i < config.RoundNums; i++ {
		round := model.Round{
			RoundNum: i + 1,
			Players:  PlayerStructGenerate(config.PlayerNames),
		}
		game.Rounds = append(game.Rounds, round)
	}
	d, err := yaml.Marshal(game)
	check(err)
	err = ioutil.WriteFile(defaultOutputFile, d, 0644)
	check(err)
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
