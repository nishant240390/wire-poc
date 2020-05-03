package main

import (
	"testing"
)

func TestInitialise(t *testing.T) {
	tests := []struct {
		name    string
	}{
		{
			name: "name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Initialize()
			if  err != nil{
				t.Errorf("Initialise() error ")
				return
			}
			got.startWar()
		})
	}
}


func TestInitialiseMock(t *testing.T) {
	tests := []struct {
		name    string
	}{
		{
			name: "name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := InitializeMock()
			if  err != nil{
				t.Errorf("Initialise() error ")
				return
			}
			got.startWar()
		})
	}
}