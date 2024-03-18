package main

import (
	"fmt"
	"strings"

	"github.com/bayashi/colorpalette"
)

var zeroByte = []byte("")

func process(origLine []byte, rule []*Rule) ([]byte, error) {
	var err error
	line := origLine
	for _, r := range rule {
		switch r.Target {
		case TARGET_WORD:
			line, err = processWord(line, r)
			if err != nil {
				return zeroByte, err
			}
		case TARGET_LINE:
			line, err = processLine(line, r)
			if err != nil {
				return zeroByte, err
			}
		default:
			return zeroByte, fmt.Errorf("wrong target %s", r.Target)
		}
	}

	return line, nil
}

func processWord(line []byte, r *Rule) ([]byte, error) {
	switch r.Type {
	case TYPE_MATCH:
		if strings.Contains(strings.TrimSpace(string(line)), r.Match) {
			coloredMatch := colorpalette.Get(r.Color).Sprintf(r.Match)
			lineString := strings.ReplaceAll(string(line), r.Match, coloredMatch)
			byteString := []byte(lineString)
			return byteString, nil
		}
	case TYPE_REGEXP:
		coloredMatch := colorpalette.Get(r.Color).Sprintf("$1")
		return r.regexp.ReplaceAll(line, []byte(coloredMatch)), nil
	default:
		return zeroByte, fmt.Errorf("wrong type %s for target=%s", r.Type, r.Target)
	}

	return line, nil
}

func processLine(line []byte, r *Rule) ([]byte, error) {
	switch r.Type {
	case TYPE_MATCH:
		if strings.Contains(strings.TrimSpace(string(line)), r.Match) {
			coloredLine := colorpalette.Get(r.Color).Sprintf("%s", line)
			return []byte(coloredLine), nil
		}
	case TYPE_REGEXP:
		if r.regexp.Match(line) {
			coloredLine := colorpalette.Get(r.Color).Sprintf("%s", line)
			return []byte(coloredLine), nil
		}
	case TYPE_PREFIX:
		if strings.HasPrefix(strings.TrimSpace(string(line)), r.Match) {
			coloredLine := colorpalette.Get(r.Color).Sprintf("%s", line)
			return []byte(coloredLine), nil
		}
	case TYPE_SUFFIX:
		if strings.HasSuffix(strings.TrimSpace(string(line)), r.Match) {
			coloredLine := colorpalette.Get(r.Color).Sprintf("%s", line)
			return []byte(coloredLine), nil
		}
	default:
		return zeroByte, fmt.Errorf("wrong type %s for target=%s", r.Type, r.Type)
	}

	return line, nil
}
