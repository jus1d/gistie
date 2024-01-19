package config

import (
	"bufio"
	"errors"
	"fmt"
	"gistie/src/internal/file"
	"os"
	"strings"
)

const (
	FileName = ".gistie.conf"
)

func GetToken() string {
	path := fmt.Sprintf("%s/%s", os.Getenv("HOME"), FileName)
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		_, err = os.Create(path)
		if err != nil {
			return ""
		}
	}

	_, content, err := file.Read(path)
	if err != nil {
		return ""
	}

	var token string
	if content == "" {
		fmt.Printf("Enter your GitHub API token: ")
		reader := bufio.NewReader(os.Stdin)
		token, _ = reader.ReadString('\n')
		os.WriteFile(path, []byte(token), 777)
	} else {
		token = content
	}

	return strings.ReplaceAll(token, "\n", "")
}
