package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

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
	Color  string         `json:"color" jsonschema:"enum=red,enum=green,enum=yellow,enum=blue,enum=magenta,enum=cyan,enum=gray,enum=light_gray,enum=white,enum=black,enum=dark_red,enum=dark_green,enum=dark_yellow,enum=dark_blue,enum=dark_magenta,enum=dark_cyan,enum=bg_red,enum=bg_green,enum=bg_yellow,enum=bg_blue,enum=bg_magenta,enum=bg_cyan,enum=error"`
	Target string         `json:"target" jsonschema:"enum=word,enum=line"`
	regexp *regexp.Regexp // If Type would be "regexp", then the compiled regexp is set here
}

func loadRule(cfp func(f string) string, ruleName string) ([]*Rule, error) {
	cfg, err := allCfg(cfp)
	if err != nil {
		return nil, err
	}

	r, isExists := cfg[ruleName]
	if !isExists {
		return nil, fmt.Errorf("'%s' doesn't exists in config", ruleName)
	}

	for i, rr := range r {
		if rr.Type == TYPE_REGEXP {
			r[i].regexp = regexp.MustCompile("(" + rr.Match + ")")
		}
		if _, ok := palette[strings.ToLower(rr.Color)]; !ok {
			r[i].Color = "error"
		} else {
			r[i].Color = strings.ToLower(rr.Color)
		}
	}

	return r, nil
}

func allCfg(cfp func(f string) string) (map[string][]*Rule, error) {
	bytes, err := loadCfg(cfp)
	if err != nil {
		return nil, err
	}

	cfg, err2 := parseCfg(bytes)
	if err2 != nil {
		return nil, err2
	}

	return cfg, nil
}

func parseCfg(bytes []byte) (map[string][]*Rule, error) {
	var result map[string][]*Rule
	err := yaml.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func loadCfg(cfp func(f string) string) ([]byte, error) {
	expectedCfgFiles := []string{}
	for _, cfgFile := range irirConfigFiles {
		cfgFullPath := cfp(cfgFile)
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
