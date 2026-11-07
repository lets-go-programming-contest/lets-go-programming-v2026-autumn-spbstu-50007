package main

import (
	"encoding/xml"
	"flag"
	"bytes"
	"fmt"
	"os"
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
}
type ValCurs struct{
	Valutbl []Valute `xml:"Valute"`
}

func main(){
	whatAflag := flag.String("config", "config.yaml", "path file")
	flag.Parse()
	readFileYaml, err := os.ReadFile(*whatAflag) 
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
	fmt.Println(len(readFileXml))


}