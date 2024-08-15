package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var (
	errorNotExist = errors.New("given attrybute type or atribute does not exists")
)

func parseAttribute(str string) (string, []string, error) {
	tagre := regexp.MustCompile(`\.(.*?)\.`)
	matches := tagre.FindAllString(str, 1)
	if len(matches) == 0 {
		return "", nil, errors.New("no attribute type provided")
	}
	attrType := matches[0]
	attrType = strings.Trim(attrType, ". ")
	if attrType == "" {
		return "", nil, errorNotExist
	}

	attrre := regexp.MustCompile(`<(.*?)>`)
	attributes := attrre.FindAllString(str, -1)
	if len(attributes) == 0 {
		return "", nil, errors.New("no attribute provided")
	}

	attr := strings.Trim(attributes[0], "< >")
	if attr == "" {
		return "", nil, errorNotExist
	}
	attributes = strings.Split(attr, " ")

	return attrType, attributes, nil
}

func parseAttributes(attributesStr []string) (map[string]Attribute, error) {
	if len(attributesStr) == 0 {
		return map[string]Attribute{}, nil
	}
	attrParsed := map[string]Attribute{}
	for _, attr := range attributesStr {
		attrType, attributes, err := parseAttribute(attr)
		if err != nil {
			return nil, err
		}
		values, err := defineAttributeTagAndValue(attrType, attributes)
		if err != nil {
			return nil, err
		}
		attrParsed[attrType] = values
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
