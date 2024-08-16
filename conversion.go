package main

import (
	"fmt"
	"strings"
)

func (node *NodeAM) toHTML() (string, error) {
	htmlTag, ok := amTagValInHTMl[node.Tag]
	if !ok {
		return "", fmt.Errorf("%v, with tag: %v", errorNotExist, node.Tag)
	}
	var attributes string
	for _, a := range node.Attributes {
		value, err := a.toHTML()
		if err != nil {
			return "", err
		}
		attributes = value + " "
	}
	attributes = strings.Trim(attributes, " ")
	tagAndValue := htmlTag + " " + attributes
	return fmt.Sprintf("<%v>%v</%v>", tagAndValue, node.Text, htmlTag), nil
}

func (a *Attribute) toHTML() (string, error) {
	switch a.Metatype {
	case single:
		htmlAttrtype, ok := attributesWithSingleValue[a.TypeAttr]
		if !ok {
			return "", fmt.Errorf("%v, with attribute single type: %v", errorNotExist, a.TypeAttr)
		}
		return fmt.Sprintf("%v=\"%v\"", htmlAttrtype, a.ValueAttr), nil
	case list:
		htmlAttrtype, ok := attributesWithSlices[a.TypeAttr]
		if !ok {
			return "", fmt.Errorf("%v, with attribute list type: %v", errorNotExist, a.TypeAttr)
		}
		var valuesStr string
		for _, v := range a.ListAttr {
			valuesStr += v + " "
		}
		valuesStr = strings.Trim(valuesStr, " ")
		return fmt.Sprintf("%v=\"%v\"", htmlAttrtype, valuesStr), nil
	}
	return "", fmt.Errorf("unkown metatype: %v", a.Metatype)
}
