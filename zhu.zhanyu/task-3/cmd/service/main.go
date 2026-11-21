package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/zhuzhanyu-der/task-3/conf"
	"github.com/zhuzhanyu-der/task-3/output"
	"github.com/zhuzhanyu-der/task-3/rates"
)

func execute(args []string) error {
	flags := flag.NewFlagSet("cbr-rates", flag.ContinueOnError)
	configPath := flags.String("config", "", "path to YAML config")

	if err := flags.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return nil
		}

		return fmt.Errorf("service: parse flags: %w", err)
	}

	if *configPath == "" {
		return conf.ErrNoPaths
	}

	opts, err := conf.Read(*configPath)
	if err != nil {
		return fmt.Errorf("service: read config: %w", err)
	}

	list, err := rates.Extract(opts.Source)
	if err != nil {
		return fmt.Errorf("service: extract rates: %w", err)
	}

	rates.SortByPriceDesc(list)

	if err := output.WriteJSON(opts.Target, list); err != nil {
		return fmt.Errorf("service: write json: %w", err)
	}

	return nil
}

func main() {
	if err := execute(os.Args[1:]); err != nil {
		panic(err)
	}
}
