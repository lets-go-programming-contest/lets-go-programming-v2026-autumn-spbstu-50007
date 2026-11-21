package main

import (
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
	"golang.org/x/text/transform"
	"gopkg.in/yaml.v3"
)

const (
	folderMode   = 0o755
	resultMode   = 0o644
	floatBitSize = 64
)

var (
	errInvalidConfig      = errors.New("input-file and output-file are required")
	errUnsupportedCharset = errors.New("unsupported charset")
	errEmptyValutes       = errors.New("no valutes found in xml")
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
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type ResultValute struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func loadConfig(configPath string) Config {
	cfgData, err := os.ReadFile(filepath.Clean(configPath))
	if err != nil {
		panic(err)
	}

	var cfg Config
	if err := yaml.Unmarshal(cfgData, &cfg); err != nil {
		panic(err)
	}

	if cfg.InputFile == "" || cfg.OutputFile == "" {
		panic(errInvalidConfig)
	}

	return cfg
}

func charsetReader(charset string, input io.Reader) (io.Reader, error) {
	switch strings.ToLower(strings.TrimSpace(charset)) {
	case "windows-1251", "cp1251":
		return transform.NewReader(input, charmap.Windows1251.NewDecoder()), nil
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedCharset, charset)
	}
}

func parseXML(inputPath string) ValCurs {
	xmlFile, err := os.Open(filepath.Clean(inputPath))
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = xmlFile.Close()
	}()

	decoder := xml.NewDecoder(xmlFile)
	decoder.CharsetReader = charsetReader

	var valCurs ValCurs
	if err := decoder.Decode(&valCurs); err != nil {
		panic(err)
	}

	if len(valCurs.Valutes) == 0 {
		panic(errEmptyValutes)
	}

	return valCurs
}

func convertValutes(valutes []Valute) []ResultValute {
	results := make([]ResultValute, 0, len(valutes))

	for _, valute := range valutes {
		var numCode int

		rawNumCode := strings.TrimSpace(valute.NumCode)
		if rawNumCode != "" {
			parsedNumCode, err := strconv.Atoi(rawNumCode)
			if err != nil {
				panic(err)
			}

			numCode = parsedNumCode
		}

		charCode := strings.TrimSpace(valute.CharCode)
		valStr := strings.ReplaceAll(strings.TrimSpace(valute.Value), ",", ".")

		valFloat, err := strconv.ParseFloat(valStr, floatBitSize)
		if err != nil {
			panic(err)
		}

		results = append(results, ResultValute{
			NumCode:  numCode,
			CharCode: charCode,
			Value:    valFloat,
		})
	}

	sort.SliceStable(results, func(i, j int) bool {
		return results[i].Value > results[j].Value
	})

	return results
}

func saveJSON(outputPath string, results []ResultValute) {
	outDir := filepath.Dir(outputPath)
	if err := os.MkdirAll(outDir, folderMode); err != nil {
		panic(err)
	}

	jsonData, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile(outputPath, jsonData, resultMode); err != nil {
		panic(err)
	}
}

func main() {
	configPath := flag.String("config", "config.yaml", "path to config file")
	flag.Parse()

	cfg := loadConfig(*configPath)
	valCurs := parseXML(cfg.InputFile)
	results := convertValutes(valCurs.Valutes)
	saveJSON(cfg.OutputFile, results)
}
