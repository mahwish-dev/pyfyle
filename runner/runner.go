// Package runner well uhh runs cprofile
package runner

import (
	"fmt"
	"os/exec"

	"pyfyle/config"

	"github.com/charmbracelet/log"
)

func Run(conf *config.Config) string {
	python := conf.PythonPath
	file := conf.FileName
	if conf.NoVenv {
		python = "python"
	}
	log.Info(fmt.Sprintf("Going to run with python = %s", python))
	cmd := exec.Command(python, "-m", "cProfile", file)
	output, err := cmd.Output()
	if err != nil {
		// TODO: handle this err
		log.Error(err.Error())
	}
	log.Info("Ran Command")
	return string(output)
}
