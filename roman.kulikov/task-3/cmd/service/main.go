package main

import (
	"flag"
	"fmt"
	"sort"

	"github.com/MrMAROdev/task-3/cbr"
	"github.com/MrMAROdev/task-3/config"
	"github.com/MrMAROdev/task-3/output"
)

func run() error {
	configPath := flag.String("config", "", "path to the YAML configuration file")
	flag.Parse()

	if *configPath == "" {
		return config.ErrEmptyPaths
	}

	cfg, err := config.Load(*configPath)
	if err != nil {
		return fmt.Errorf("main: load config: %w", err)
	}

	rates, err := cbr.Parse(cfg.InputFile)
	if err != nil {
		return fmt.Errorf("main: parse xml: %w", err)
	}

	sort.SliceStable(rates, func(i, j int) bool {
		return rates[i].Value > rates[j].Value
	})

	if err := output.Write(cfg.OutputFile, rates); err != nil {
		return fmt.Errorf("main: write json: %w", err)
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}
