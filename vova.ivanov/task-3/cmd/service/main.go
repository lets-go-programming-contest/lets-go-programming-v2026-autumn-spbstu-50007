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
	dirPerm  = 0o750
	filePerm = 0o600
)

type config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type xmlCurrency struct {
	NumCode  int    `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type xmlRates struct {
	XMLName xml.Name      `xml:"ValCurs"`
	Items   []xmlCurrency `xml:"Valute"`
}

type currency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func main() {
	configPath := flag.String("config", "", "path to config file")
	flag.Parse()

	if *configPath == "" {
		panic("specify a YAML file with -config")
	}

	cfg := readConfig(*configPath)
	rates := readRates(cfg.InputFile)

	sort.Slice(rates, func(i, j int) bool {
		return rates[i].Value > rates[j].Value
	})

	writeJSON(cfg.OutputFile, rates)
}

func readConfig(path string) config {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var cfg config

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		panic(err)
	}

	cfg.InputFile, cfg.OutputFile = strings.TrimSpace(cfg.InputFile), strings.TrimSpace(cfg.OutputFile)

	if cfg.InputFile == "" || cfg.OutputFile == "" {
		panic("input-file and output-file must be specified")
	}

	return cfg
}

func readRates(path string) []currency {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel

	var document xmlRates

	if err := decoder.Decode(&document); err != nil {
		panic(err)
	}

	rates := make([]currency, 0, len(document.Items))

	for _, item := range document.Items {
		valueText := strings.ReplaceAll(strings.TrimSpace(item.Value), ",", ".")

		value, err := strconv.ParseFloat(valueText, 64)
		if err != nil {
			panic(err)
		}

		rates = append(rates, currency{
			NumCode:  item.NumCode,
			CharCode: strings.TrimSpace(item.CharCode),
			Value:    value,
		})
	}

	return rates
}

func writeJSON(path string, rates []currency) {
	data, err := json.MarshalIndent(rates, "", "    ")
	if err != nil {
		panic(err)
	}

	dir := filepath.Dir(path)

	if err := os.MkdirAll(dir, dirPerm); err != nil {
		panic(err)
	}

	if err := os.WriteFile(path, data, filePerm); err != nil {
		panic(err)
	}
}