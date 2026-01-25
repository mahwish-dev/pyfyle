package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"syscall"

	"pyfyle/config"
	"pyfyle/outputs"
	"pyfyle/parse"
	"pyfyle/runner"

	"github.com/charmbracelet/log"
)

func main() {
	conf := config.MakeConfig()
	output := runner.Run(conf)
	fc, pr, err := parse.Parse(output)
	if err != nil {
		log.Fatal(err.Error())
	}
	file, err := outputs.CreateOutputs(fc, pr, *conf)
	if err != nil {
		log.Fatal(err.Error())
	}
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	file = filepath.Join(cwd, "pyfyle", "csv", file)
	scriptPath := filepath.Join(cwd, "pyfyle", "bin", "pyfyle-viewer.sh")
	bashPath, err := exec.LookPath("bash")
	if err != nil {
		panic("Could not find bash executable")
	}
	args := []string{"bash", scriptPath, file}
	env := os.Environ()
	execErr := syscall.Exec(bashPath, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
