package main

import (
	"errors"
	"fmt"
	"os"
)

var (
	ErrMissingDestinationDir = errors.New("missing destination dir")
	ErrMissingComand         = errors.New("missing command to execute")
)

const requiredArgs = 2

// получение аргументов из командной строки.
func parseArgs() (dir string, command []string, err error) {
	// Проверим на наличие обязательных аргументов
	parseCommand := make([]string, 0)
	incomingArgs := os.Args[1:]
	if len(incomingArgs) >= requiredArgs {
		if len(incomingArgs[0]) > 0 {
			dir = incomingArgs[0]
		}
		if len(incomingArgs[1]) > 0 {
			parseCommand = append(parseCommand, incomingArgs[1])
		}
	} else {
		switch len(incomingArgs) {
		case 0:
			err = ErrMissingDestinationDir
		case 1:
			err = ErrMissingComand
		}
		return dir, parseCommand, err
	}
	if len(incomingArgs[2:]) > 0 {
		parseCommand = append(parseCommand, incomingArgs[2:]...)
	}
	return dir, parseCommand, err
}

func main() {
	dir, customCmd, err := parseArgs()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	cmdEnv, err := ReadDir(dir)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	RunCmd(customCmd, cmdEnv)
}
