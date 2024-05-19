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

func GetInterterByLanguage(language string) (interpreterFactory, error) {
	if language == "java" {
		return &javaInterpreterFactory{}, nil
	}
	return nil, errors.New("invalid language")
}
