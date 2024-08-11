package main

import (
	"fmt"
	"testing"
)

// parseAM
func TestParseAM_ValidString(t *testing.T) {
	type Test struct {
		input    string
		expected NodeAM
	}

	tests := []Test{
		{
			input: " this is a header <-hd-",
			expected: NodeAM{
				Text: "this is a header",
				Tag:  "hd",
			},
		},
		{
			input: "-hd->this is a header ",
			expected: NodeAM{
				Text: "this is a header",
				Tag:  "hd",
			},
		},
	}

	for _, test := range tests {
		p := NewAmParser()
		node, err := p.parseAM(test.input)
		if err != nil {
			t.Errorf(fmt.Sprintf("function did not expect an error but got: %v", err))
			continue
		}
		if node.Tag != test.expected.Tag {
			t.Errorf(fmt.Sprintf("expected tag did not match expected: %v, got: %v", test.expected.Tag, node.Tag))
			continue
		}
		if node.Text != test.expected.Text {
			t.Errorf(fmt.Sprintf("expected text did not match expected: %v, got: %v", test.expected.Text, node.Text))
			continue
		}

	}

}
