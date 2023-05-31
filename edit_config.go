package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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

	if err := createConfigDir(file); err != nil {
		return err.Error()
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

func createConfigDir(dirPath string) error {
	d := filepath.Dir(dirPath)
	dir, err := os.Stat(d)
	if err != nil || !dir.IsDir() {
		err2 := os.MkdirAll(d, os.FileMode(0700)&os.ModePerm)
		if err2 != nil {
			return err2
		}
	}

	return nil
}
