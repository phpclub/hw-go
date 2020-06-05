package main

import (
	"errors"
	"log"
	"os"
)

var (
	ErrMissingDestinationDir = errors.New("missing destination dir")
	ErrMissingComand         = errors.New("missing command to execute")
)

func main() {
	switch len(os.Args) {
	// Выведем ошибку если нет переданных аргументов
	case 1:
		log.Fatal(ErrMissingDestinationDir)
	case 2: //nolint:gomnd
		log.Fatal(ErrMissingComand)
	}
	dirPath, customCmd := os.Args[1], os.Args[2:]

	cmdEnv, err := ReadDir(dirPath)

	if err != nil {
		log.Fatal(err)
	}

	RunCmd(customCmd, cmdEnv)
}
