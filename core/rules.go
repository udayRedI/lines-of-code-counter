package core

import "strings"

type Rule interface {
	IsValid(line string) bool
}

type BlankLineRule struct {
}

func (b *BlankLineRule) IsValid(line string) bool {
	return strings.Trim(line, " ") == ""
}

type SingleCommentLineRule struct {
	commentPattern string
}

func (c *SingleCommentLineRule) IsValid(line string) bool {
	line = strings.Trim(line, " ")
	return strings.HasPrefix(line, c.commentPattern)
}
