package storage

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

const outputDirMode = 0o755

var (
	errXMLEncoding = errors.New("unsupported XML encoding")
	errExtraXML    = errors.New("unexpected data after XML root")
)

func ReadXML(path string, data any) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("open XML %q: %w", path, err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if !strings.EqualFold(charset, "windows-1251") {
			return nil, fmt.Errorf("%w %q", errXMLEncoding, charset)
		}

		return charmap.Windows1251.NewDecoder().Reader(input), nil
	}

	if err := decoder.Decode(data); err != nil {
		return fmt.Errorf("decode XML %q: %w", path, err)
	}

	var extra any

	if err := decoder.Decode(&extra); err == nil {
		return errExtraXML
	} else if !errors.Is(err, io.EOF) {
		return fmt.Errorf("decode XML %q: %w", path, err)
	}

	return nil
}

func WriteJSON(path string, data any) error {
	encoded, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return fmt.Errorf("encode JSON: %w", err)
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, outputDirMode); err != nil {
		return fmt.Errorf("create output directory: %w", err)
	}

	file, err := os.CreateTemp(dir, ".output-*.json")
	if err != nil {
		return fmt.Errorf("create output file: %w", err)
	}

	defer func() {
		if err := os.Remove(file.Name()); err != nil && !errors.Is(err, os.ErrNotExist) {
			panic(err)
		}
	}()

	_, writeErr := file.Write(encoded)
	closeErr := file.Close()

	if writeErr != nil {
		return fmt.Errorf("write JSON: %w", writeErr)
	}

	if closeErr != nil {
		return fmt.Errorf("close JSON file: %w", closeErr)
	}

	if err := os.Rename(file.Name(), path); err != nil {
		return fmt.Errorf("replace JSON file: %w", err)
	}

	return nil
}
