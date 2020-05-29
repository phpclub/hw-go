package main

import (
	"bufio"
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

var (
	ErrFailedReadDir  = errors.New("failed reading directory")
	ErrFailedReadFile = errors.New("failed read file")
)

type Environment map[string]string

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	if !strings.HasSuffix(dir, "/") {
		dir += "/"
	}
	files, err := ioutil.ReadDir(dir)
	deprecatedSymbol := "="
	if err != nil {
		return Environment{}, ErrFailedReadDir
	}
	// Создадим map с размером на кол-во найденных файлов
	cfgEnv := make(map[string]string, len(files))
	for _, file := range files {
		// имя файла не должно содержать =
		if !strings.ContainsAny(file.Name(), deprecatedSymbol) {
			// Получим первую строку из файла
			content, err := readLineFromFile(dir + file.Name())
			if err != nil {
				return Environment{}, ErrFailedReadFile
			}
			cfgEnv[file.Name()] = content
		}
	}
	return cfgEnv, nil
}

func readLineFromFile(path string) (str string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		str = cleanText(sc.Text())
		err = sc.Err()
		return
	}
	return "", err
}

func cleanText(s string) string {
	// терминальные нули (0x00) заменяются на перевод строки (\n).
	s = strings.Replace(s, "\x00", "\n", -1)
	// пробелы и табуляция в конце удаляются.
	s = strings.TrimRight(s, "\t")
	s = strings.TrimRight(s, " ")
	return s
}
