package main

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

func dumpRule(rule string) string {
	r, err := loadRule(cfgFilePath, rule)
	if err != nil {
		return fmt.Sprintf("could not find %s, %s", rule, err.Error())
	}

	y, err := yaml.Marshal(r)
	if err != nil {
		return fmt.Sprintf("could not get YAML %s, %s", rule, err.Error())
	}

	return string(y)
}

func dumpRules() string {
	cfg, err := allCfg(cfgFilePath)
	if err != nil {
		return fmt.Sprintf("could not get config %s", err.Error())
	}

	var rules []string
	for k := range cfg {
		rules = append(rules, k)
	}

	return strings.Join(rules, "\n")
}
