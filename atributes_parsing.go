package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var (
	errorNotExist = errors.New("given tag or atribute does not exists")
)

func parseAttribute(str string) (string, []string, error) {
	tagre := regexp.MustCompile(`\.(.*?)\.`)
	matches := tagre.FindAllString(str, 1)
	if len(matches) == 0 {
		return "", nil, errors.New("no tag provided")
	}
	tag := matches[0]
	attrre := regexp.MustCompile(`<(.*?)>`)

	attributes := attrre.FindAllString(str, -1)
	if len(attributes) == 0 {
		return "", nil, errors.New("no attribute provided")
	}
	attr := strings.Trim(attributes[0], "< >")
	attributes = strings.Split(attr, " ")

	tag = strings.Trim(tag, ".")
	return tag, attributes, nil
}

func parseAttributes(attributesStr []string) (map[string]Attribute, error) {
	if len(attributesStr) == 0 {
		return map[string]Attribute{}, nil
	}
	attrParsed := map[string]Attribute{}
	for _, attr := range attributesStr {
		tag, attributes, err := parseAttribute(attr)
		if err != nil {
			return nil, err
		}
		attribute, err := defineAttributeTagAndValue(tag, attributes)
		if err != nil {
			return nil, err
		}
		attrParsed[tag] = attribute
	}
	return attrParsed, nil
}

func defineAttributeTagAndValue(typeAttr string, attributes []string) (Attribute, error) {
	if len(attributes) == 0 {
		return Attribute{}, errors.New("no values provided")
	}
	_, exists := attributesWithSlices[typeAttr]
	if exists {
		return Attribute{
			TypeAttr: typeAttr,
			ListAttr: attributes,
			Metatype: list,
		}, nil
	}
	_, exists = attributesWithSingleValue[typeAttr]
	if exists {
		return Attribute{
			TypeAttr:  typeAttr,
			ValueAttr: attributes[0],
			Metatype:  single,
		}, nil
	}
	return Attribute{}, fmt.Errorf("%v, with atribute: %v", errorNotExist, typeAttr)
}
