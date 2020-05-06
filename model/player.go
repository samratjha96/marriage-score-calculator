package model

type Player struct {
	Name            string `yaml:"name"`
	Score           int    `yaml:"score"`
	Winner          bool   `yaml:"winner"`
	RoundOneCleared bool   `yaml:"pachayo"`
	NormalizedScore int    `yaml:"-"`
}
