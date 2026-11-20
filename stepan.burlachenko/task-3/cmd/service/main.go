package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
	"gopkg.in/yaml.v2"
)

type FilePath struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type Valute struct {
	NumCode   string `xml:"NumCode"`
	CharCode  string `xml:"CharCode"`
	Nominal   string `xml:"Nominal"`
	Name      string `xml:"Name"`
	Value     string `xml:"Value"`
	VunitRate string `xml:"VunitRate"`
}

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valutes []Valute `xml:"Valute"`
}

type OutputStruct struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

const (
	dirPerm  = 0o755
	filePerm = 0o600
)

func parseFlags() string {
	var configFlag string

	flag.StringVar(&configFlag, "config", "config.yaml", "Enter file name")
	flag.Parse()

	return configFlag
}

func loadConfig(path string) FilePath {
	yamlFile, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var cfg FilePath

	if err := yaml.Unmarshal(yamlFile, &cfg); err != nil {
		panic(err)
	}

	return cfg
}

func parseXML(filePath string) ValCurs {
	xmlFile, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = xmlFile.Close()
	}()

	decoder := xml.NewDecoder(xmlFile)
	decoder.CharsetReader = charset.NewReaderLabel

	var valCurs ValCurs

	if err := decoder.Decode(&valCurs); err != nil {
		panic(err)
	}

	return valCurs
}

func processData(valCurs ValCurs) []OutputStruct {
	result := make([]OutputStruct, 0, len(valCurs.Valutes))

	for _, data := range valCurs.Valutes {
		numCodeInt := 0
		trimmedNumCode := strings.TrimSpace(data.NumCode)

		if trimmedNumCode != "" {
			parseCode, err := strconv.Atoi(trimmedNumCode)
			if err != nil {
				panic(err)
			}

			numCodeInt = parseCode
		}

		tempString := strings.ReplaceAll(strings.TrimSpace(data.Value), ",", ".")

		valueFloat, err := strconv.ParseFloat(tempString, 64)
		if err != nil {
			panic(err)
		}

		temp := OutputStruct{
			NumCode:  numCodeInt,
			CharCode: strings.TrimSpace(data.CharCode),
			Value:    valueFloat,
		}

		result = append(result, temp)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Value > result[j].Value
	})

	return result
}

func saveJSON(outputPath string, result []OutputStruct) {
	jsonResult, err := json.MarshalIndent(result, "", "\t")
	if err != nil {
		panic(err)
	}

	dir := filepath.Dir(outputPath)

	if err := os.MkdirAll(dir, dirPerm); err != nil {
		panic(err)
	}

	if err = os.WriteFile(outputPath, jsonResult, filePerm); err != nil {
		panic(err)
	}
}

func main() {
	configPath := parseFlags()
	cfg := loadConfig(configPath)
	valCurs := parseXML(cfg.InputFile)
	result := processData(valCurs)
	saveJSON(cfg.OutputFile, result)
}
