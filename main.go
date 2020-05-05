package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"

	"marriage/model"
)

type InitializationConfig struct {
	PlayerNames []string `yaml:"players"`
	RoundNums   int      `yaml:"rounds"`
}

func ValidateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}
	return nil
}

// InitializeConfig returns a new decoded InitializationConfig struct from a yml file
func InitializeConfig(configPath string) (*InitializationConfig, error) {
	// Create InitializationConfig structure
	config := &InitializationConfig{}

	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

// ParseFlags will create and parse the CLI flags
// and return the path to be used elsewhere
func ParseFlags() (string, error) {
	// String that contains the configured configuration path
	var configPath string

	// Set up a CLI flag called "-config" to allow users
	// to supply the configuration file
	flag.StringVar(&configPath, "f", "./config.yml", "path to config file")

	flag.Parse()

	// Validate the path first
	if err := ValidateConfigPath(configPath); err != nil {
		return "", err
	}

	// Return the configuration path
	return configPath, nil
}

// Nicely format and print the yaml from the InitializationConfig struct to the terminal
func PrettyPrintInitConfig(data []byte) {
	var config InitializationConfig
	err := yaml.Unmarshal(data, &config)
	check(err)
	d, err := yaml.Marshal(config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("%s\n\n", string(d))
}

/*
	--------------- Everything related to Game Configuration -------------------------
*/

// Turn a InitializationConfig struct to a GameConfig struct
func GenerateGameConfig(config *InitializationConfig) {
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
	err = ioutil.WriteFile("generated.yml", d, 0644)
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

// Nicely format and print the yaml from the GameConfig struct to the terminal
func PrettyPrintGameConfig(data []byte) {
	var config model.GameConfig
	err := yaml.Unmarshal(data, &config)
	check(err)
	d, err := yaml.Marshal(config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("%s\n\n", string(d))
}

func main() {
	// Generate our config based on the config supplied
	// by the user in the flags
	cfgPath, err := ParseFlags()
	check(err)
	cfg, err := InitializeConfig(cfgPath)
	check(err)
	GenerateGameConfig(cfg)
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// func main() {
// 	data, err := ioutil.ReadFile("game.yml")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	var config GameConfig
// 	if err := config.Parse(data); err != nil {
// 		log.Fatal(err)
// 	}
// 	PrettyPrint(config)
// }
