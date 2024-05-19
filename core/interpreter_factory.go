package core

import "errors"

type interpreterFactory interface {
	getBlankLineRule() Rule
	getSingleCommentLineRule() Rule
}

type javaInterpreterFactory struct{}

func (j *javaInterpreterFactory) getBlankLineRule() Rule {
	return &BlankLineRule{}
}

func (j *javaInterpreterFactory) getSingleCommentLineRule() Rule {
	return &SingleCommentLineRule{commentPattern: "//"}
}

type goInterpreterFactory struct{}

func (j *goInterpreterFactory) getBlankLineRule() Rule {
	return &BlankLineRule{}
}

func (j *goInterpreterFactory) getSingleCommentLineRule() Rule {
	return &SingleCommentLineRule{commentPattern: "//"}
}

func GetInterterByLanguage(fileExtension string) (interpreterFactory, error) {
	if fileExtension == "java" {
		return &javaInterpreterFactory{}, nil
	} else if fileExtension == "go" {
		return &goInterpreterFactory{}, nil
	}
	return nil, errors.New("invalid language")
}
