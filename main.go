package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/udayRedI/lines-of-code-counter/core"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter file path to analyse (Hit enter to exit): ")
		scanner.Scan()
		text := scanner.Text()
		if len(text) == 0 {
			break
		}
		processor := core.NewFileProcessor()
		processor.Process(text)
		fmt.Printf("Code Count: %d\n", processor.CodeCount)
		fmt.Printf("Comment Count: %d\n", processor.CommentCount)
		fmt.Printf("Blank Count: %d\n", processor.BlankCount)
		fmt.Printf("Total Count: %d\n", processor.TotalCount)
	}

}
