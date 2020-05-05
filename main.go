package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type GameConfig struct {
	Rounds []struct {
		Round   int `yaml: "round"`
		Players []struct {
			Name            string `yaml:"name"`
			Score           int    `yaml:"score"`
			Winner          bool   `yaml:"winner"`
			RoundOneCleared bool   `yaml:"pachayo"`
		} `yaml: "players"`
	} `yaml:"rounds"`
}

func (g *GameConfig) Parse(data []byte) error {
	return yaml.Unmarshal(data, g)
}

func main() {
	data, err := ioutil.ReadFile("game.yml")
	if err != nil {
		log.Fatal(err)
	}
	var config GameConfig
	if err := config.Parse(data); err != nil {
		log.Fatal(err)
	}
	d, err := yaml.Marshal(config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("%s\n\n", string(d))
}
