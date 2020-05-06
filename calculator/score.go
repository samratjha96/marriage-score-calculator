package calculator

import (
	"fmt"
	"log"
	"marriage/model"
	"os"
)

func simpleScoring(player1 *model.Player, player2 *model.Player) {
	if player1.Score > player2.Score {
		deficit := player1.Score - player2.Score
		player1.NormalizedScore += deficit
		player2.NormalizedScore -= deficit
	}
}

func winnerScoring(player1 *model.Player, player2 *model.Player) {
	if player1.Winner {
		if player2.RoundOneCleared {
			player1.NormalizedScore += 3
			player2.NormalizedScore -= 3
		} else {
			player1.NormalizedScore += 10
			player2.NormalizedScore -= 10
		}
	}
}

func fullNormalizedScore(player1 *model.Player, player2 *model.Player) {
	simpleScoring(player1, player2)
	simpleScoring(player2, player1)
	winnerScoring(player1, player2)
	winnerScoring(player2, player1)
}

func scoreRound(round *model.Round) {
	players := round.Players
	winner := ""
	for i := 0; i < len(players); i++ {
		// If some player has already been identified as a winner in this round, unless that winner if players[i] itself
		// then we have found an extra winner and we should fail
		if players[i].Winner {
			if winner != "" && winner != players[i].Name {
				log.Fatalf("There should only be one winner per round. Found many winners in round %d", round.RoundNum)
			}
			winner = players[i].Name
		}
		for j := i + 1; j < len(round.Players); j++ {
			if i == j {
				continue
			}
			if players[i].Name == players[j].Name {
				log.Fatal("Comparing two players with same name")
			}
			// If some player has already been identified as a winner in this round, unless that winner if players[j] itself
			// then we have found an extra winner and we should fail
			if players[j].Winner {
				if winner != "" && winner != players[i].Name {
					log.Fatalf("There should only be one winner per round. Found many winners in round %d", round.RoundNum)
				}
				winner = players[j].Name
			}
			fullNormalizedScore(&players[i], &players[j])
		}
	}
	// If no winner was identified then we need to alert the user
	if winner == "" {
		log.Fatalf("There should at least be one winner in round %d", round.RoundNum)
	}
}

func ScoreGame(filename string) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		log.Fatalf("%s is not a valid game.yml file", filename)
	}
	game := model.GameConfig{}
	gameConfig, err := game.FromYaml(filename)
	if err != nil {
		log.Fatal(err)
	}
	rounds := gameConfig.Rounds
	for i := 0; i < len(rounds); i++ {
		scoreRound(&rounds[i])
	}
	for i := 0; i < len(rounds); i++ {
		players := rounds[i].Players
		for j := 0; j < len(players); j++ {
			fmt.Printf("\n%s normalized score: %d\n", players[j].Name, players[j].NormalizedScore)
		}
	}
}
