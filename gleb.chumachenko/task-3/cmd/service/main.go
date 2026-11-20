package main

import (
	"errors"
	"flag"
	"sort"

	"github.com/Sallle-1/task-3/config"
	"github.com/Sallle-1/task-3/currency"
	"github.com/Sallle-1/task-3/storage"
)

var errConfigPath = errors.New("specify a YAML file with -config")

func main() {
	path, err := readConfigPath()
	check(err)

	cfg, err := config.Load(path)
	check(err)

	var document currency.Document

	check(storage.ReadXML(cfg.InputFile, &document))
	check(document.Validate())

	sort.SliceStable(document.Currencies, func(i, j int) bool {
		return *document.Currencies[i].Value > *document.Currencies[j].Value
	})

	check(storage.WriteJSON(cfg.OutputFile, document.Currencies))
}

func readConfigPath() (string, error) {
	path := flag.String("config", "", "path to the YAML file")
	flag.Parse()

	if *path == "" {
		return "", errConfigPath
	}

	return *path, nil
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
