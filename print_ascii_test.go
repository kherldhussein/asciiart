package main

import (
	"testing"
)

func TestPrintArt(t *testing.T) {
	asciMap, err := ReadAscii("standard.txt")
	if err != nil {
		t.Errorf("Error reading ASCII map: %v", err)
		return
	}

	testCases := []struct {
		name        string
		input       string
		expected    string
		expectedErr bool
	}{
		{
			name:        "Print a single word",
			input:       "Hello",
			expected:    " _    _          _   _               \n| |  | |        | | | |              \n| |__| |   ___  | | | |   ___        \n|  __  |  / _ \\ | | | |  / _ \\       \n| |  | | |  __/ | | | | | (_) |      \n|_|  |_|  \\___| |_| |_|  \\___/       \n                                    \n                                    \n",
			expectedErr: false,
		},
		{
			name:        "Print unsupported escape sequence",
			input:       "\\t",
			expected:    "",
			expectedErr: true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			err := PrintArt(tt.input, asciMap)
			if (err != nil) != tt.expectedErr {
				t.Errorf("Go (%q) error = %v, Expected %v", tt.input, err, tt.expectedErr)
				return
			}
		})
	}
}
