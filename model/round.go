package model

type Round struct {
	RoundNum int      `yaml: "round"`
	Players  []Player `yaml: "players"`
}
