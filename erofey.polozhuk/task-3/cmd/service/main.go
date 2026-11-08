package main

import (
	"bytes"
	"encoding/xml"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sort"
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

func main(){
	whatApath := flag.String("config", "config.yaml", "path file")
	flag.Parse()
	readFileYaml, err := os.ReadFile(*whatApath) 
	if err != nil {
		panic(err)
	}
	var config Config
	unmarerr:= yaml.Unmarshal(readFileYaml, &config)
	if unmarerr!=nil{
		panic(unmarerr)
	}
	readFileXml, errXml := os.ReadFile(config.InputFile)
	if errXml!=nil {
		panic(errXml)
	}
	decodeOldWinEncod := xml.NewDecoder(bytes.NewReader(readFileXml))
	decodeOldWinEncod.CharsetReader = charset.NewReaderLabel
	var valCurs ValCurs 
	unmarerr = decodeOldWinEncod.Decode(&valCurs)
	if unmarerr != nil {
		panic(unmarerr)
	}
	fmt.Println(len(valCurs.Valutbl))
	for i, _ := range valCurs.Valutbl{
		val := valCurs.Valutbl[i].Value
		var errvalf error
		valCurs.Valutbl[i].ValueFloat, errvalf =  strconv.ParseFloat(strings.ReplaceAll(val, ",", "."), 64)
		if errvalf != nil {
			panic(errvalf)
		}
	}
	sort.SliceStable(valCurs.Valutbl, valCurs.ValCursSOrtBy)
	

}