package file

import (
	"bufio"
	"os"
	"path/filepath"
)

func Read(path string) (filename string, content string, err error) {
	file, err := os.Open(path)
	if err != nil {
		return "", "", err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}

	return filepath.Base(file.Name()), content, nil
}

func Exists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}
