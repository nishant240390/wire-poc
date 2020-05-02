package main

import (
	"testing"
)

func TestInitialise(t *testing.T) {
	tests := []struct {
		name    string
	}{
		{
			name : "test1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Initialise()
			if err != nil {
				t.Errorf("Initialise() error ")
				return
			}
		})
	}
}

