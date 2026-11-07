package main

import (
	"flag"
	"os"
	"fmt"
	"gopkg.in/yaml.v3"
)

type Config struct{
	InputFile string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
	
}


func main(){
	whatAflag := flag.String("config", "config.yaml", "path file")
	flag.Parse()
	readFile, err := os.ReadFile(*whatAflag) 
	if err != nil {
		panic(err)
	}
	var config Config
	unmarerr:= yaml.Unmarshal(readFile, &config)
	if unmarerr!=nil{
		panic(unmarerr)
	}
	fmt.Println(config.InputFile,config.OutputFile)

}