package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"flag"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
	"gopkg.in/yaml.v3"
)

type Config struct{
	InputFile string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
	
}

type Valute struct{
	NumCode string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Nominal float64 `xml:"Nominal"`
	Name string	`xml:"Name"`
	Value string `xml:"Value"`
	ValueFloat float64
}
type ValCurs struct{
	Valutbl []Valute `xml:"Valute"`
}
func (v *ValCurs) ValCursSOrtBy(i , j int ) bool {
		return v.Valutbl[i].ValueFloat > v.Valutbl[j].ValueFloat
}

type RezalJson struct{
	NumCode int `json:"num_code"`
	CharCode string `json:"char_code"`
    Value float64 `json:"value"`
}

func main(){
	whatApath := flag.String("config", "config.yaml", "path file")
	flag.Parse()
	readFileYaml, err := os.ReadFile(*whatApath) 
	if err != nil {
		log.Fatal(err)
	}
	var config Config
	unmarerr:= yaml.Unmarshal(readFileYaml, &config)
	if unmarerr!=nil{
		log.Fatal(unmarerr)
	}
	readFileXml, errXml := os.ReadFile(config.InputFile)
	if errXml!=nil {
		log.Fatal(errXml)
	}
	decodeOldWinEncod := xml.NewDecoder(bytes.NewReader(readFileXml))
	decodeOldWinEncod.CharsetReader = charset.NewReaderLabel
	var valCurs ValCurs 
	unmarerr = decodeOldWinEncod.Decode(&valCurs)
	if unmarerr != nil {
		log.Fatal(unmarerr)
	}
	for i, _ := range valCurs.Valutbl {
		val := valCurs.Valutbl[i].Value
		var errvalf error
		valCurs.Valutbl[i].ValueFloat, errvalf =  strconv.ParseFloat(strings.ReplaceAll(val, ",", "."), 64)
		if errvalf != nil {
			log.Fatal(errvalf)
		}
	}
	var JsonStructs []RezalJson
	sort.SliceStable(valCurs.Valutbl, valCurs.ValCursSOrtBy)
	for i := range valCurs.Valutbl{
		val := valCurs.Valutbl[i].NumCode
		intZnachNumCode, errNumCode  := strconv.ParseInt(val ,10, 32)
		if errNumCode != nil {
			log.Fatal(errNumCode)
		}
		lastJson:= RezalJson{NumCode: int(intZnachNumCode), CharCode: valCurs.Valutbl[i].CharCode, Value: valCurs.Valutbl[i].ValueFloat}
		JsonStructs = append(JsonStructs, lastJson)
	}
	
	jsonBytes, errjon := json.MarshalIndent(JsonStructs, "", "  ")
	if errjon != nil {
		log.Fatal(errjon)
	}
	resDir := filepath.Dir(config.OutputFile)
	errmkdir := os.MkdirAll(resDir, 0o775)
	if errmkdir!=nil{
		log.Fatal(errmkdir)
	}
	errWrite := os.WriteFile(config.OutputFile, jsonBytes, 0o644)
	if errWrite != nil {
		log.Fatal(errWrite)
	}
}