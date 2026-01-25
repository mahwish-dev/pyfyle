// Package config handles setting of flags
package config

import (
	"errors"
	"os"
	"path"

	toml "github.com/pelletier/go-toml/v2"
	flag "github.com/spf13/pflag"
)

type Config struct {
	FileName         string
	NoVenv           bool
	PythonPath       string
	OutputMarkdown   bool
	Template         string
	DashboardEnabled bool
}

type tomlConfig struct {
	DashboardEnabled bool
}

func MakeConfig() *Config {
	conf := Config{}

	flag.BoolVar(&conf.NoVenv, "noVenv", false, "disable setting venv")
	flag.StringVar(&conf.PythonPath, "PythonPath", getDefaultPython(), "path to python")
	flag.BoolVar(&conf.OutputMarkdown, "outputMarkdown", false, "output markdown")
	flag.StringVar(&conf.Template, "template", "", "path to markdown template")
	flag.StringVar(&conf.FileName, "filename", "", "name of file")
	flag.Parse()
	conf.DashboardEnabled = parseToml()

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

func parseToml() bool {
	var v tomlConfig
	cwd, err := os.Getwd()
	if err != nil {
		// TODO: handle this err
		return false
	}
	pathToToml := path.Join(cwd, "pyfyle.toml")
	bytes, err := os.ReadFile(pathToToml)
	if err != nil {
		// TODO: handle this err
		return false
	}
	err = toml.Unmarshal(bytes, &v)
	if err != nil {
		// TODO: handle this err
		return false
	}
	return v.DashboardEnabled
}
