package colors

import "runtime"

const (
	FgBlack   = "30"
	FgRed     = "31"
	FgGreen   = "32"
	FgYellow  = "33"
	FgBlue    = "34"
	FgMagenta = "35"
	FgCyan    = "36"
	FgWhite   = "37"
	BgBlack   = "40"
	BgRed     = "41"
	BgGreen   = "42"
	BgYellow  = "43"
	BgBlue    = "44"
	BgMagenta = "45"
	BgCyan    = "46"
	BgWhite   = "47"
	Default   = "39"
	Reset     = "\033[0m"
)

func Colorize(text string, fgColor string, params ...string) string {
	if runtime.GOOS == "windows" {
		return text
	}

	var color string

	if len(params) > 0 {
		color = "\033[" + fgColor + ";" + params[0] + "m"
	} else {
		color = "\033[" + fgColor + "m"
	}

	return color + text + Reset
}
