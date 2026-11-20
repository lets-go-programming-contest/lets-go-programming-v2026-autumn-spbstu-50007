package cbr

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

var (
	ErrNoRates     = errors.New("no currencies found")
	ErrUnsupported = errors.New("unsupported xml encoding")
)

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

func Parse(path string) ([]Rate, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("cbr.Parse: %w", err)
	}

	defer func() { _ = file.Close() }()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if strings.EqualFold(charset, "windows-1251") {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}

		return nil, fmt.Errorf("%w: %s", ErrUnsupported, charset)
	}

	var root xmlRoot
	if err := decoder.Decode(&root); err != nil {
		return nil, fmt.Errorf("cbr.Parse: %w", err)
	}

	if len(root.Valutes) == 0 {
		return nil, ErrNoRates
	}

	rates := make([]Rate, 0, len(root.Valutes))

	for _, item := range root.Valutes {
		numCode := 0

		if trimmed := strings.TrimSpace(item.NumCode); trimmed != "" {
			if n, err := strconv.Atoi(trimmed); err == nil {
				numCode = n
			}
		}

		value := 0.0

		if trimmed := strings.TrimSpace(item.Value); trimmed != "" {
			if f, err := strconv.ParseFloat(strings.ReplaceAll(trimmed, ",", "."), 64); err == nil {
				value = f
			}
		}

		rates = append(rates, Rate{
			NumCode:  numCode,
			CharCode: strings.TrimSpace(item.CharCode),
			Value:    value,
		})
	}

	if len(rates) == 0 {
		return nil, ErrNoRates
	}

	return rates, nil
}
