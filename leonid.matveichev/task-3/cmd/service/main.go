package main

import (
	"flag"
	"fmt"
	"sort"

	"github.com/Leonid7557/task-3/cfg"
	"github.com/Leonid7557/task-3/parse"
	"github.com/Leonid7557/task-3/save"
)

func main() {
	configPath := flag.String("config", "", "path to the YAML configuration file")
	flag.Parse()

	if *configPath == "" {
		panic("main: -config flag is required")
	}

	appConfig, err := cfg.Read(*configPath)
	if err != nil {
		panic(fmt.Sprintf("main: %v", err))
	}

	rates, err := parse.DecodeXML(appConfig.InputFile)
	if err != nil {
		panic(fmt.Sprintf("main: %v", err))
	}

	sort.Slice(rates, func(i, j int) bool {
		return rates[i].Value > rates[j].Value
	})

	if err := save.WriteJSON(appConfig.OutputFile, rates); err != nil {
		panic(fmt.Sprintf("main: %v", err))
	}
}
