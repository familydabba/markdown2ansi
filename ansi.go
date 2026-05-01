package markdown2ansi

const (
	Reset         = "\x1b[0m"
	Bold          = "\x1b[1m"
	Italic        = "\x1b[3m"
	Underline     = "\x1b[4m"
	Strikethrough = "\x1b[9m"

	BgReset     = "\x1b[49m"
	BgBlack     = "\x1b[40m"
	BgRed       = "\x1b[41m"
	BgGreen     = "\x1b[42m"
	BgYellow    = "\x1b[43m"
	BgBlue      = "\x1b[44m"
	BgMagenta   = "\x1b[45m"
	BgCyan      = "\x1b[46m"
	BgWhite     = "\x1b[47m"
	BgLightGray = "\x1b[48;5;246m"
)

var (
	Black   = "\x1b[30m"
	Red     = "\x1b[31m"
	Green   = "\x1b[32m"
	Yellow  = "\x1b[33m"
	Blue    = "\x1b[34m"
	Magenta = "\x1b[35m"
	Cyan    = "\x1b[36m"
	White   = "\x1b[37m"

	BrightBlack   = "\x1b[90m"
	BrightRed     = "\x1b[91m"
	BrightGreen   = "\x1b[92m"
	BrightYellow  = "\x1b[93m"
	BrightBlue    = "\x1b[94m"
	BrightMagenta = "\x1b[95m"
	BrightCyan    = "\x1b[96m"
	BrightWhite   = "\x1b[97m"
)

type Theme struct {
	Text       string
	Bold       string
	Italic     string
	Code       string
	CodeBlock  string
	Link       string
	Header1    string
	Header2    string
	Header3    string
	LinkBrace  string
	ListBullet string
}

var DefaultTheme = Theme{
	Text:       Reset,
	Bold:       Bold,
	Italic:     Italic,
	Code:       BrightGreen,
	CodeBlock:  Black + BgLightGray,
	Link:       Blue,
	Header1:    Bold + BrightCyan + Underline,
	Header2:    Bold + BrightBlue,
	Header3:    Bold + BrightWhite,
	LinkBrace:  Magenta,
	ListBullet: Green,
}

func (t Theme) ToAnsiCode(name string) string {
	switch name {
	case "reset":
		return Reset
	case "bold":
		return t.Bold
	case "italic":
		return t.Italic
	case "code":
		return t.Code
	case "codeblock":
		return t.CodeBlock
	case "link":
		return t.Link
	case "header1":
		return t.Header1
	case "header2":
		return t.Header2
	case "header3":
		return t.Header3
	case "linkbrace":
		return t.LinkBrace
	case "listbullet":
		return t.ListBullet
	default:
		return t.Text
	}
}
