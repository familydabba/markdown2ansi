package markdown2ansi

import (
	"regexp"
	"strings"
)

type Renderer struct {
	width int
	theme Theme
}

func NewRenderer() *Renderer {
	return &Renderer{
		width: 80,
		theme: DefaultTheme,
	}
}

func (r *Renderer) SetWidth(w int) {
	r.width = w
}

func (r *Renderer) SetTheme(t Theme) {
	r.theme = t
}

func (r *Renderer) Render(markdown string) string {
	tokens := Parse(markdown)
	return r.renderTokens(tokens)
}

func (r *Renderer) renderTokens(tokens Tokens) string {
	var result string
	for _, token := range tokens {
		switch token.Type {
		case TokenText:
			wrapped := r.wrapText(token.Content, r.width)
			result += r.theme.Text + wrapped
		case TokenBold:
			result += r.theme.Bold + token.Content + Reset
		case TokenItalic:
			result += r.theme.Italic + token.Content + Reset
		case TokenCode:
			result += r.theme.Code + token.Content + Reset
		case TokenCodeBlock:
			result += r.theme.CodeBlock + token.Content + Reset
		case TokenLink:
			result += r.theme.Link + token.Content + Reset + " (" + r.theme.LinkBrace + token.URL + Reset + ")"
		case TokenHeader1:
			result += r.theme.Header1 + token.Content + Reset + "\n"
		case TokenHeader2:
			result += r.theme.Header2 + token.Content + Reset + "\n"
		case TokenHeader3:
			result += r.theme.Header3 + token.Content + Reset + "\n"
		case TokenHeader4, TokenHeader5, TokenHeader6:
			result += r.theme.Bold + token.Content + Reset + "\n"
		case TokenListItem:
			result += r.theme.ListBullet + "• " + Reset + token.Content + "\n"
		case TokenOrderedListItem:
			result += r.theme.ListBullet + token.Content + Reset + "\n"
		case TokenHorizontalRule:
			result += Cyan + strings.Repeat("─", r.width) + Reset + "\n"
		case TokenBlockquote:
			result += Yellow + "│ " + Reset + token.Content + "\n"
		case TokenLineBreak:
			result += "\n"
		default:
			result += token.Content
		}
	}
	return result
}

func (r *Renderer) wrapText(text string, width int) string {
	lines := strings.Split(text, "\n")
	var result string
	for _, line := range lines {
		if len(line) <= width {
			result += line + "\n"
			continue
		}
		result += r.wrapLine(line, width) + "\n"
	}
	return result
}

func (r *Renderer) wrapLine(text string, width int) string {
	var result string
	for len(text) > width {
		cut := width
		for cut > 0 && text[cut] != ' ' {
			cut--
		}
		if cut == 0 {
			cut = width
		}
		result += text[:cut] + "\n"
		text = text[cut:]
		if len(text) > 0 && text[0] == ' ' {
			text = text[1:]
		}
	}
	result += text
	return result
}

func Render(markdown string) string {
	r := NewRenderer()
	return r.Render(markdown)
}

var (
	reBold       = regexp.MustCompile(`\*\*([^*]+)\*\*`)
	reItalicStar = regexp.MustCompile(`\*([^*]+)\*`)
	reItalicUS   = regexp.MustCompile(`_([^_]+)_`)
	reCode       = regexp.MustCompile("`([^`]+)`")
	reLink       = regexp.MustCompile(`\[([^\]]+)\]\(([^)]+)\)`)
	reCodeBlock  = regexp.MustCompile("(?s)```(\\w*)\\n(.+?)```")
)

func ProcessInline(text string) string {
	text = reLink.ReplaceAllString(text, "$1 ($2)")

	text = reCode.ReplaceAllStringFunc(text, func(s string) string {
		matches := reCode.FindStringSubmatch(s)
		if len(matches) > 1 {
			return BrightGreen + matches[1] + Reset
		}
		return s
	})

	text = reBold.ReplaceAllStringFunc(text, func(s string) string {
		matches := reBold.FindStringSubmatch(s)
		if len(matches) > 1 {
			return Bold + matches[1] + Reset
		}
		return s
	})

	text = reItalicStar.ReplaceAllStringFunc(text, func(s string) string {
		matches := reItalicStar.FindStringSubmatch(s)
		if len(matches) > 1 {
			return Italic + matches[1] + Reset
		}
		return s
	})

	text = reItalicUS.ReplaceAllStringFunc(text, func(s string) string {
		matches := reItalicUS.FindStringSubmatch(s)
		if len(matches) > 1 {
			return Italic + matches[1] + Reset
		}
		return s
	})

	return text
}
