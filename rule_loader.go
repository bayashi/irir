package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/adrg/xdg"
	yaml "gopkg.in/yaml.v3"
)

const irir = "irir"

var (
	irirDir         = irir
	irirFile        = irir + "_rule"
	irirConfigFiles = [2]string{irirFile + ".yaml", irirFile + ".yml"}
)

const (
	TARGET_WORD = "word"
	TARGET_LINE = "line"

	TYPE_MATCH  = "match"
	TYPE_PREFIX = "prefix"
	TYPE_SUFFIX = "suffix"
	TYPE_REGEXP = "regexp"
)

type Rule struct {
	Type   string         `json:"type" jsonschema:"enum=match,enum=prefix,enum=suffix,enum=regexp"`
	Match  string         `json:"match" jsonschema:"string"`
	Color  string         `json:"color" jsonschema:"string"`
	Target string         `json:"target" jsonschema:"enum=word,enum=line"`
	regexp *regexp.Regexp // If Type would be "regexp", then the compiled regexp is set here
}

func loadRule(ruleName string) ([]*Rule, error) {
	bytes, err := loadCfg()
	if err != nil {
		return []*Rule{}, err
	}

	cfg, err2 := parseCfg(bytes)
	if err2 != nil {
		return []*Rule{}, err
	}

	r, isExists := cfg[ruleName]
	if !isExists {
		return []*Rule{}, fmt.Errorf("'%s' doesn't exists in config", ruleName)
	}

	for i, rr := range r {
		if rr.Type == TYPE_REGEXP {
			m, _ := rr.re()
			r[i].regexp = regexp.MustCompile(m)
		}
		if _, ok := palette[strings.ToLower(rr.Color)]; !ok {
			r[i].Color = "error"
		} else {
			r[i].Color = strings.ToLower(rr.Color)
		}
	}

	return r, nil
}

func parseCfg(bytes []byte) (map[string][]*Rule, error) {
	var result map[string][]*Rule
	err := yaml.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func loadCfg() ([]byte, error) {
	expectedCfgFiles := []string{}
	for _, cfgFile := range irirConfigFiles {
		cfgFullPath := fullPath(cfgFile)
		expectedCfgFiles = append(expectedCfgFiles, cfgFullPath)
		if _, err := os.Stat(cfgFullPath); err != nil {
			continue
		}
		bytes, err := os.ReadFile(cfgFullPath)
		if err != nil {
			continue
		}

		return bytes, nil
	}

	return nil, fmt.Errorf("could not read rule config file %s", strings.Join(expectedCfgFiles, ", "))
}

func fullPath(fileName string) string {
	return filepath.Join(xdg.ConfigHome, irirDir, fileName)
}

func (r *Rule) re() (string, string) {
	m := strings.Split(r.Match, "\n")
	if len(m) == 2 {
		return m[0], m[1]
	}

	return "(" + r.Match + ")", "$1"
}
