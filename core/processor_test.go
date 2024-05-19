package core

import (
	"errors"
	"testing"
)

type FileExtractorMock struct {
	fileContents []string
	extension    string
	extErr       error
}

func (f *FileExtractorMock) Extract(fileName string) ([]string, *string, error) {
	return f.fileContents, &f.extension, f.extErr
}

func TestJavaFileProcessor(t *testing.T) {

	type testCase struct {
		title                 string
		expectedFileProcessor fileProcessor
		fileExtractor         iFileExtractor
		expError              error
	}

	testCases := []testCase{
		{
			title: "should have valid count",
			expectedFileProcessor: fileProcessor{
				CodeCount:    5,
				CommentCount: 1,
				BlankCount:   2,
				TotalCount:   8,
			},
			expError: nil,
			fileExtractor: &FileExtractorMock{
				fileContents: []string{
					"",
					"",
					"// comment",
					"public class TestClass {",
					"    public void testMethod() {",
					"        System.out.println(\"Hello, World!\"); \\ test comment",
					"    }",
					"}",
				},
				extension: "java",
				extErr:    nil,
			},
		},
		{
			title: "should return error when file is invalid",
			fileExtractor: &FileExtractorMock{
				fileContents: []string{
					"",
					"",
					"// comment",
					"public class TestClass {",
					"    public void testMethod() {",
					"        System.out.println(\"Hello, World!\"); \\ test comment",
					"    }",
					"}",
				},
				extension: "java",
				extErr:    errors.New("failed to open file"),
			},
		},
		{
			title: "should return error when file extension  is invalid",
			fileExtractor: &FileExtractorMock{
				fileContents: []string{
					"",
					"",
					"// comment",
					"public class TestClass {",
					"    public void testMethod() {",
					"        System.out.println(\"Hello, World!\"); \\ test comment",
					"    }",
					"}",
				},
				extension: "gava",
				extErr:    errors.New("invalid file extension"),
			},
		},
	}

	errCheck := func(gotErr error, tc testCase) {
		if gotErr == nil {
			if tc.expError != nil && tc.fileExtractor.(*FileExtractorMock).extErr != nil {
				t.Errorf("expected no error but got error")
			}
		} else {
			if gotErr != tc.expError && gotErr != tc.fileExtractor.(*FileExtractorMock).extErr {
				t.Errorf("expected Error but doesn't match")
			}
		}
	}

	for _, tc := range testCases {
		t.Run(tc.title, func(t *testing.T) {
			processor := fileProcessor{}
			processor.fileExtractor = tc.fileExtractor
			gotErr := processor.Process("TestFile.java")
			errCheck(gotErr, tc)
			if processor.TotalCount != tc.expectedFileProcessor.TotalCount {
				t.Errorf("expected TotalCount %d, got %d", tc.expectedFileProcessor.TotalCount, processor.TotalCount)
			}
			if processor.CodeCount != tc.expectedFileProcessor.CodeCount {
				t.Errorf("expected CodeCount %d, got %d", tc.expectedFileProcessor.CodeCount, processor.CodeCount)
			}
			if processor.CommentCount != tc.expectedFileProcessor.CommentCount {
				t.Errorf("expected CommentCount %d, got %d", tc.expectedFileProcessor.CommentCount, processor.CommentCount)
			}
			if processor.BlankCount != tc.expectedFileProcessor.BlankCount {
				t.Errorf("expected BlankCount %d, got %d", tc.expectedFileProcessor.BlankCount, processor.BlankCount)
			}
		})
	}
}
