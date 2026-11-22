package main
import(
    "os"
    "strings"
    "sort"
    "encoding/json"
    "encoding/xml"
    "flag"
    "fmt"
    "path/filepath"
    "strconv"
"gopkg.in/yaml.v3"
"bytes"
"golang.org/x/net/html/charset"
)

type Config struct
    {InputFile string `yaml:"input-file"`
    OutputFile string `yaml:"output-file"`
}

type FindV struct
    {XMLName xml.Name `xml:"ValCurs"`
    Items []Valute `xml:"Valute"`}


type Valute struct 
    {NCode int `xml:"NumCode"`
    CCode string `xml:"CharCode"`
    Rate string `xml:"Value"`}

type Result struct
    {NCode int `json:"num_code"`
    CCode string `json:"char_code"`
    Rate float64 `json:"value"`}


func main(){
    cfgPath := flag.String("config", "", "path to yaml config")
    flag.Parse()
    if *cfgPath == ""{
        panic("config is required")

    }

    ymlData, err := os.ReadFile(*cfgPath)
    if err != nil{
        panic(err)
    }
    var cfg Config 
    err = yaml.Unmarshal(ymlData, &cfg)
    if err != nil {
        panic(err)
    }

    xmlData, err := os.ReadFile(cfg.InputFile)
    if err != nil{
        panic(err)
    }

   var raw FindV
decoder := xml.NewDecoder(bytes.NewReader(xmlData))
decoder.CharsetReader = charset.NewReaderLabel
err = decoder.Decode(&raw)
if err != nil {
    panic(err)
}

    var results []Result 
    for _, item := range raw.Items {
        cleanRate := strings.ReplaceAll(item.Rate, ",", ".")
        val, err :=  strconv.ParseFloat(cleanRate, 64)
        if err != nil{
            panic(fmt.Sprintf("invalid rate value", item.Rate))
        }
        results = append(results, Result{
			NCode: item.NCode,
			CCode: item.CCode,
			Rate:  val,
    })
}
    sort.Slice(results, func(i, j int) bool {
		return results[i].Rate > results[j].Rate
	})

	outDir := filepath.Dir(cfg.OutputFile)
	if outDir != "." && outDir != "" {
		if err := os.MkdirAll(outDir, 0755); err != nil {
			panic(err)
		}
	}
	jsonBytes, err := json.MarshalIndent(results, "", "    ")
	if err != nil {
		panic(err)
	}

    err = os.WriteFile(cfg.OutputFile, jsonBytes, 0644)
	if err != nil {
		panic(err)
	}
} 