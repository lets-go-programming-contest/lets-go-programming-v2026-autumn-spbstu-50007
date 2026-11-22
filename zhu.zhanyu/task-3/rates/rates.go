package rates

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

var (
	ErrNoCurrencies       = errors.New("xml contains no currency entries")
	ErrUnsupportedCharset = errors.New("unsupported xml charset")
	ErrInvalidNumCode     = errors.New("invalid NumCode")
	ErrInvalidValue       = errors.New("invalid Value")
)

const (
	intBase   = 10
	floatBits = 64
)

type Currency struct {
	Code     int     `json:"num_code"`
	Abridged string  `json:"char_code"`
	Price    float64 `json:"value"`
}

type quoteXML struct {
	Digits  string `xml:"NumCode"`
	Letters string `xml:"CharCode"`
	Amount  string `xml:"Value"`
}

type cursXML struct {
	XMLName xml.Name   `xml:"ValCurs"`
	Entries []quoteXML `xml:"Valute"`
}

func Extract(path string) ([]Currency, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("rates.Extract: %w", err)
	}

	defer func() { _ = file.Close() }()

	return decode(file)
}

func decode(r io.Reader) ([]Currency, error) {
	parser := xml.NewDecoder(r)
	parser.CharsetReader = charsetReader

	var doc cursXML
	if err := parser.Decode(&doc); err != nil {
		return nil, fmt.Errorf("rates.decode: %w", err)
	}

	if len(doc.Entries) == 0 {
		return nil, ErrNoCurrencies
	}

	list := make([]Currency, 0, len(doc.Entries))

	for _, entry := range doc.Entries {
		code, err := parseNumCode(strings.TrimSpace(entry.Digits))
		if err != nil {
			return nil, err
		}

		price, err := parseValue(strings.TrimSpace(entry.Amount))
		if err != nil {
			return nil, err
		}

		list = append(list, Currency{
			Code:     code,
			Abridged: strings.TrimSpace(entry.Letters),
			Price:    price,
		})
	}

	return list, nil
}

func parseNumCode(text string) (int, error) {
	if text == "" {
		return 0, nil
	}

	code, err := strconv.Atoi(text)
	if err != nil {
		return 0, fmt.Errorf("%w %q: %w", ErrInvalidNumCode, text, err)
	}

	return code, nil
}

func parseValue(text string) (float64, error) {
	if text == "" {
		return 0, nil
	}

	value, err := strconv.ParseFloat(strings.ReplaceAll(text, ",", "."), floatBits)
	if err != nil {
		return 0, fmt.Errorf("%w %q: %w", ErrInvalidValue, text, err)
	}

	return value, nil
}

func charsetReader(label string, input io.Reader) (io.Reader, error) {
	switch {
	case strings.EqualFold(label, "windows-1251"):
		return charmap.Windows1251.NewDecoder().Reader(input), nil
	case strings.EqualFold(label, "utf-8"):
		return input, nil
	default:
		return nil, fmt.Errorf("%w: %s", ErrUnsupportedCharset, label)
	}
}

func SortByPriceDesc(list []Currency) {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].Price > list[j].Price
	})
}
