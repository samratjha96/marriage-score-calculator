package calculator

import (
	"fmt"
	"log"
	"marriage/model"
)

func computeDeficit(player1 model.Player, player2 model.Player) int {
	if player1.Name == player2.Name {
		log.Fatal("Comparing two players with same name")
	}
	if player1.IsWinner() {
		fmt.Println("Winner is", player1.Name)
	} else {
		fmt.Println("No winner")
	}
	return 0
}

func scoreRound(round model.Round) {
	players := round.Players
	for i := 0; i < len(players); i++ {
		for j := i + 1; j < len(round.Players); j++ {
			if i == j {
				continue
			}
			computeDeficit(players[i], players[j])
		}
	}
}

func ScoreGame(game model.GameConfig) {
	gameConfig, err := game.FromYaml(game.Filename)
	if err != nil {
		log.Fatal(err)
	}
	rounds := gameConfig.Rounds
	for i := 0; i < len(rounds); i++ {
		scoreRound(rounds[i])
	}
}
