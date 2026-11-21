package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
	"gopkg.in/yaml.v3"
)

const dirPermissions = 0o755

type AppConfig struct {
	SourcePath string `yaml:"input-file"`
	TargetPath string `yaml:"output-file"`
}

type Rate float64

func (r *Rate) UnmarshalText(text []byte) error {
	clean := strings.ReplaceAll(strings.TrimSpace(string(text)), ",", ".")

	number, err := strconv.ParseFloat(clean, 64)
	if err != nil {
		return fmt.Errorf("parse rate %q: %w", clean, err)
	}

	*r = Rate(number)

	return nil
}

type Currency struct {
	Numeric int    `json:"num_code"  xml:"NumCode"`
	Letters string `json:"char_code" xml:"CharCode"`
	Amount  Rate   `json:"value"     xml:"Value"`
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

	defer func() {
		if closeErr := src.Close(); closeErr != nil {
			panic(closeErr)
		}
	}()

	dec := xml.NewDecoder(src)
	dec.CharsetReader = func(_ string, input io.Reader) (io.Reader, error) {
		return charmap.Windows1251.NewDecoder().Reader(input), nil
	}

	var bulletin Bulletin
	if err := dec.Decode(&bulletin); err != nil {
		panic(err)
	}

	return bulletin.Items
}

func saveCurrencies(path string, currencies []Currency) {
	if err := os.MkdirAll(filepath.Dir(path), dirPermissions); err != nil {
		panic(err)
	}

	dst, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	defer func() {
		if closeErr := dst.Close(); closeErr != nil {
			panic(closeErr)
		}
	}()

	enc := json.NewEncoder(dst)
	enc.SetIndent("", "    ")

	if err := enc.Encode(currencies); err != nil {
		panic(err)
	}
}
