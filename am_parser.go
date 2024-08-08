package main

import (
	"errors"
	"regexp"
	"strings"
)

const (
	paragraph = "paragraph"
)

var (
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
			if i+1 == len(texts) {
				return NodeAM{}, errorOutOfRange
			}
			tag := parseTag(t)[0]
			return NodeAM{
				text:       texts[i+1],
				tag:        tag,
				attributes: make(map[string]string),
			}, nil
		}
		if t[:2] == "<-" {
			if i-1 < 0 {
				return NodeAM{}, errorOutOfRange
			}
			tag := parseTag(t)[0]
			return NodeAM{
				text:       texts[i-1],
				tag:        tag,
				attributes: make(map[string]string),
			}, nil
		}
	}
	return NodeAM{}, errors.New("unknwon error")
}

func parseTag(str string) []string {
	if str[0] == '<' {
		s := str[2 : len(str)-1]
		return strings.Split(s, " ")
	}
	s := str[1 : len(str)-2]
	return strings.Split(s, " ")
}
