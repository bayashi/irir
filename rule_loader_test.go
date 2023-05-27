package main

import (
	"reflect"
	"strings"
	"testing"

	a "github.com/bayashi/actually"
)

// Verify that an enum list for jsonschema on `Color` of `Rule` struct is same as palette
func TestRuleColorEnum(t *testing.T) {
	colors := []string{}
	for _, c := range orderedColors {
		colors = append(colors, "enum=" + c)
	}
	expect := strings.Join(colors, ",")

	ty := reflect.TypeOf(Rule{})
	tag, ok := ty.FieldByName("Color")
	a.Got(ok).True(t)
	a.Got(tag.Tag.Get("jsonschema")).Expect(expect).Same(t)
}
