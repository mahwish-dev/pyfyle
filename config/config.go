// Package config handles setting of flags
package config

import flag "github.com/spf13/pflag"

type Config struct {
	NoVenv          bool
	VenvPath        string
	OutputMarkdown  bool
	Template        string
	IncludeBuiltins bool
	SortBy          string
}

func MakeConfig() *Config {
	conf := Config{}
	flag.BoolVar(&conf.NoVenv, "noVenv", false, "disable setting venv")
	flag.StringVar(&conf.VenvPath, "venvPath", "", "path to venv")
	flag.BoolVar(&conf.OutputMarkdown, "outputMarkdown", false, "output markdown")
	flag.StringVar(&conf.Template, "template", "", "path to markdown template")
	flag.BoolVar(&conf.IncludeBuiltins, "includeBuiltins", true, "includeBuiltins")
	flag.StringVar(&conf.SortBy, "sortby", "", "sortby")
	flag.Parse()

	return &conf
}
