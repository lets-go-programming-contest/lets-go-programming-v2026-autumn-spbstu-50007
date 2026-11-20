package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"go.yaml.in/yaml/v3"
	"golang.org/x/text/encoding/charmap"
)

const (
	outputDirMode  = 0o755
	outputFileMode = 0o600
)

var (
	errEmptyPaths          = errors.New("input-file and output-file must be non-empty")
	errUnsupportedEncoding = errors.New("unsupported XML encoding")
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

type xmlDocument struct {
	XMLName    xml.Name      `xml:"ValCurs"`
	Currencies []xmlCurrency `xml:"Valute"`
}

type currency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func main() {
	configPath := flag.String("config", "", "path to YAML config")
	flag.Parse()

	if *configPath == "" {
		panic("specify a YAML file with -config")
	}

	cfg, err := readConfig(*configPath)
	if err != nil {
		panic(err)
	}

	currencies, err := readCurrencies(cfg.InputFile)
	if err != nil {
		panic(err)
	}

	sort.SliceStable(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})

	if err := writeCurrencies(cfg.OutputFile, currencies); err != nil {
		panic(err)
	}
}

func readConfig(path string) (config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return config{}, fmt.Errorf("read config %q: %w", path, err)
	}

	var cfg config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return config{}, fmt.Errorf("decode config %q: %w", path, err)
	}

	cfg.InputFile, cfg.OutputFile = strings.TrimSpace(cfg.InputFile), strings.TrimSpace(cfg.OutputFile)

	if cfg.InputFile == "" || cfg.OutputFile == "" {
		return config{}, errEmptyPaths
	}

	return cfg, nil
}

func readCurrencies(path string) ([]currency, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read XML %q: %w", path, err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = func(name string, input io.Reader) (io.Reader, error) {
		if strings.EqualFold(name, "windows-1251") {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}

		return nil, fmt.Errorf("%w: %s", errUnsupportedEncoding, name)
	}

	var doc xmlDocument
	if err := decoder.Decode(&doc); err != nil {
		return nil, fmt.Errorf("decode XML %q: %w", path, err)
	}

	result := make([]currency, 0, len(doc.Currencies))

	for _, item := range doc.Currencies {
		value, err := strconv.ParseFloat(
			strings.ReplaceAll(strings.TrimSpace(item.Value), ",", "."),
			64,
		)
		if err != nil {
			return nil, fmt.Errorf("parse value for %s: %w", item.CharCode, err)
		}

		result = append(result, currency{
			NumCode:  item.NumCode,
			CharCode: strings.TrimSpace(item.CharCode),
			Value:    value,
		})
	}

	return result, nil
}

func writeCurrencies(path string, currencies []currency) error {
	data, err := json.MarshalIndent(currencies, "", "    ")
	if err != nil {
		return fmt.Errorf("encode JSON: %w", err)
	}

	directory := filepath.Dir(path)
	if err := os.MkdirAll(directory, outputDirMode); err != nil {
		return fmt.Errorf("create output directory: %w", err)
	}

	if err := os.WriteFile(path, data, outputFileMode); err != nil {
		return fmt.Errorf("write JSON %q: %w", path, err)
	}

	return nil
}