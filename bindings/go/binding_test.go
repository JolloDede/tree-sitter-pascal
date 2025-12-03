package tree_sitter_pascal_test

import (
	"testing"

	tree_sitter_pascal "github.com/jollodede/tree-sitter-pascal/bindings/go"
	tree_sitter "github.com/tree-sitter/go-tree-sitter"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_pascal.Language())
	if language == nil {
		t.Errorf("Error loading Pascal grammar")
	}
}
