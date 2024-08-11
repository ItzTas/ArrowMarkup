package main

import "regexp"

type NodeAM struct {
	Text       string
	Tag        string
	Attributes map[string]Attribute
}

type Attribute struct {
	TypeAttr  string
	ValueAttr string
	MappAtrr  map[string]string
	ListAttr  []string
}

type AMParser struct {
	regex *regexp.Regexp
}

func NewAmParser() *AMParser {
	return &AMParser{
		regex: regexp.MustCompile(`(<-\w+-)|(-\w+->)|([^<>-]+)`),
	}
}
