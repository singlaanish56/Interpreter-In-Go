package ast

import (
	"github.com/singlaanish56/Interpreter-In-Go/token"
	"testing"
)

func TestString(t *testing.T) {

	rootNode := &ASTRootNode{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Identifier:"let"},
				Variable: &Variable{Token: token.Token{Type: token.VARIABLE, Identifier: "ab"}, Value: "ab"},
				Value: &Variable{Token: token.Token{Type: token.VARIABLE, Identifier: "another one"}, Value: "another one"},
			},
		},
	}

	if rootNode.String() != "let ab = another one;"{
		t.Errorf("rootNode.String() errored out %q", rootNode.String())
	}
}