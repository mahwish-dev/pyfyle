// Package config handles setting of flags
package config

import flag "github.com/spf13/pflag"

type Config struct {
	noVenv          bool
	venvPath        string
	outputMarkdown  bool
	template        string
	includeBuiltins bool
	sortby          string
}

func MakeConfig() *Config {
	conf := Config{}
	flag.BoolVar(&conf.noVenv, "noVenv", false, "disable setting venv")
	flag.StringVar(&conf.venvPath, "venvPath", "", "path to venv")
	flag.BoolVar(&conf.outputMarkdown, "outputMarkdown", false, "output markdown")
	flag.StringVar(&conf.template, "template", "", "path to markdown template")
	flag.BoolVar(&conf.includeBuiltins, "includeBuiltins", true, "includeBuiltins")
	flag.StringVar(&conf.sortby, "sortby", "", "sortby")
	flag.Parse()

	return &conf
}
