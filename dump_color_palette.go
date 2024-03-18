package main

import (
	"strings"

	"github.com/bayashi/colorpalette"
)

func dumpColorPalette() string {
	colors := []string{}
	for _, c := range colorpalette.List() {
		colors = append(colors, colorpalette.Get(c).Sprintf(c))
	}

	return strings.Join(colors, "\n")
}
