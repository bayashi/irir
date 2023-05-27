package main

import (
	"testing"

	a "github.com/bayashi/actually"
)

func TestPalette(t *testing.T) {
	a.Got(len(palette)).Expect(len(orderedColors)).Same(t)
}
