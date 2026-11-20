package cfg

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	ErrEmptyInput  = errors.New("input-file is empty")
	ErrEmptyOutput = errors.New("output-file is empty")
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func Read(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("cfg.Read: %w", err)
	}

	var config Config

	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("cfg.Read: %w", err)
	}

	if config.InputFile == "" {
		return nil, ErrEmptyInput
	}

	if config.OutputFile == "" {
		return nil, ErrEmptyOutput
	}

	return &config, nil
}
