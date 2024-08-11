package main

import "regexp"

type metaAttrTypes int

const (
	single metaAttrTypes = iota
	list
	mapping
)

var attributesWithSlices = map[string]struct{}{
	"class": {},
}

var attributesWithSingleValue = map[string]struct{}{
	"href": {},
}

const (
	paragraph = "paragraph"
)

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
	Metatype  metaAttrTypes
}

type AMParser struct {
	regex *regexp.Regexp
}

func NewAmParser() *AMParser {
	return &AMParser{
		regex: regexp.MustCompile(`(<-\w+-)|(-\w+->)|([^<>-]+)`),
	}
}
