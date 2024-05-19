package core

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
)

type iFileExtractor interface {
	Extract(fileName string) ([]string, *string, error)
}

type fileExtractor struct {
}

func (f *fileExtractor) Extract(fileName string) ([]string, *string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, nil, errors.New("failed to open file")
	}

	extension := path.Ext(file.Name())
	extension = strings.Trim(extension, ".")

	defer file.Close()

	// Create a new scanner to read the file
	scanner := bufio.NewScanner(file)

	fileContents := []string{}
	// Read the file line by line
	for scanner.Scan() {
		line := scanner.Text()
		fileContents = append(fileContents, line)
	}

	// Check for errors while reading the file
	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return fileContents, &extension, nil
}

type FileProcessor struct {
	CodeCount    uint32
	CommentCount uint32
	BlankCount   uint32
	TotalCount   uint32
	fileName     string

	fileExtractor iFileExtractor
}

func (j *FileProcessor) Process(fileName string) error {
	fileContents, extension, fileErr := j.fileExtractor.Extract(fileName)
	if fileErr != nil {
		return fileErr
	}
	interpreter, interErr := GetInterterByLanguage(*extension)
	if interErr != nil {
		return interErr
	}

	j.TotalCount = uint32(len(fileContents))
	j.fileName = fileName
	for _, line := range fileContents {
		if interpreter.getBlankLineRule().IsValid(line) {
			j.BlankCount++
		} else if interpreter.getSingleCommentLineRule().IsValid(line) {
			j.CommentCount++
		} else {
			j.CodeCount++
		}
	}
	return nil
}

func (j *FileProcessor) Report() {
	fmt.Printf("File Name: %s\n", j.fileName)
	fmt.Printf("Code Count: %d\n", j.CodeCount)
	fmt.Printf("Comment Count: %d\n", j.CommentCount)
	fmt.Printf("Blank Count: %d\n", j.BlankCount)
	fmt.Printf("Total Count: %d\n", j.TotalCount)
	fmt.Printf("\n")
	fmt.Printf("\n")
}

func NewFileProcessor() *FileProcessor {
	fileExt := fileExtractor{}
	return &FileProcessor{
		fileExtractor: &fileExt,
	}
}
