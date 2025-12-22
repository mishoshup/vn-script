package tree_sitter_vn_script_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_vn_script "github.com/mishoshup/vn-script.git/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_vn_script.Language())
	if language == nil {
		t.Errorf("Error loading VN Script grammar")
	}
}
