package parse

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

var ErrEmptyRates = errors.New("no currencies found")

type Rate struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

type xmlValute struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type xmlRoot struct {
	XMLName xml.Name    `xml:"ValCurs"`
	Valutes []xmlValute `xml:"Valute"`
}

func DecodeXML(path string) ([]Rate, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("parse.DecodeXML: %w", err)
	}
	defer func() { _ = file.Close() }()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if strings.EqualFold(charset, "windows-1251") {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}
		return nil, fmt.Errorf("unsupported encoding: %s", charset)
	}

	var root xmlRoot

	if err := decoder.Decode(&root); err != nil {
		return nil, fmt.Errorf("parse.DecodeXML: %w", err)
	}

	if len(root.Valutes) == 0 {
		return nil, ErrEmptyRates
	}

	rates := make([]Rate, 0, len(root.Valutes))

	for _, valute := range root.Valutes {
		var numCode int

		if valute.NumCode != "" {
			var err error

			numCode, err = strconv.Atoi(valute.NumCode)
			if err != nil {
				return nil, fmt.Errorf("parse.DecodeXML: %w", err)
			}
		}

		valueStr := strings.ReplaceAll(valute.Value, ",", ".")

		var value float64

		if valueStr != "" {
			var err error

			value, err = strconv.ParseFloat(valueStr, 64)
			if err != nil {
				return nil, fmt.Errorf("parse.DecodeXML: %w", err)
			}
		}

		rates = append(rates, Rate{
			NumCode:  numCode,
			CharCode: strings.TrimSpace(valute.CharCode),
			Value:    value,
		})
	}

	if len(rates) == 0 {
		return nil, ErrEmptyRates
	}

	return rates, nil
}
