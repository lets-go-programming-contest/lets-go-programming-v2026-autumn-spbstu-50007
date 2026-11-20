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

func main() {

	var configFlag string
	flag.StringVar(&configFlag, "config", "config.yaml", "Enter file name")
	flag.Parse()

	YamlFile, err := os.ReadFile(configFlag)
	if err != nil {
		panic(err)
	}

	var path FilePath
	err = yaml.Unmarshal(YamlFile, &path)
	if err != nil {
		panic(err)
	}

	XMLFile, err := os.Open(path.InputFile)
	if err != nil {
		panic(err)
	}
	defer XMLFile.Close()

	decoder := xml.NewDecoder(XMLFile)
	decoder.CharsetReader = charset.NewReaderLabel

	var valCurs ValCurs

	err = decoder.Decode(&valCurs)
	if err != nil {
		panic(err)
	}

	var result []OutputStruct
	for _, data := range valCurs.Valutes {

		numCodeInt, err := strconv.Atoi(strings.TrimSpace(data.NumCode))
		if err != nil {
			panic(err)
		}

		tempString := strings.ReplaceAll(strings.TrimSpace(data.Value), ",", ".")
		valueInt, err := strconv.ParseFloat(tempString, 64)
		if err != nil {
			panic(err)
		}

		temp := OutputStruct{
			NumCode:  numCodeInt,
			CharCode: strings.TrimSpace(data.CharCode),
			Value:    valueInt,
		}

		result = append(result, temp)

	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Value > result[j].Value
	})

	jsonResult, err := json.MarshalIndent(result, "", "	")
	if err != nil {
		panic(err)
	}

	err = os.MkdirAll(filepath.Dir(path.OutputFile), 0755)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(path.OutputFile, jsonResult, 0644)
	if err != nil {
		panic(err)
	}
}
