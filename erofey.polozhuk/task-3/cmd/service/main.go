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
	NumCode    string  `xml:"NumCode"`
	CharCode   string  `xml:"CharCode"`
	Nominal    float64 `xml:"Nominal"`
	Name       string  `xml:"Name"`
	Value      string  `xml:"Value"`
	ValueFloat float64 `xml:"-"`
}

type ValCurs struct {
	Valutbl []Valute `xml:"Valute"`
}

type RezalJSON struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func parseNumCode(raw string) int {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return 0
	}

	code, err := strconv.ParseInt(raw, 10, 32)
	if err != nil {
		panic(err)
	}

	return int(code)
}

func main() {
	whatApath := flag.String("config", "config.yaml", "path file")
	flag.Parse()

	readFileYaml, err := os.ReadFile(*whatApath)
	if err != nil {
		panic(err)
	}

	var config Config

	unmarerr := yaml.Unmarshal(readFileYaml, &config)
	if unmarerr != nil {
		panic(unmarerr)
	}

	readFileXML, errXML := os.ReadFile(config.InputFile)
	if errXML != nil {
		panic(errXML)
	}

	decodeOldWinEncod := xml.NewDecoder(bytes.NewReader(readFileXML))
	decodeOldWinEncod.CharsetReader = charset.NewReaderLabel

	var valCurs ValCurs

	unmarerr = decodeOldWinEncod.Decode(&valCurs)
	if unmarerr != nil {
		panic(unmarerr)
	}

	for index := range valCurs.Valutbl {
		val := valCurs.Valutbl[index].Value

		var errvalf error

		valCurs.Valutbl[index].ValueFloat, errvalf = strconv.ParseFloat(strings.ReplaceAll(val, ",", "."), 64)
		if errvalf != nil {
			panic(errvalf)
		}
	}

	jsonStructs := make([]RezalJSON, 0, len(valCurs.Valutbl))

	sort.SliceStable(valCurs.Valutbl, func(left, right int) bool {
		return valCurs.Valutbl[left].ValueFloat > valCurs.Valutbl[right].ValueFloat
	})

	for index := range valCurs.Valutbl {
		lastJSON := RezalJSON{
			NumCode:  parseNumCode(valCurs.Valutbl[index].NumCode),
			CharCode: valCurs.Valutbl[index].CharCode,
			Value:    valCurs.Valutbl[index].ValueFloat,
		}
		jsonStructs = append(jsonStructs, lastJSON)
	}

	jsonBytes, errjon := json.MarshalIndent(jsonStructs, "", "  ")
	if errjon != nil {
		panic(errjon)
	}

	resDir := filepath.Dir(config.OutputFile)

	errmkdir := os.MkdirAll(resDir, dirPerm)
	if errmkdir != nil {
		panic(errmkdir)
	}

	errWrite := os.WriteFile(config.OutputFile, jsonBytes, filePerm)
	if errWrite != nil {
		panic(errWrite)
	}
}