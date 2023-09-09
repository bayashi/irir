package main

import (
	"testing"

	a "github.com/bayashi/actually"
)

const RULE = "irir_test_rule"
const DEFAULT_RULE = RULE + "_default"

func TestGetDefaultRule(t *testing.T) {
	t.Setenv(ENV_KEY_IRIR_DEFAULT_RULE, RULE)

	a.Got(getDefaultRule()).Expect(RULE).Same(t)
}

func TestTargetRule(t *testing.T) {
	for name, tt := range map[string]struct {
		args   []string
		expect string
	}{
		"no arg, then set default rule": {
			args: []string{},
			expect: DEFAULT_RULE,
		},
		"blank arg, then set default rule": {
			args: []string{""},
			expect: DEFAULT_RULE,
		},
		"just one normal rule": {
			args: []string{RULE},
			expect: RULE,
		},
		"just one normal rule with wrap command": {
			args: []string{RULE, "--", "ls", "-lah"},
			expect: RULE,
		},
		"no rule with wrap command, then set default rule": {
			args: []string{"--", "ls", "-lah"},
			expect: DEFAULT_RULE,
		},
		"blank rule with wrap command, then set default rule": {
			args: []string{"", "--", "ls", "-lah"},
			expect: DEFAULT_RULE,
		},
	}{
		t.Setenv(ENV_KEY_IRIR_DEFAULT_RULE, DEFAULT_RULE)
		t.Run(name, func(t *testing.T) {
			o := &options{}
			o.targetRule(tt.args)
			a.Got(o.rule).Expect(tt.expect).Same(t)
		})
	}
}
