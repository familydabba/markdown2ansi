package markdown2ansi

import (
	"regexp"
	"strings"
)

type TokenType int

const (
	TokenText TokenType = iota
	TokenBold
	TokenItalic
	TokenCode
	TokenCodeBlock
	TokenLink
	TokenHeader1
	TokenHeader2
	TokenHeader3
	TokenHeader4
	TokenHeader5
	TokenHeader6
	TokenListItem
	TokenOrderedListItem
	TokenHorizontalRule
	TokenBlockquote
	TokenLineBreak
)

type Token struct {
	Type    TokenType
	Content string
	URL     string
	Lang    string
	Level   int
}

type Tokens []Token

func (t Tokens) String() string {
	result := ""
	for _, token := range t {
		switch token.Type {
		case TokenText:
			result += token.Content
		case TokenBold:
			result += Bold + token.Content + Reset
		case TokenItalic:
			result += Italic + token.Content + Reset
		case TokenCode:
			result += Cyan + token.Content + Reset
		case TokenCodeBlock:
			result += Magenta + token.Content + Reset
		case TokenLink:
			result += Blue + token.Content + Reset + " (" + Magenta + token.URL + Reset + ")"
		case TokenHeader1:
			result += Bold + Underline + BrightCyan + token.Content + Reset + "\n"
		case TokenHeader2:
			result += Bold + BrightBlue + token.Content + Reset + "\n"
		case TokenHeader3:
			result += Bold + token.Content + Reset + "\n"
		case TokenHeader4, TokenHeader5, TokenHeader6:
			result += Bold + token.Content + Reset + "\n"
		case TokenListItem:
			result += Green + "• " + Reset + token.Content + "\n"
		case TokenOrderedListItem:
			result += Green + token.Content + Reset + "\n"
		case TokenHorizontalRule:
			result += Cyan + strings.Repeat("─", 60) + Reset + "\n"
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

func Parse(markdown string) Tokens {
	lines := strings.Split(markdown, "\n")
	var tokens Tokens

	reBold = regexp.MustCompile(`\*\*([^*]+)\*\*`)
	reItalicStar = regexp.MustCompile(`\*([^*]+)\*`)
	reItalicUS = regexp.MustCompile(`_([^_]+)_`)
	reCode = regexp.MustCompile("`([^`]+)`")
	reLink = regexp.MustCompile(`\[([^\]]+)\]\(([^)]+)\)`)

	inCodeBlock := false
	codeLang := ""

	for i := 0; i < len(lines); i++ {
		line := lines[i]

		if strings.HasPrefix(line, "```") && !inCodeBlock {
			inCodeBlock = true
			matches := regexp.MustCompile("```(\\w*)").FindStringSubmatch(line)
			if len(matches) > 1 {
				codeLang = matches[1]
			}
			tokens = append(tokens, Token{Type: TokenCodeBlock, Lang: codeLang, Content: ""})
			continue
		}

		if inCodeBlock {
			if strings.HasPrefix(line, "```") {
				inCodeBlock = false
				codeLang = ""
			} else {
				tokens = append(tokens, Token{Type: TokenCodeBlock, Content: line + "\n"})
			}
			continue
		}

		if strings.HasPrefix(line, "# ") {
			line = strings.TrimPrefix(line, "# ")
			line = processInline(line, reBold, reItalicStar, reItalicUS, reCode, reLink)
			tokens = append(tokens, Token{Type: TokenHeader1, Content: line})
			continue
		}
		if strings.HasPrefix(line, "## ") {
			line = strings.TrimPrefix(line, "## ")
			line = processInline(line, reBold, reItalicStar, reItalicUS, reCode, reLink)
			tokens = append(tokens, Token{Type: TokenHeader2, Content: line})
			continue
		}
		if strings.HasPrefix(line, "### ") {
			line = strings.TrimPrefix(line, "### ")
			line = processInline(line, reBold, reItalicStar, reItalicUS, reCode, reLink)
			tokens = append(tokens, Token{Type: TokenHeader3, Content: line})
			continue
		}
		if strings.HasPrefix(line, "#### ") {
			line = strings.TrimPrefix(line, "#### ")
			line = processInline(line, reBold, reItalicStar, reItalicUS, reCode, reLink)
			tokens = append(tokens, Token{Type: TokenHeader4, Content: line})
			continue
		}
		if strings.HasPrefix(line, "##### ") {
			line = strings.TrimPrefix(line, "##### ")
			line = processInline(line, reBold, reItalicStar, reItalicUS, reCode, reLink)
			tokens = append(tokens, Token{Type: TokenHeader5, Content: line})
			continue
		}
		if strings.HasPrefix(line, "###### ") {
			line = strings.TrimPrefix(line, "###### ")
			line = processInline(line, reBold, reItalicStar, reItalicUS, reCode, reLink)
			tokens = append(tokens, Token{Type: TokenHeader6, Content: line})
			continue
		}

		if strings.HasPrefix(line, "> ") {
			line = strings.TrimPrefix(line, "> ")
			line = processInline(line, reBold, reItalicStar, reItalicUS, reCode, reLink)
			tokens = append(tokens, Token{Type: TokenBlockquote, Content: line})
			continue
		}

		if strings.HasPrefix(line, "- ") || strings.HasPrefix(line, "* ") || strings.HasPrefix(line, "+ ") {
			line = strings.TrimPrefix(strings.TrimPrefix(strings.TrimPrefix(line, "- "), "* "), "+ ")
			line = processInline(line, reBold, reItalicStar, reItalicUS, reCode, reLink)
			tokens = append(tokens, Token{Type: TokenListItem, Content: line})
			continue
		}

		orderedListRe := regexp.MustCompile(`^(\d+)\.\s+(.+)$`)
		if matches := orderedListRe.FindStringSubmatch(line); len(matches) > 0 {
			content := matches[1] + ". " + matches[2]
			content = processInline(content, reBold, reItalicStar, reItalicUS, reCode, reLink)
			tokens = append(tokens, Token{Type: TokenOrderedListItem, Content: content})
			continue
		}

		hrRe := regexp.MustCompile(`^[-*_]{3,}$`)
		if hrRe.MatchString(strings.TrimSpace(line)) {
			tokens = append(tokens, Token{Type: TokenHorizontalRule})
			continue
		}

		if line == "" {
			tokens = append(tokens, Token{Type: TokenLineBreak})
			continue
		}

		line = processInline(line, reBold, reItalicStar, reItalicUS, reCode, reLink)
		tokens = append(tokens, Token{Type: TokenText, Content: line + "\n"})
	}

	return tokens
}

func processInline(line string, reBold, reItalicStar, reItalicUS, reCode *regexp.Regexp, reLink *regexp.Regexp) string {
	line = reLink.ReplaceAllString(line, "$1 ($2)")

	line = reCode.ReplaceAllStringFunc(line, func(s string) string {
		matches := reCode.FindStringSubmatch(s)
		if len(matches) > 1 {
			return BrightGreen + matches[1] + Reset
		}
		return s
	})

	line = reBold.ReplaceAllStringFunc(line, func(s string) string {
		matches := reBold.FindStringSubmatch(s)
		if len(matches) > 1 {
			return Bold + matches[1] + Reset
		}
		return s
	})

	line = reItalicStar.ReplaceAllStringFunc(line, func(s string) string {
		matches := reItalicStar.FindStringSubmatch(s)
		if len(matches) > 1 {
			return Italic + matches[1] + Reset
		}
		return s
	})

	line = reItalicUS.ReplaceAllStringFunc(line, func(s string) string {
		matches := reItalicUS.FindStringSubmatch(s)
		if len(matches) > 1 {
			return Italic + matches[1] + Reset
		}
		return s
	})

	return line
}
