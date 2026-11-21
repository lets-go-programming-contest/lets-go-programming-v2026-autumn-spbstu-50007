package output

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/zhuzhanyu-der/task-3/rates"
)

const (
	dirPerm  = 0o755
	filePerm = 0o644
)

func Write(path string, currencies []rates.Currency) error {
	data, err := json.MarshalIndent(currencies, "", "    ")
	if err != nil {
		return fmt.Errorf("output.Write: %w", err)
	}
	if err := os.MkdirAll(filepath.Dir(path), dirPerm); err != nil {
		return fmt.Errorf("output.Write: %w", err)
	}
	if err := os.WriteFile(path, data, filePerm); err != nil {
		return fmt.Errorf("output.Write: %w", err)
	}
	return nil
}
