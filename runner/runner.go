// Package runner well uhh runs cprofile
package runner

import (
	"os/exec"

	"pyfyle/config"
)

func Run(conf *config.Config) string {
	python := conf.PythonPath
	file := conf.FileName
	if conf.NoVenv {
		python = "python"
	}
	cmd := exec.Command(python, "-m", "cProfile", file)
	output, err := cmd.Output()
	if err != nil {
		// TODO: handle this err
	}
	return string(output)
}
