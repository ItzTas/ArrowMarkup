package main

import (
	"errors"
	"regexp"
)

const (
	paragraph = "paragraph"
)

const (
	errorOutOfRange = errors.New("text out of range")
)

type NodeAM struct {
	text       string
	tag        string
	attributes map[string]string
}

type AMParser struct {
	regex *regexp.Regexp
}

func NewAmParser() *AMParser {
	return &AMParser{
		regex: regexp.MustCompile(`(<-\w+-)|(-\w+->)|([^<>-]+)`),
	}
}

func (p *AMParser) parseAM(str string) (NodeAM, error) {
	texts := p.regex.FindAllString(str, -1)
	if len(texts) == 0 {
		return NodeAM{}, errors.New("invalid string")
	}
	if len(texts) == 1 {
		return NodeAM{
			text:       texts[0],
			tag:        paragraph,
			attributes: make(map[string]string),
		}, nil
	}

	for i, t := range texts {
		if t[0] == '-' && t[len(t)-2:] == "->" {
			if i+1 > len(texts) {
				return NodeAM{}, errorOutOfRange
			}
		}
	}
}
