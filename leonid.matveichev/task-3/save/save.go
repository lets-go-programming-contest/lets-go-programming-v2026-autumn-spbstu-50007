package save

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Leonid7557/task-3/parse"
)

const (
	dirPerm  = 0o755
	filePerm = 0o600
)

func WriteJSON(path string, rates []parse.Rate) error {
	data, err := json.MarshalIndent(rates, "", "  ")
	if err != nil {
		return fmt.Errorf("save.WriteJSON: %w", err)
	}

	dir := filepath.Dir(path)

	if err := os.MkdirAll(dir, dirPerm); err != nil {
		return fmt.Errorf("save.WriteJSON: %w", err)
	}

	if err := os.WriteFile(path, data, filePerm); err != nil {
		return fmt.Errorf("save.WriteJSON: %w", err)
	}

	return nil
}
