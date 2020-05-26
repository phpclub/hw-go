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
func parseArgs() (dir string, command string, args map[int]string, err error) {
	// Проверим на наличие обязательных аргументов
	incomingArgs := os.Args[1:]
	if len(incomingArgs) >= requiredArgs {
		if len(incomingArgs[0]) > 0 {
			dir = incomingArgs[0]
		}
		if len(incomingArgs[1]) > 0 {
			command = incomingArgs[1]
		}
	} else {
		switch len(incomingArgs) {
		case 0:
			err = ErrMissingDestinationDir
		case 1:
			err = ErrMissingComand
		}
		return dir, command, args, err
	}
	if len(incomingArgs[2:]) > 0 {
		// Создадим map c переданными аргументами
		args = make(map[int]string, len(incomingArgs[2:]))
		for k, v := range incomingArgs[2:] {
			args[k] = v
		}
	}
	return dir, command, args, err
}

func main() {
	dir, cmd, args, err := parseArgs()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(dir, cmd, args)
	cmdEnv, err := ReadDir(dir)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(cmdEnv)

}
