package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"

	flag "github.com/spf13/pflag"
)

var (
	version     = ""
	installFrom = "Source"
)

type options struct {
	rule string
}

func parseArgs() *options {
	o := &options{}

	var (
		flagHelp           bool
		flagVersion        bool
		flagDumpSchema     bool
		flagDumpColors     bool
		flagDumpConfigPath bool
		flagEditConfig     bool
		flagDumpRule       bool
		flagDumpRules      bool
	)

	flag.BoolVarP(&flagHelp, "help", "h", false, "Show help (This message) and exit")
	flag.BoolVarP(&flagVersion, "version", "v", false, "Show version and build info and exit")
	flag.BoolVarP(&flagDumpSchema, "dump-schema", "", false, "Dump JSON Schema to validate the rule YAML config file")
	flag.BoolVarP(&flagDumpColors, "dump-colors", "", false, "Dump color palette")
	flag.BoolVarP(&flagDumpConfigPath, "dump-config-path", "", false, "Dump config file path")
	flag.BoolVarP(&flagEditConfig, "edit-config", "", false, "Invoke $EDITOR (or vi) to edit config YAML file")
	flag.BoolVarP(&flagDumpRule, "dump-rule", "", false, "Dump specified rule")
	flag.BoolVarP(&flagDumpRules, "dump-rules", "", false, "Show rules from config file")

	flag.Parse()

	if flagHelp {
		putHelp(fmt.Sprintf("Version %s", getVersion()))
	}

	if flagVersion {
		putErr(versionDetails())
		os.Exit(exitOK)
	}

	if flagDumpSchema {
		fmt.Println(dumpJSONSchema())
		os.Exit(exitOK)
	}

	if flagDumpColors {
		fmt.Println(dumpColorPalette())
		os.Exit(exitOK)
	}

	if flagDumpConfigPath {
		fmt.Println(cfgFilePath(irirConfigFiles[0]))
		os.Exit(exitOK)
	}

	if flagEditConfig {
		ret := editConfig(cfgFilePath(irirConfigFiles[0]))
		if ret != "" {
			fmt.Println(ret)
			os.Exit(exitErr)
		}
		os.Exit(exitOK)
	}

	if flagDumpRules {
		fmt.Println(dumpRules())
		os.Exit(exitOK)
	}

	o.targetRule()

	if flagDumpRule {
		fmt.Println(dumpRule(o.rule))
		os.Exit(exitOK)
	}

	if len(flag.Args()) == 0 && o.rule == "" {
		putHelp(fmt.Sprintf("Version %s", getVersion()))
	}

	return o
}

func versionDetails() string {
	goos := runtime.GOOS
	goarch := runtime.GOARCH
	compiler := runtime.Version()

	return fmt.Sprintf(
		"Version %s - %s.%s (compiled:%s, %s)",
		getVersion(),
		goos,
		goarch,
		compiler,
		installFrom,
	)
}

func getVersion() string {
	if version != "" {
		return version
	}
	i, ok := debug.ReadBuildInfo()
	if !ok {
		return "Unknown"
	}

	return i.Main.Version
}

func (o *options) targetRule() {
	for _, arg := range flag.Args() {
		if o.rule != "" {
			putHelp(fmt.Sprintf("Err: Wrong args. Unnecessary arg [%s]", arg))
		}
		if arg == "-" {
			continue
		}
		o.rule = arg
	}
}
