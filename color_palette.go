package main

import "github.com/fatih/color"

// See more details:
// * https://github.com/fatih/color
// * https://en.wikipedia.org/wiki/ANSI_escape_code#Colors
var palette = map[string]*color.Color{
	// Bright color text
	"red":        color.New(color.FgHiRed),
	"green":      color.New(color.FgHiGreen),
	"yellow":     color.New(color.FgHiYellow),
	"blue":       color.New(color.FgHiBlue),
	"magenta":    color.New(color.FgHiMagenta),
	"cyan":       color.New(color.FgHiCyan),
	"gray":       color.New(color.FgHiBlack),
	"light_gray": color.New(color.FgWhite),
	"white":      color.New(color.FgHiWhite),

	// Darker color text
	"dark_red":     color.New(color.FgRed),
	"dark_green":   color.New(color.FgGreen),
	"dark_yellow":  color.New(color.FgYellow),
	"dark_blue":    color.New(color.FgBlue),
	"dark_magenta": color.New(color.FgMagenta),
	"dark_cyan":    color.New(color.FgCyan),
	"black":        color.New(color.FgBlack),

	// White text on background color
	"bg_red":     color.New(color.FgHiWhite, color.BgRed),
	"bg_green":   color.New(color.FgHiWhite, color.FgGreen),
	"bg_yellow":  color.New(color.FgHiWhite, color.FgYellow),
	"bg_blue":    color.New(color.FgHiWhite, color.FgBlue),
	"bg_magenta": color.New(color.FgHiWhite, color.FgMagenta),
	"bg_cyan":    color.New(color.FgHiWhite, color.FgCyan),

	// For error text
	"error": color.New(color.FgHiBlue, color.BgRed, color.Underline),
}
