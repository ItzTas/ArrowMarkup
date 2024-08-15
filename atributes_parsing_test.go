package main

import (
	"reflect"
	"testing"
)

// ParseAttribute
func TestParseAttribute_ValidAttribute(t *testing.T) {
	tests := []struct {
		name               string
		input              string
		expectedAttrType   string
		expectedAttrValues []string
	}{
		{
			name:               "valid class attribute",
			input:              ".class.< testclass anotherclass >",
			expectedAttrType:   "class",
			expectedAttrValues: []string{"testclass", "anotherclass"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			attrType, values, err := parseAttribute(tt.input)

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if !reflect.DeepEqual(tt.expectedAttrValues, values) {
				t.Errorf("expected values: %v, but got: %v", tt.expectedAttrValues, values)
				return
			}

			if attrType != tt.expectedAttrType {
				t.Errorf("expected attribute type: %v, but got: %v", tt.expectedAttrType, attrType)
			}
		})
	}
}

func TestParseAttribute_ExpectedErrors(t *testing.T) {
	tests := []struct {
		name                 string
		input                string
		expectedErrorMessage string
	}{
		{
			name:                 "no attribute type provided",
			input:                ". .<invalid>",
			expectedErrorMessage: errorNotExist.Error(),
		},
		{
			name:                 "no attribute type .. provided",
			input:                "< no type >",
			expectedErrorMessage: "no attribute type provided",
		},
		{
			name:                 "no attributes provided",
			input:                ".class.<>",
			expectedErrorMessage: errorNotExist.Error(),
		},
		{
			name:                 "no attributes <> provided",
			input:                ".class.",
			expectedErrorMessage: "no attribute provided",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, err := parseAttribute(tt.input)

			if err == nil {
				t.Errorf("expected error but got none")
				return
			}

			if err.Error() != tt.expectedErrorMessage {
				t.Errorf("expected error message: %v, got: %v", tt.expectedErrorMessage, err.Error())
			}
		})
	}
}

// parseAttributes
func TestParseAttributes_ValidInput(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected map[string]Attribute
	}{
		{
			name:  "valid class and href attributes",
			input: []string{".class.< testclass anotherclass >", ".href.< testlink >"},
			expected: map[string]Attribute{
				"class": {
					TypeAttr: "class",
					ListAttr: []string{"testclass", "anotherclass"},
					Metatype: list,
				},
				"href": {
					TypeAttr:  "href",
					ValueAttr: "testlink",
					Metatype:  single,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			attributes, err := parseAttributes(tt.input)
			if err != nil {
				t.Errorf("did not expect error but got: %v, with test: %v", err, tt)
			}
			if !reflect.DeepEqual(attributes, tt.expected) {
				t.Errorf("expected values did not match with test: %v, expected: %v, got: %v", tt.name, tt.expected, attributes)
			}
		})
	}
}
