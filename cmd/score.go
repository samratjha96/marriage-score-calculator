package cmd

import (
	"fmt"
	"log"
	"marriage/model"
	"os"
	"sort"

	"github.com/spf13/cobra"
)

var ScoreFile string

var scoreCmd = &cobra.Command{
	Use:   "score",
	Short: "Score the game",
	Run: func(cmd *cobra.Command, args []string) {
		ScoreGame(ScoreFile)
	},
}

func init() {
	rootCmd.AddCommand(scoreCmd)
	scoreCmd.Flags().StringVarP(&ScoreFile, "game", "g", defaultGameFile, "yml file representing the game")
}

func ScoreGame(filename string) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		log.Fatalf("%s does not exist. Please create a valid game.yml file", filename)
	}
	game := model.GameConfig{}
	gameConfig, err := game.FromYaml(filename)
	if err != nil {
		log.Fatal(err)
	}
	rounds := gameConfig.Rounds
	for i := 0; i < len(rounds); i++ {
		validateWinnerInRound(&rounds[i])
		zeroOutNonRoundOneCleared(&rounds[i])
		scoreRound(&rounds[i])
	}
	aggregatedCount := aggregateNormalizedScores(rounds)
	printResults(aggregatedCount)
}

// Every round must contain only one winner. Otherwise the configuration is invalid and we exit
func validateWinnerInRound(round *model.Round) error {
	winner := 0
	players := round.Players
	for i := 0; i < len(players); i++ {
		if players[i].Winner {
			winner += 1
		}
		if winner > 1 {
			break
		}
	}
	if winner == 0 {
		log.Fatalf("No winner found for round %d\n", round.RoundNum)
	} else if winner != 1 {
		log.Fatalf("More than one winner found for round %d\n", round.RoundNum)
	}
	return nil
}

// If a player has not shown 3 sequences when the round ends, the game rules state that they earned 0 points for the round
// Ideally the score keeper would not have put in any points for a player that did not clear the first round
// But this function provides that safety anyway in case they mistakenly forgot to
func zeroOutNonRoundOneCleared(round *model.Round) error {
	players := round.Players
	for i := 0; i < len(players); i++ {
		if !players[i].RoundOneCleared {
			fmt.Printf("Zeroed out score for %s in round %d as they didn't show 3 sequences\n", players[i].Name, round.RoundNum)
			players[i].Score = 0
		}
	}
	return nil
}

func scoreRound(round *model.Round) {
	players := round.Players
	for i := 0; i < len(players); i++ {
		for j := i + 1; j < len(round.Players); j++ {
			if i == j {
				continue
			}
			if players[i].Name == players[j].Name {
				log.Fatal("Comparing two players with same name")
			}
			fullNormalizedScore(&players[i], &players[j])
		}
	}
}

// Simply compute the difference in two player's raw score
func simpleScoring(player1 *model.Player, player2 *model.Player) {
	if player1.Score > player2.Score {
		deficit := player1.Score - player2.Score
		player1.NormalizedScore += deficit
		player2.NormalizedScore -= deficit
	}
}

// Factor in the winner's multiplier in the normalized score
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

func aggregateNormalizedScores(rounds []model.Round) map[string]int {
	aggregatedCounts := make(map[string]int)
	for i := 0; i < len(rounds); i++ {
		players := rounds[i].Players
		for j := 0; j < len(players); j++ {
			_, ok := aggregatedCounts[players[j].Name]
			if ok {
				aggregatedCounts[players[j].Name] += players[j].NormalizedScore
			} else {
				aggregatedCounts[players[j].Name] = players[j].NormalizedScore
			}
		}
	}
	return aggregatedCounts
}

func printResults(aggregatedCount map[string]int) {
	fmt.Println("------------------- Final Scores ------------------------")
	type kv struct {
		Key   string
		Value int
	}

	var slices []kv
	for k, v := range aggregatedCount {
		slices = append(slices, kv{k, v})
	}

	sort.Slice(slices, func(i, j int) bool {
		return slices[i].Value > slices[j].Value
	})

	for _, kv := range slices {
		fmt.Printf("%s: %d\n", kv.Key, kv.Value)
	}
}
