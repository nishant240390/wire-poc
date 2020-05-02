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
			stark,err := NewStarkProvider("john")

			lannister,err := NewLannisterProvider("cersie")

			got, err := Initialize(stark,lannister)
			if  err != nil{
				t.Errorf("Initialise() error ")
				return
			}
			got.startWar()
		})
	}
}