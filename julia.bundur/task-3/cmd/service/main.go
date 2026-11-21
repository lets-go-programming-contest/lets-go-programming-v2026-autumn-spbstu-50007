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
	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	SourcePath string `yaml:"input-file"`
	TargetPath string `yaml:"output-file"`
}

type Rate float64

func (r *Rate) UnmarshalText(text []byte) error {
	clean := strings.ReplaceAll(strings.TrimSpace(string(text)), ",", ".")

	number, err := strconv.ParseFloat(clean, 64)
	if err != nil {
		return err
	}

	*r = Rate(number)

	return nil
}

type Currency struct {
	Numeric int    `xml:"NumCode"  json:"num_code"`
	Letters string `xml:"CharCode" json:"char_code"`
	Amount  Rate   `xml:"Value"    json:"value"`
}

type Bulletin struct {
	Items []Currency `xml:"Valute"`
}

func main() {
	cfgFlag := flag.String("config", "", "path to YAML config file")
	flag.Parse()

	if *cfgFlag == "" {
		panic("flag -config is required")
	}

	settings := parseSettings(*cfgFlag)
	currencies := loadCurrencies(settings.SourcePath)

	sort.Slice(currencies, func(a, b int) bool {
		return currencies[a].Amount > currencies[b].Amount
	})

	saveCurrencies(settings.TargetPath, currencies)
}

func parseSettings(path string) AppConfig {
	raw, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var settings AppConfig
	if err := yaml.Unmarshal(raw, &settings); err != nil {
		panic(err)
	}

	return settings
}

func loadCurrencies(path string) []Currency {
	src, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer src.Close()

	dec := xml.NewDecoder(src)
	dec.CharsetReader = charset.NewReaderLabel // файлы в windows-1251

	var bulletin Bulletin
	if err := dec.Decode(&bulletin); err != nil {
		panic(err)
	}

	return bulletin.Items
}

func saveCurrencies(path string, currencies []Currency) {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		panic(err)
	}

	dst, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer dst.Close()

	enc := json.NewEncoder(dst)
	enc.SetIndent("", "    ")

	if err := enc.Encode(currencies); err != nil {
		panic(err)
	}
}
