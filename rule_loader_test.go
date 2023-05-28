package main

import (
	"os"
	"path/filepath"
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

func createCfgFile(t *testing.T, content string) string {
	temp, _ := filepath.Abs(t.TempDir())
	path := filepath.Join(temp, irirConfigFiles[0])
	f, _ := os.Create(path)
	defer f.Close()
	_, err := f.Write([]byte(content))
	if err != nil {
		panic("could not write: " + path + ", " + content)
	}

	return path
}

func TestAllCfgMatch(t *testing.T) {
	path := createCfgFile(t, "foo:\n- type: match\n  match: bar\n  color: red\n  target: word\n")
	cfgFilePath := func(f string) string {
		return path
	}
	cfg, err := loadRule(cfgFilePath, "foo")
	a.Got(err).NoError(t)
	a.Got(len(cfg)).Expect(1).Same(t)
	a.Got(cfg[0].Match).Expect("bar").Same(t)

	_, err2 := loadRule(cfgFilePath, "no")
	a.Got(err2.Error()).Expect("'no' doesn't exists in config")
}

func TestAllCfgRegexp(t *testing.T) {
	path := createCfgFile(t, "foo:\n- type: regexp\n  match: Ba.\n  color: green\n  target: word\n")
	cfgFilePath := func(f string) string {
		return path
	}
	cfg, err := loadRule(cfgFilePath, "foo")
	a.Got(err).NoError(t)
	a.Got(len(cfg)).Expect(1).Same(t)
	a.Got(cfg[0].regexp.String()).Expect("(Ba.)").Same(t)
}
