package main

import (
	"log"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env
func RunCmd(cmd []string, env Environment) (returnCode int) {
	// set environment
	for key, value := range env {
		if value == "" {
			os.Unsetenv(key)
		} else {
			_, ok := os.LookupEnv(key)
			if ok {
				os.Unsetenv(key)
			}
			os.Setenv(key, value)
		}
	}

	// prepare command
	command := exec.Command(cmd[0], cmd[0:]...)
	command.Env = os.Environ() // ?
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	if err := command.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			return exitError.ExitCode()
		} else {
			log.Fatal(err)
		}
	}
	return
}
