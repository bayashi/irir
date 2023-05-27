package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

const (
	ENV_EDITOR     = "EDITOR"
	DEFAULT_EDITOR = "vi"
)

func editConfig(file string) string {
	if runtime.GOOS == "windows" {
		return "Not supported this option on Windows"
	}

	editor := editor()
	c := exec.Command(editor, file)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	err := c.Run()
	if err != nil {
		return fmt.Sprintf("could not open editor %s, %s", editor, err.Error())
	}

	return ""
}

func editor() string {
	editor := os.Getenv(ENV_EDITOR)
	if editor == "" {
		editor = DEFAULT_EDITOR
	}

	return editor
}
