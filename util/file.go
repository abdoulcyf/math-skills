package util

import (
	"io"
	"log"
	"os"
)

func ReadFileToStr(filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			// Handle the error when closing the file
			log.Printf("Error closing file: %v", err)
		}
	}(f)

	data, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
