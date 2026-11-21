package rates

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
	ErrEmptyFile   = errors.New("xml contains no currency entries")
	ErrBadEncoding = errors.New("xml uses an unsupported charset")
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

	parser := xml.NewDecoder(file)
	parser.CharsetReader = func(label string, input io.Reader) (io.Reader, error) {
		if strings.EqualFold(label, "windows-1251") {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}
		return nil, fmt.Errorf("%w: %s", ErrBadEncoding, label)
	}

	var doc cursXML
	if err := parser.Decode(&doc); err != nil {
		return nil, fmt.Errorf("rates.Extract: %w", err)
	}
	if len(doc.Entries) == 0 {
		return nil, ErrEmptyFile
	}

	list := make([]Currency, 0, len(doc.Entries))
	for _, q := range doc.Entries {
		code, err := strconv.Atoi(strings.TrimSpace(q.Digits))
		if err != nil {
			return nil, fmt.Errorf("rates.Extract: invalid NumCode %q: %w", q.Digits, err)
		}
		price, err := strconv.ParseFloat(
			strings.ReplaceAll(strings.TrimSpace(q.Amount), ",", "."), 64,
		)
		if err != nil {
			return nil, fmt.Errorf("rates.Extract: invalid Value %q: %w", q.Amount, err)
		}
		list = append(list, Currency{
			Code:     code,
			Abridged: strings.TrimSpace(q.Letters),
			Price:    price,
		})
	}

	return list, nil
}
