package mkfile

import (
	"fmt"
	"os"
)

func CreateFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
        return fmt.Errorf("creating file: %w", err)
    }
	defer file.Close()
	return nil
}