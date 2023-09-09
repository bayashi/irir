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

func TestSetWrapCommand(t *testing.T) {
	for name, tt := range map[string]struct {
		args  []string
		expectCmdName string
		expectCmdArgs []string
	}{
		"no wrap command": {
			args: []string{},
			expectCmdName: "",
			expectCmdArgs: nil,
		},
		"no wrap command, just only rule": {
			args: []string{"rule"},
			expectCmdName: "",
			expectCmdArgs: nil,
		},
		"wrap command with rule": {
			args: []string{"rule", "--", "ls", "-la"},
			expectCmdName: "ls",
			expectCmdArgs: []string{"-la"},
		},
		"wrap command without rule": {
			args: []string{"--", "ls", "-la"},
			expectCmdName: "ls",
			expectCmdArgs: []string{"-la"},
		},
	}{
		t.Run(name, func(t *testing.T) {
			o := &options{}
			o.setWrapCommand(tt.args)
			a.Got(o.wrapCmdName).Expect(tt.expectCmdName).Same(t)
			a.Got(o.wrapCmdArgs).Expect(tt.expectCmdArgs).Same(t)
		})
	}
}
