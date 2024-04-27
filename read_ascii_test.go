package main

import (
	"reflect"
	"testing"
)

func TestReadASCIIMapFromFile(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    [][]string
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadASCIIMapFromFile(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadASCIIMapFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadASCIIMapFromFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
