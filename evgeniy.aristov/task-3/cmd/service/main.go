package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"flag"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"go.yaml.in/yaml/v3"
	"golang.org/x/net/html/charset"
)

const (
	folderMode = 0o755
	resultMode = 0o644
)

type fileOptions struct {
	Source string `yaml:"input-file"`
	Target string `yaml:"output-file"`
}

type xmlQuote struct {
	Number int    `xml:"NumCode"`
	Symbol string `xml:"CharCode"`
	Amount string `xml:"Value"`
}

type currencyQuote struct {
	Number int     `json:"num_code"`
	Symbol string  `json:"char_code"`
	Amount float64 `json:"value"`
}

func main() {
	configPath := flag.String("config", "config.yaml", "path to YAML configuration")
	flag.Parse()

	options := readOptions(*configPath)
	quotes := collectQuotes(options.Source)

	sort.SliceStable(quotes, func(left, right int) bool {
		return quotes[left].Amount > quotes[right].Amount
	})

	saveQuotes(options.Target, quotes)
}

func readOptions(filename string) fileOptions {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var options fileOptions
	if err := yaml.Unmarshal(content, &options); err != nil {
		panic(err)
	}

	if options.Source == "" || options.Target == "" {
		panic("input-file and output-file must be specified")
	}

	return options
}

func collectQuotes(filename string) []currencyQuote {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var table struct {
		XMLName xml.Name   `xml:"ValCurs"`
		Entries []xmlQuote `xml:"Valute"`
	}

	decoder := xml.NewDecoder(bytes.NewReader(content))
	decoder.CharsetReader = charset.NewReaderLabel

	if err := decoder.Decode(&table); err != nil {
		panic(err)
	}

	quotes := make([]currencyQuote, 0, len(table.Entries))

	for _, entry := range table.Entries {
		amount, err := strconv.ParseFloat(strings.ReplaceAll(entry.Amount, ",", "."), 64)
		if err != nil {
			panic(err)
		}

		quotes = append(quotes, currencyQuote{
			Number: entry.Number,
			Symbol: entry.Symbol,
			Amount: amount,
		})
	}

	return quotes
}

func saveQuotes(filename string, quotes []currencyQuote) {
	content, err := json.MarshalIndent(quotes, "", "    ")
	if err != nil {
		panic(err)
	}

	if err := os.MkdirAll(filepath.Dir(filename), folderMode); err != nil {
		panic(err)
	}

	if err := os.WriteFile(filename, content, resultMode); err != nil {
		panic(err)
	}
}
