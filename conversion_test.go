package main

import (
	"fmt"
	"testing"
)

func TestToHTML_ValidNode(t *testing.T) {
	tests := []struct {
		name           string
		input          NodeAM
		expectedOutput string
	}{
		{
			name: "valid h2 tag and class attribute",
			input: NodeAM{
				Text: "test text",
				Tag:  "h2d",
				Attributes: map[string]Attribute{
					"class": {
						TypeAttr: "class",
						ListAttr: []string{
							"t", "test",
						},
						Metatype: list,
					},
				},
			},
			expectedOutput: "<h2 class=\"t test\">test text</h2>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := tt.input.toHTML()
			if err != nil {
				t.Fatalf(fmt.Sprintf("did not expect error in test: %v but got: %v", tt.name, err))
			}

			if r != tt.expectedOutput {
				t.Fatalf(fmt.Sprintf("expected: %v but got: %v", tt.expectedOutput, r))
			}
		})
	}
}
