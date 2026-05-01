package markdown2ansi

import (
	"strings"
	"testing"
)

func TestThemeToAnsiCode(t *testing.T) {
	theme := DefaultTheme

	tests := []struct {
		name    string
		input   string
		wantStr string
	}{
		{"bold", "bold", Bold},
		{"italic", "italic", Italic},
		{"code", "code", BrightGreen},
		{"codeblock", "codeblock", Black + BgLightGray},
		{"link", "link", Blue},
		{"header1", "header1", Bold + BrightCyan + Underline},
		{"header2", "header2", Bold + BrightBlue},
		{"header3", "header3", Bold + BrightWhite},
		{"linkbrace", "linkbrace", Magenta},
		{"listbullet", "listbullet", Green},
		{"unknown", "unknown", Reset},
		{"reset", "reset", Reset},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := theme.ToAnsiCode(tt.input)
			if result != tt.wantStr {
				t.Errorf("Theme.ToAnsiCode(%q) = %q, want %q", tt.input, result, tt.wantStr)
			}
		})
	}
}

func TestANSIConstants(t *testing.T) {
	if Reset == "" {
		t.Error("Reset should not be empty")
	}
	if Bold == "" {
		t.Error("Bold should not be empty")
	}
	if Italic == "" {
		t.Error("Italic should not be empty")
	}
	if Cyan == "" {
		t.Error("Cyan should not be empty")
	}
	if !strings.Contains(Cyan, "\x1b[") {
		t.Error("Cyan should contain ANSI escape")
	}
}

func TestColorsExist(t *testing.T) {
	colors := []string{
		Black, Red, Green, Yellow, Blue, Magenta, Cyan, White,
		BrightBlack, BrightRed, BrightGreen, BrightYellow, BrightBlue, BrightMagenta, BrightCyan, BrightWhite,
	}

	for _, color := range colors {
		if color == "" {
			t.Error("Color should not be empty")
		}
		if !strings.Contains(color, "\x1b[") {
			t.Errorf("Color %q should contain ANSI escape", color)
		}
	}
}
