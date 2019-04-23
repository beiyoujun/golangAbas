package deferDo

import (
	"testing"
)

func Test_deferdo(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"xsj"}, {"徐士杰"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deferdo()
		})
	}
}

func Test_tr(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr()
		})
	}
}
