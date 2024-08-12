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

func TestParseAttribute_MustErrorWhenNoAttributeTypeOrAttributesAreProvided(t *testing.T) {
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
			name:                 "no attributes provided",
			input:                "< no type >",
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
