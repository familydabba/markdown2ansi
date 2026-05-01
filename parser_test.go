package markdown2ansi

import (
	"strings"
	"testing"
)

func TestParseHeaders(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantType TokenType
	}{
		{"header1", "# Header 1", TokenHeader1},
		{"header2", "## Header 2", TokenHeader2},
		{"header3", "### Header 3", TokenHeader3},
		{"header4", "#### Header 4", TokenHeader4},
		{"header5", "##### Header 5", TokenHeader5},
		{"header6", "###### Header 6", TokenHeader6},
		{"list", "- list item", TokenListItem},
		{"ordered_list", "1. ordered item", TokenOrderedListItem},
		{"blockquote", "> quote", TokenBlockquote},
		{"hr", "---", TokenHorizontalRule},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokens := Parse(tt.input)
			if len(tokens) == 0 {
				t.Errorf("Parse(%q) returned no tokens", tt.input)
				return
			}
			if tokens[0].Type != tt.wantType {
				t.Errorf("Parse(%q)[0].Type = %v, want %v", tt.input, tokens[0].Type, tt.wantType)
			}
		})
	}
}

func TestParseMixedContent(t *testing.T) {
	input := `# Header

This is a paragraph.

- Item 1
- Item 2

> A quote

---

Some text with **bold** and *italic*.
`
	tokens := Parse(input)

	if len(tokens) < 5 {
		t.Errorf("Parse() returned %d tokens, want at least 5", len(tokens))
	}
}

func TestParsePreservesContent(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantCont string
	}{
		{"header1", "# My Header", "My Header"},
		{"header2", "## Another", "Another"},
		{"list", "- item", "item"},
		{"ordered", "1. first", "first"},
		{"blockquote", "> quote text", "quote text"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokens := Parse(tt.input)
			if len(tokens) == 0 {
				t.Fatalf("Parse(%q) returned no tokens", tt.input)
			}
			if tokens[0].Content != tt.wantCont && !strings.Contains(tokens[0].Content, tt.wantCont) {
				t.Errorf("Parse(%q).Content = %q, want to contain %q", tt.input, tokens[0].Content, tt.wantCont)
			}
		})
	}
}

func TestTokensString(t *testing.T) {
	tokens := Tokens{
		{Type: TokenText, Content: "Hello "},
		{Type: TokenBold, Content: "World"},
		{Type: TokenLineBreak},
	}

	result := tokens.String()
	if result == "" {
		t.Error("Tokens.String() should not be empty")
	}
	if len(result) < len("Hello World") {
		t.Error("Tokens.String() should contain token content")
	}
}
