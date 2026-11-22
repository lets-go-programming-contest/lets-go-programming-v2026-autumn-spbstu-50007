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

	"golang.org/x/text/encoding/charmap"
	"gopkg.in/yaml.v3"
)

const (
	dirPermissions  = 0o755
	filePermissions = 0o600
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	NumCode  int     `xml:"NumCode"`
	CharCode string  `xml:"CharCode"`
	ValueStr string  `xml:"Value"`
	Value    float64 `xml:"-"`
}

type OutputValute struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

var ErrUnknownCharset = errors.New("unknown charset")

func main() {
	configPath := flag.String("config", "config.yaml", "path to the configuration file")
	flag.Parse()

	cfg, err := loadConfig(*configPath)
	if err != nil {
		panic(fmt.Sprintf("Failed to load config: %v", err))
	}

	valCurs, err := parseXML(cfg.InputFile)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse XML: %v", err))
	}

	result, err := processData(valCurs)
	if err != nil {
		panic(fmt.Sprintf("Failed to process data: %v", err))
	}

	err = saveJSON(cfg.OutputFile, result)
	if err != nil {
		panic(fmt.Sprintf("Failed to save JSON: %v", err))
	}

	fmt.Println("File saved to:", cfg.OutputFile)
}

func loadConfig(path string) (*Config, error) {
	configData, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Config

	err = yaml.Unmarshal(configData, &cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to parse yaml: %w", err)
	}

	return &cfg, nil
}

func parseXML(path string) (*ValCurs, error) {
	xmlData, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read xml file: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(xmlData))

	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if charset == "windows-1251" {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}

		return nil, fmt.Errorf("%w: %s", ErrUnknownCharset, charset)
	}

	var valCurs ValCurs

	err = decoder.Decode(&valCurs)
	if err != nil {
		return nil, fmt.Errorf("failed to parse xml: %w", err)
	}

	return &valCurs, nil
}

func processData(valCurs *ValCurs) ([]OutputValute, error) {
	for index := range valCurs.Valutes {
		cleanValue := strings.ReplaceAll(valCurs.Valutes[index].ValueStr, ",", ".")

		val, err := strconv.ParseFloat(cleanValue, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse value %q: %w", valCurs.Valutes[index].ValueStr, err)
		}

		valCurs.Valutes[index].Value = val
	}

	sort.Slice(valCurs.Valutes, func(i, j int) bool {
		return valCurs.Valutes[i].Value > valCurs.Valutes[j].Value
	})

	result := make([]OutputValute, 0, len(valCurs.Valutes))

	for _, v := range valCurs.Valutes {
		result = append(result, OutputValute{
			NumCode:  v.NumCode,
			CharCode: v.CharCode,
			Value:    v.Value,
		})
	}

	return result, nil
}

func saveJSON(path string, data []OutputValute) error {
	dir := filepath.Dir(path)

	err := os.MkdirAll(dir, dirPermissions)
	if err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	jsonData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return fmt.Errorf("failed to encode json: %w", err)
	}

	err = os.WriteFile(path, jsonData, filePermissions)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
