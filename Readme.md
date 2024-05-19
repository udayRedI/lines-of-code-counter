
# Lines of code counter
Generate report for the following:
1. Number of blank lines
2. Number of lines with comments
3. Number of lines with code
4. Total number of lines in the file  

## Run application
1. Install go 1.22.2 or higher
2. Run `go run main.go` from project root

## Run all tests
`go test ./...`

## Following are basic components
1. `fileExtractor` - File extraction is abstracted out for two reasons:
    1. Dependency inversion: `fileExtractor` implements `iFileExtractor` and any new way of getting file contents can implement this interface.
    2. Easy mocking: While writing unit tests for the client of `fileExtractor` success and failure testcases can be simulated.
2. `interpreterFactory` - Abstracts away multiple types rules needed to count lines. This makes more sense when counting multiple line comments.
3. `fileProcessor` - This struct has `fileExtractor`. For now this is an composition relationship as there's no one using `iFileExtractor` interface but some other client does it can be a aggregation. `fileProcessor` uses instantiates a `interpreterFactory` to loop through every line and increment count of necessary fields. This also uses `GetInterterByLanguage` based on extension to instantiate `interpreterFactory`.

## Test coverage:
`fileProcessor` is unit tested by mocking `FileExtractorMock`. Mock helps in simulating positive and negative test cases.