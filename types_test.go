package nit_test

import (
	"go/ast"
	"go/token"
	"testing"

	"github.com/MarioCarrion/nit"
)

//nolint:dupl
func TestTypesValidator_Validate(t *testing.T) {
	tests := [...]struct {
		name          string
		filename      string
		expectedError bool
	}{
		{
			"OK",
			"types_valid.go",
			false,
		},
		{
			"OK: group",
			"types_group3.go",
			false,
		},
		{
			"Error: parenthesized declaration",
			"types_paren.go",
			true,
		},
		{
			"Error: group 1",
			"types_group1.go",
			true,
		},
		{
			"Error: group 2",
			"types_group2.go",
			true,
		},
		{
			"Error: sorted",
			"types_sorted.go",
			true,
		},
		{
			"Error: group 4",
			"types_group4.go",
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(ts *testing.T) {
			f, fset := newParserFile(ts, tt.filename)

			comments := nit.NewBreakComments(fset, f.Comments)
			validator := nit.NewTypesValidator(comments)

			for _, s := range f.Decls {
				switch g := s.(type) {
				case *ast.GenDecl:
					if g.Tok == token.TYPE {
						if err := validator.Validate(g, fset); tt.expectedError != (err != nil) {
							ts.Fatalf("expected error %t, got %s", tt.expectedError, err)
						}
						break
					}
				}
			}
		})
	}
}
