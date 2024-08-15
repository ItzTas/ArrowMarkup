package main

import (
	"errors"
	"regexp"
	"strings"
)

var (
	errorOutOfRange = errors.New("text out of range")
)

func (p *AMParser) parseAM(str string) (NodeAM, error) {
	str = strings.Trim(str, " ")
	texts := p.regex.FindAllString(str, -1)
	if len(texts) == 0 {
		return NodeAM{}, errors.New("invalid string")
	}
	if len(texts) == 1 {
		if p.isAM(texts[0]) {
			return NodeAM{}, errorOutOfRange
		}
		return NodeAM{
			Text: strings.Trim(texts[0], " "),
			Tag:  paragraph,
		}, nil
	}

	for i, t := range texts {
		if len(t) < 3 {
			continue
		}
		if t[0] == '-' && t[len(t)-2:] == "->" {
			if i+1 == len(texts) {
				return NodeAM{}, errorOutOfRange
			}
			tag := parseTag(t)
			attributes, err := getAttributes(t)
			if err != nil {
				return NodeAM{}, err
			}
			return NodeAM{
				Text:       strings.Trim(texts[i+1], " "),
				Tag:        tag,
				Attributes: attributes,
			}, nil
		}
		if t[:2] == "<-" {
			if i-1 < 0 {
				return NodeAM{}, errorOutOfRange
			}
			tag := parseTag(t)
			attributes, err := getAttributes(t)
			if err != nil {
				return NodeAM{}, err
			}
			return NodeAM{
				Text:       strings.Trim(texts[i-1], " "),
				Tag:        tag,
				Attributes: attributes,
			}, nil
		}
	}
	return NodeAM{
		Text: strings.Trim(str, " "),
		Tag:  paragraph,
	}, nil
}

func getAttributes(str string) (map[string]Attribute, error) {
	str = strings.Trim(str, " ")
	var s string
	if str[0] == '<' {
		s = str[2 : len(str)-1]
		s = strings.Trim(s, " ")
	} else {
		s = str[1 : len(str)-2]
		s = strings.Trim(s, " ")
	}
	regex := regexp.MustCompile(`\.(\w+)\.<([^>]+)>`)
	toParse := regex.FindAllString(s, -1)
	return parseAttributes(toParse)
}

func (p *AMParser) isAM(str string) bool {
	return str == p.regex.FindString(str)
}

func parseTag(str string) string {
	str = strings.Trim(str, " ")
	if str[0] == '<' {
		s := str[2 : len(str)-1]
		s = strings.Trim(s, " ")
		return strings.Split(s, " ")[0]
	}
	s := str[1 : len(str)-2]
	s = strings.Trim(s, " ")
	return strings.Split(s, " ")[0]
}
