package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"

	"github.com/adrg/xdg"
	"golang.org/x/term"
)


var cfgFilePath = func(fileName string) string {
	return filepath.Join(xdg.ConfigHome, irirDir, fileName)
}

func run() error {
	o := parseArgs()

	rule, err := loadRule(cfgFilePath, o.rule)
	if err != nil {
		return err
	}

	if o.wrapCmdName != "" {
		return wrapCommand(o, rule)
	}

	if term.IsTerminal(int(syscall.Stdin)) {
		os.Exit(exitOK)
	}

	routine(os.Stdin, rule)

	return nil
}

func wrapCommand(o *options, rule []*Rule) error {
	cmd := exec.Command(o.wrapCmdName, o.wrapCmdArgs...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("%s pipe error %#v, %w", irir, cmd.String(), err)
	}

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("%s error at start %#v, %w", irir, cmd.String(), err)
	}

	if err := routine(stdout, rule); err != nil {
		return fmt.Errorf("%s stdout error %#v, %w", irir, cmd.String(), err)
	}

    if err := cmd.Wait(); err != nil {
        return fmt.Errorf("%s exec error %#v, %w", irir, cmd.String(), err)
    }

	return nil
}

func routine(r io.Reader, rule []*Rule) error {
	s := bufio.NewScanner(r)
	for s.Scan() {
		line := s.Bytes()
		result, err := process(line, rule)
		if err != nil {
			putErr(err.Error())
			result = line
		}
		if _, err := os.Stdout.Write(result); err != nil {
			return err
		}
		if _, err := os.Stdout.WriteString("\n"); err != nil {
			return fmt.Errorf("fail to write string %w", err)
		}
	}

	return nil
}
