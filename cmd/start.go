/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"log"
	"marriage/configuration"

	"github.com/spf13/cobra"
)

var Config string
var Out string

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a game",
	Long: `Reads in the default or provided configuration file and generates a game.yml file

The game.yml file contains the layout of how the game will be scored. The game is based on rounds
and the players playing the game. Who the players are and how many rounds there will be is determined
by the config file

Use the game.yml file to track the scoring throughout the rounds.`,
	Run: func(cmd *cobra.Command, args []string) {
		generateGame(Config)
	},
}

func generateGame(configFilePath string) {
	if err := configuration.ValidateConfigPath(configFilePath); err != nil {
		log.Fatal(err)
	}
	cfg, err := configuration.MarshalConfigFile(configFilePath)
	if err != nil {
		log.Fatal(err)
	}
	game := configuration.GenerateGameConfig(cfg, Out)
	fmt.Printf("Game file: %s successfully created\n", game.Filename)
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringVarP(&Config, "file", "f", "config.yml", "Initial configuration file")
	startCmd.Flags().StringVarP(&Out, "out", "o", "game.yml", "Generated yml file representing game configuration")
}
