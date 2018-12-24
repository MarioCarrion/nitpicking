package nitpicking_test

import (
	"path/filepath"
	"testing"

	"github.com/MarioCarrion/nitpicking"
)

func TestNitpicker_Validate(t *testing.T) {
	tests := [...]struct {
		name          string
		filename      string
		expectedError bool
	}{
		{
			"OK",
			"nitpicker_valid.go",
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(ts *testing.T) {
			n := nitpicking.Nitpicker{}

			if err := n.Validate(filepath.Join("testdata", tt.filename)); tt.expectedError != (err != nil) {

				ts.Fatalf("expected error %t, got %s", tt.expectedError, err)
			}
		})
	}
}
