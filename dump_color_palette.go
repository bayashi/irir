package main

import "strings"

func dumpColorPalette() string {
	colors := []string{}
	for _, c := range orderedColors {
		colors = append(colors, palette[c].Sprintf(c))
	}

	return strings.Join(colors, "\n")
}
