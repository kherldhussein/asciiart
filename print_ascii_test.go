package main

import "testing"

func TestPrintArt(t *testing.T) {
	type args struct {
		str     string
		asciMap [][]string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintArt(tt.args.str, tt.args.asciMap)
		})
	}
}
