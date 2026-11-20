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

func (v *ValCurs) ValCursSOrtBy(i, j int) bool {
	return v.Valutbl[i].ValueFloat > v.Valutbl[j].ValueFloat
}

type RezalJSON struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
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

	readFileXml, errXml := os.ReadFile(config.InputFile)
	if errXml != nil {
		panic(errXml)
	}

	decodeOldWinEncod := xml.NewDecoder(bytes.NewReader(readFileXml))
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

	sort.SliceStable(valCurs.Valutbl, valCurs.ValCursSOrtBy)

	for index := range valCurs.Valutbl {
		val := strings.TrimSpace(valCurs.Valutbl[index].NumCode)

		var intZnachNumCode int64

		if val != "" {
			var errNumCode error
			intZnachNumCode, errNumCode = strconv.ParseInt(val, 10, 32)
			if errNumCode != nil {
				panic(errNumCode)
			}

		}
		lastJson := RezalJSON{NumCode: int(intZnachNumCode), CharCode: valCurs.Valutbl[index].CharCode, Value: valCurs.Valutbl[index].ValueFloat}
		jsonStructs = append(jsonStructs, lastJson)
	}

	jsonBytes, errjon := json.MarshalIndent(jsonStructs, "", "  ")
	if errjon != nil {
		panic(errjon)
	}

	resDir := filepath.Dir(config.OutputFile)
	errmkdir := os.MkdirAll(resDir, 0o750)
	if errmkdir != nil {
		panic(errmkdir)
	}

	errWrite := os.WriteFile(config.OutputFile, jsonBytes, 0o600)
	if errWrite != nil {
		panic(errWrite)
	}
}
