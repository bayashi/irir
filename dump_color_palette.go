package main

import "strings"


func dumpColorPalette() string {
	colors := []string{}
	for k := range palette {
		colors = append(colors, "enum=" + k)
	}

	return strings.Join(colors, ",")
}