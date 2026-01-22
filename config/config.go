// Package config handles setting of flags
package config

import (
	"errors"
	"os"
	"path"

	flag "github.com/spf13/pflag"
)

type Config struct {
	NoVenv          bool
	PythonPath      string
	OutputMarkdown  bool
	Template        string
	IncludeBuiltins bool
	SortBy          string
}

func MakeConfig() *Config {
	conf := Config{}

	flag.BoolVar(&conf.NoVenv, "noVenv", false, "disable setting venv")
	flag.StringVar(&conf.PythonPath, "PythonPath", getDefaultPython(), "path to python")
	flag.BoolVar(&conf.OutputMarkdown, "outputMarkdown", false, "output markdown")
	flag.StringVar(&conf.Template, "template", "", "path to markdown template")
	flag.BoolVar(&conf.IncludeBuiltins, "includeBuiltins", true, "includeBuiltins")
	flag.StringVar(&conf.SortBy, "sortby", "", "sortby")
	flag.Parse()

	return &conf
}

func getDefaultPython() string {
	cwd, err := os.Getwd()
	if err != nil {
		// TODO: handle this err
	}
	pathToPy := path.Join(cwd, ".venv", "bin", "python")
	if _, err := os.Stat(pathToPy); errors.Is(err, os.ErrNotExist) {
		return "python"
	}
	return pathToPy
}
