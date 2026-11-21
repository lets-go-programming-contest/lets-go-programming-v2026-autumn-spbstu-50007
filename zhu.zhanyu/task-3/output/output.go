package output

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/zhuzhanyu-der/task-3/rates"
)

const dirPerm os.FileMode = 0o755

func WriteJSON(path string, currencies []rates.Currency) error {
	if err := os.MkdirAll(filepath.Dir(path), dirPerm); err != nil {
		return fmt.Errorf("output.WriteJSON: mkdir: %w", err)
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("output.WriteJSON: create: %w", err)
	}

	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")

	if err := enc.Encode(currencies); err != nil {
		_ = file.Close()

		return fmt.Errorf("output.WriteJSON: encode: %w", err)
	}

	if err := file.Close(); err != nil {
		return fmt.Errorf("output.WriteJSON: close: %w", err)
	}

	return nil
}
