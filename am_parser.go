package main

import (
	"errors"
	"regexp"
	"strings"

	"github.com/ItzTas/arrowmarkup/internal/models"
)

const (
	paragraph = "paragraph"
)

var (
	errorOutOfRange = errors.New("Text out of range")
)

type AMParser struct {
	regex *regexp.Regexp
}

func NewAmParser() *AMParser {
	return &AMParser{
		regex: regexp.MustCompile(`(<-\w+-)|(-\w+->)|([^<>-]+)`),
	}
}

func (p *AMParser) parseAM(str string) (models.NodeAM, error) {
	str = strings.Trim(str, " ")
	texts := p.regex.FindAllString(str, -1)
	if len(texts) == 0 {
		return models.NodeAM{}, errors.New("invalid string")
	}
	if len(texts) == 1 {
		if p.isAM(texts[0]) {
			return models.NodeAM{}, errorOutOfRange
		}
		return models.NodeAM{
			Text:       texts[0],
			Tag:        paragraph,
			Attributes: make(map[string]string),
		}, nil
	}

	for i, t := range texts {
		if t[0] == '-' && t[len(t)-2:] == "->" && len(t) >= 3 {
			if i+1 == len(texts) {
				return models.NodeAM{}, errorOutOfRange
			}
			tag := parseTag(t)[0]
			return models.NodeAM{
				Text:       texts[i+1],
				Tag:        tag,
				Attributes: make(map[string]string),
			}, nil
		}
		if t[:2] == "<-" && len(t) >= 3 {
			if i-1 < 0 {
				return models.NodeAM{}, errorOutOfRange
			}
			tag := parseTag(t)[0]
			return models.NodeAM{
				Text:       texts[i-1],
				Tag:        tag,
				Attributes: make(map[string]string),
			}, nil
		}
	}
	return models.NodeAM{
		Text:       str,
		Tag:        paragraph,
		Attributes: make(map[string]string),
	}, nil
}

func (p *AMParser) isAM(str string) bool {
	return str == p.regex.FindString(str)
}

func parseTag(str string) []string {
	if str[0] == '<' {
		s := str[2 : len(str)-1]
		return strings.Split(s, " ")
	}
	s := str[1 : len(str)-2]
	return strings.Split(s, " ")
}
