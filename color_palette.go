package main

import c "github.com/fatih/color"

// See more details:
// * https://github.com/fatih/color
// * https://en.wikipedia.org/wiki/ANSI_escape_code#Colors
var palette = map[string]*c.Color{
	// Bright color text
	"red":        c.New(c.FgHiRed),
	"green":      c.New(c.FgHiGreen),
	"yellow":     c.New(c.FgHiYellow),
	"blue":       c.New(c.FgHiBlue),
	"magenta":    c.New(c.FgHiMagenta),
	"cyan":       c.New(c.FgHiCyan),
	"gray":       c.New(c.FgHiBlack),
	"light_gray": c.New(c.FgWhite),
	"white":      c.New(c.FgHiWhite),

	// Darker color text
	"black":        c.New(c.FgBlack),
	"dark_red":     c.New(c.FgRed),
	"dark_green":   c.New(c.FgGreen),
	"dark_yellow":  c.New(c.FgYellow),
	"dark_blue":    c.New(c.FgBlue),
	"dark_magenta": c.New(c.FgMagenta),
	"dark_cyan":    c.New(c.FgCyan),

	// White text on background color
	"bg_red":     c.New(c.FgHiWhite, c.BgRed),
	"bg_green":   c.New(c.FgHiWhite, c.BgGreen),
	"bg_yellow":  c.New(c.FgHiWhite, c.BgYellow),
	"bg_blue":    c.New(c.FgHiWhite, c.BgBlue),
	"bg_magenta": c.New(c.FgHiWhite, c.BgMagenta),
	"bg_cyan":    c.New(c.FgHiWhite, c.BgCyan),

	// For error text
	"error": c.New(c.FgHiBlue, c.BgRed, c.Underline),
}

var orderedColors = []string{
	"red",
	"green",
	"yellow",
	"blue",
	"magenta",
	"cyan",
	"gray",
	"light_gray",
	"white",

	"black",
	"dark_red",
	"dark_green",
	"dark_yellow",
	"dark_blue",
	"dark_magenta",
	"dark_cyan",

	"bg_red",
	"bg_green",
	"bg_yellow",
	"bg_blue",
	"bg_magenta",
	"bg_cyan",

	"error",
}
