package conf

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

var ErrNoPaths = errors.New("config file must define both input-file and output-file")

type Settings struct {
	Source string `yaml:"input-file"`
	Target string `yaml:"output-file"`
}

func Read(path string) (Settings, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return Settings{}, fmt.Errorf("conf.Read: %w", err)
	}

	var opts Settings
	if err := yaml.Unmarshal(content, &opts); err != nil {
		return Settings{}, fmt.Errorf("conf.Read: %w", err)
	}

	opts.Source = strings.TrimSpace(opts.Source)
	opts.Target = strings.TrimSpace(opts.Target)
	if opts.Source == "" || opts.Target == "" {
		return Settings{}, ErrNoPaths
	}

	return opts, nil
}
