package main

import (
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
