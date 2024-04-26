package main

import "testing"

func TestPrintArt(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		asciMap   [][]string
		expectErr bool
	}{
		{
			name:      "Empty string",
			input:     "",
			asciMap:   nil,
			expectErr: false,
		},
		{
			name:      "Newline escape sequence",
			input:     "\\n",
			asciMap:   nil,
			expectErr: false,
		},
		{
			name:      "Unsupported escape sequence",
			input:     "\\t",
			asciMap:   nil,
			expectErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := PrintArt(tt.input, tt.asciMap)
			if tt.expectErr && err == nil {
				t.Errorf("Expected an error but got none")
			}
			if !tt.expectErr && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}
