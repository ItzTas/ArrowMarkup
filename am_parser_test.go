package main

import (
	"reflect"
	"testing"
)

// parseAM
func TestParseAM_ValidString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected NodeAM
	}{
		{
			name:  "header with arrow",
			input: " this is a header <-hd-",
			expected: NodeAM{
				Text: "this is a header",
				Tag:  "hd",
			},
		},
		{
			name:  "header with reversed arrow",
			input: "-hd->this is a header ",
			expected: NodeAM{
				Text: "this is a header",
				Tag:  "hd",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewAmParser()
			node, err := p.parseAM(tt.input)
			if err != nil {
				t.Errorf("function did not expect an error but got: %v", err)
				return
			}

			if node.Tag != tt.expected.Tag {
				t.Errorf("expected tag: %v, got: %v", tt.expected.Tag, node.Tag)
				return
			}

			if node.Text != tt.expected.Text {
				t.Errorf("expected text: %v, got: %v", tt.expected.Text, node.Text)
			}
		})
	}
}

func TestParseAM_ExpectedErrors(t *testing.T) {
	tests := []struct {
		name                string
		input               string
		expectedErrorString string
	}{
		{
			name:                "must error when front arrow points to nothing",
			input:               "-hd->",
			expectedErrorString: errorOutOfRange.Error(),
		},
		{
			name:                "must error when inverse arrow points to nothing",
			input:               "<-hd-",
			expectedErrorString: errorOutOfRange.Error(),
		},
		{
			name:                "must error when front arrow points to nothing",
			input:               "some text -hd->",
			expectedErrorString: errorOutOfRange.Error(),
		},
		{
			name:                "must error when inverse arrow points to nothing",
			input:               "<-hd- some text",
			expectedErrorString: errorOutOfRange.Error(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewAmParser()
			_, err := p.parseAM(tt.input)
			if err == nil {
				t.Errorf("expected error but got: %v with test: %v", err, tt)
				return
			}
			if err.Error() != tt.expectedErrorString {
				t.Errorf("expected error: %v but got: %v", tt.expectedErrorString, err.Error())
			}
		})
	}
}

// getAttributes
func GetAttibutes_ValidInput(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput map[string]Attribute
	}{
		{
			name:  "valid header with class input",
			input: "<-hd .class.< testclass anotherclass >",
			expectedOutput: map[string]Attribute{
				"class": {
					TypeAttr: "class",
					ListAttr: []string{"testclass", "anotherclass"},
					Metatype: list,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := getAttributes(tt.input)
			if err != nil {
				t.Errorf("did not expect error but got: %v, in test: %v", err, tt.name)
			}
			if !reflect.DeepEqual(result, tt.expectedOutput) {
				t.Errorf("expected output and output did not match output: %v expected: \n%v\n", result, tt.expectedOutput)
			}
		})
	}
}
