package markdown2ansi

import (
	"strings"
	"testing"
)

func TestRenderHeader1(t *testing.T) {
	input := "# Header"
	result := Render(input)

	if !strings.Contains(result, Bold) {
		t.Error("Render header should contain Bold")
	}
}

func TestRenderHeader2(t *testing.T) {
	input := "## Header"
	result := Render(input)

	if !strings.Contains(result, Bold) {
		t.Error("Render header should contain Bold")
	}
}

func TestRenderHorizontalRule(t *testing.T) {
	input := "---"
	result := Render(input)

	if !strings.Contains(result, "─") {
		t.Error("Render HR should contain horizontal line character")
	}
}

func TestRenderCodeBlock(t *testing.T) {
	input := "```go\nfmt.Println(\"hello\")\n```"
	result := Render(input)

	if result == "" {
		t.Error("Render code block should not be empty")
	}
}

func TestRenderMultipleElements(t *testing.T) {
	input := `# Title

This is **bold** and *italic* text.

- Item 1
- Item 2

> A quote

---

[link](http://example.com)
`
	result := Render(input)

	if result == "" {
		t.Error("Render multiple elements should not be empty")
	}
	if !strings.Contains(result, Bold) {
		t.Error("Render should contain bold")
	}
	if !strings.Contains(result, "•") {
		t.Error("Render should contain list bullet")
	}
}

func TestRenderer_SetWidth(t *testing.T) {
	r := NewRenderer()
	r.SetWidth(40)
	if r.width != 40 {
		t.Errorf("SetWidth() = %d, want 40", r.width)
	}
}

func TestRenderer_SetTheme(t *testing.T) {
	r := NewRenderer()
	custom := Theme{
		Text:       White,
		Bold:       Bold,
		Italic:     Italic,
		Code:       Green,
		CodeBlock:  Magenta,
		Link:       Blue,
		Header1:    Red,
		Header2:    Red,
		Header3:    Red,
		LinkBrace:  Magenta,
		ListBullet: Green,
	}
	r.SetTheme(custom)

	if r.theme.Text != White {
		t.Errorf("SetTheme() Text = %q, want %q", r.theme.Text, White)
	}
}
