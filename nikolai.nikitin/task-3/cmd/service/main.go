// Package main converts CBR currency rates from XML to sorted JSON.
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

	"golang.org/x/net/html/charset"
	"gopkg.in/yaml.v3"
)

const (
	dirPerm  = 0o750
	filePerm = 0o600
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type Valute struct {
	NumCode  int    `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type ValCurs struct {
	Valutes []Valute `xml:"Valute"`
}

type Currency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func readConfig(path string) Config {
	data, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		panic(err)
	}

	var cfg Config

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		panic(err)
	}

	return cfg
}

func readCurrencies(path string) []Currency {
	xmlData, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		panic(err)
	}

	var curs ValCurs

	decoder := xml.NewDecoder(bytes.NewReader(xmlData))
	decoder.CharsetReader = charset.NewReaderLabel

	err = decoder.Decode(&curs)
	if err != nil {
		panic(err)
	}

	currencies := make([]Currency, 0, len(curs.Valutes))

	for _, valute := range curs.Valutes {
		text := strings.ReplaceAll(valute.Value, ",", ".")

		number, err := strconv.ParseFloat(text, 64)
		if err != nil {
			panic(err)
		}

		currencies = append(currencies, Currency{
			NumCode:  valute.NumCode,
			CharCode: valute.CharCode,
			Value:    number,
		})
	}

	return currencies
}

func saveJSON(path string, currencies []Currency) {
	err := os.MkdirAll(filepath.Dir(path), dirPerm)
	if err != nil {
		panic(err)
	}

	result, err := json.MarshalIndent(currencies, "", "    ")
	if err != nil {
		panic(err)
	}

	result = append(result, '\n')

	err = os.WriteFile(path, result, filePerm)
	if err != nil {
		panic(err)
	}
}

func main() {
	configPath := flag.String("config", "config.yaml", "path to config file")
	flag.Parse()

	cfg := readConfig(*configPath)
	currencies := readCurrencies(cfg.InputFile)

	sort.Slice(currencies, func(i int, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})

	saveJSON(cfg.OutputFile, currencies)
}
