package main

import (
	"fmt"
	"gistie/src/internal/config"
	"gistie/src/internal/file"
	"gistie/src/internal/gist"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: gister ./main.go\n")
		fmt.Printf("ERROR: No path provided\n")
	}

	path := os.Args[1]
	if !file.Exists(path) {
		fmt.Printf("ERROR: File '%s' doesn't exists\n", path)
	}

	filename, content, err := file.Read(path)
	if err != nil {
		fmt.Printf("ERROR: Can't read file '%s'\n%s\n", path, err.Error())
	}

	token := config.GetToken()

	url, err := gist.Create(token, filename, "", content)
	if err != nil {
		fmt.Printf("ERROR: Can't create a Gist: %s\n", err.Error())
	}

	fmt.Printf("You can found your Gist here -> %s\n", url)
}
