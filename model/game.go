package model

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type GameConfig struct {
	Rounds   []Round `yaml:"rounds"`
	Filename string  `yaml:"-"`
}

func (g *GameConfig) ToYaml(filename string) error {
	d, err := yaml.Marshal(g)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, d, 0644)
	return err
}

func (g *GameConfig) FromYaml(filename string) (*GameConfig, error) {
	config := &GameConfig{}

	// Open config file
	file, err := os.Open(filename)
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
