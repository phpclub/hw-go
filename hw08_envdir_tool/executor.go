package main

import (
	"log"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	var customCmd *exec.Cmd
	if len(cmd) == 0 {
		// Пустой []
		return 1
	}

	if len(cmd[1:]) > 0 {
		customCmd = exec.Command(cmd[0], cmd[1:]...) //nolint:gosec
	} else {
		customCmd = exec.Command(cmd[0]) //nolint:gosec
	}

	for key, val := range env {
		_, ok := os.LookupEnv(key)
		// Если такая есть - удалим из окружения
		if ok {
			err := os.Unsetenv(key)
			if err != nil {
				return 1
			}
		}
		// Добавляем в новые переменные если они не пустые
		if len(val) > 0 {
			err := os.Setenv(key, val)
			if err != nil {
				return 1
			}
		}
	}
	customCmd.Env = os.Environ()
	customCmd.Stdout = os.Stdout
	customCmd.Stderr = os.Stderr
	err := customCmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	return 0
}
