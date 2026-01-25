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
)

func main() {
	conf := config.MakeConfig()
	output := runner.Run(conf)
	fc, pr, err := parse.Parse(output)
	if err != nil {
		panic(err)
	}
	file, err := outputs.CreateOutputs(fc, pr, *conf)
	if err != nil {
		panic(err)
	}
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	file = filepath.Join(cwd, "csv", file)
	println(file)
	scriptPath := filepath.Join(cwd, "bin", "pyfyle-viewer.sh")
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
