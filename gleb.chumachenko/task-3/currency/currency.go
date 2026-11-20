package currency

import (
	"encoding/xml"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var errIncomplete = errors.New("incomplete currency data")

type Rate float64

func (value *Rate) UnmarshalText(text []byte) error {
	decimal := strings.ReplaceAll(strings.TrimSpace(string(text)), ",", ".")

	number, err := strconv.ParseFloat(decimal, 64)
	if err != nil {
		return fmt.Errorf("parse currency value %q: %w", text, err)
	}

	*value = Rate(number)

	return nil
}

type Currency struct {
	NumCode  int    `json:"num_code"  xml:"NumCode"`
	CharCode string `json:"char_code" xml:"CharCode"`
	Value    *Rate  `json:"value"     xml:"Value"`
}

type Document struct {
	XMLName    xml.Name   `xml:"ValCurs"`
	Currencies []Currency `xml:"Valute"`
}

func (document *Document) Validate() error {
	if document.Currencies == nil {
		document.Currencies = []Currency{}
	}

	for i := range document.Currencies {
		item := &document.Currencies[i]
		item.CharCode = strings.TrimSpace(item.CharCode)

		if item.Value == nil || (item.NumCode == 0) != (item.CharCode == "") {
			return errIncomplete
		}
	}

	return nil
}
