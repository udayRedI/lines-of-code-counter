package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/udayRedI/lines-of-code-counter/core"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter file path to analyse (Hit enter to exit): ")
		scanner.Scan()
		filePath := scanner.Text()

		if len(filePath) == 0 {
			break
		}

		flatFiles := map[string][]string{}

		ffsErr := flattenFileStructure(filePath, &flatFiles)
		if ffsErr != nil {
			log.Printf("Failed to process %s", filePath)
			continue
		}

		for _, files := range flatFiles {
			for _, file := range files {
				processor := core.NewFileProcessor()
				processErr := processor.Process(file)
				if processErr == nil {
					processor.Report()
				}
			}
		}

	}

}

func flattenFileStructure(inputPath string, flatFiles *map[string][]string) error {

	file, openErr := os.Open(inputPath)
	if openErr != nil {
		log.Print("Cant open file")
		return openErr
	}

	info, statErr := os.Stat(file.Name())
	if statErr != nil {
		log.Print("Invalid file or folder")
		return statErr
	}

	if info.IsDir() {
		entries, err := os.ReadDir(file.Name())
		if err != nil {
			log.Fatal(err)
		}

		for _, e := range entries {

			if strings.HasPrefix(e.Name(), ".") {
				// Avoid hidden folders
				continue
			}
			absPath, _ := filepath.Abs(filepath.Join(file.Name(), e.Name()))
			flattenFileStructure(absPath, flatFiles)
		}

	} else {
		parentFolder := filepath.Dir(file.Name())
		(*flatFiles)[parentFolder] = append((*flatFiles)[parentFolder], inputPath)
	}

	return nil
}
