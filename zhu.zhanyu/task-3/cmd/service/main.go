package main

import (
	"flag"
	"fmt"
	"sort"

	"github.com/zhuzhanyu-der/task-3/conf"
	"github.com/zhuzhanyu-der/task-3/output"
	"github.com/zhuzhanyu-der/task-3/rates"
)

func execute() error {
	path := flag.String("config", "", "path to the YAML configuration file")
	flag.Parse()
	if *path == "" {
		return conf.ErrNoPaths
	}

	opts, err := conf.Read(*path)
	if err != nil {
		return fmt.Errorf("service: read config: %w", err)
	}

	list, err := rates.Extract(opts.Source)
	if err != nil {
		return fmt.Errorf("service: extract rates: %w", err)
	}

	sort.SliceStable(list, func(i, j int) bool {
		return list[i].Price > list[j].Price
	})

	if err := output.Write(opts.Target, list); err != nil {
		return fmt.Errorf("service: write json: %w", err)
	}
	return nil
}

func main() {
	if err := execute(); err != nil {
		panic(err)
	}
}
