package main

import (
	"os"
)

func saveFile(filename string, data []byte) error {
	return os.WriteFile(filename, data, 0644)
}
