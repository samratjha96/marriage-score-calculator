package model

type InitializationConfig struct {
	PlayerNames []string `yaml:"players"`
	RoundNums   int      `yaml:"rounds"`
}
