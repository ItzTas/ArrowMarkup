package main

import "regexp"

type metaAttrTypes int

const (
	single metaAttrTypes = iota
	list
	mapping
)

var attributesWithSlices = map[string]string{
	"class": "class",
}

var attributesWithSingleValue = map[string]string{
	"href": "href",
}

var amTagValInHTMl = map[string]string{
	"hd":  "h1",
	"h2d": "h2",
	"h3d": "h3",
	"h4d": "h4",
	"h5d": "h5",
	"h6d": "h6",

	"lik":       "a",
	"paragraph": "p",
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
		regex: regexp.MustCompile(`(<-.*-)|(-.*->)`),
	}
}
