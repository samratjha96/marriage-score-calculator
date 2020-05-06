package configuration

import (
	"fmt"
	"marriage/model"
	"os"

	"gopkg.in/yaml.v2"
)

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

// Marshals the config.yml or user supplied file into a InitializationConfig struct
func MarshalConfigFile(configPath string) (*model.InitializationConfig, error) {
	// Create InitializationConfig structure
	config := &model.InitializationConfig{}

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
