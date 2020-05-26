package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

type Environment map[string]string

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return Environment{}, fmt.Errorf("failed reading directory: %s", dir)
	}
	for _, file := range files {
		//fmt.Println(file.Name())
		// Получим первую строку из файла
		content, err := ReadLineFromFile(dir + file.Name())
		if err != nil {
			return Environment{}, fmt.Errorf("failed read file: %s", file.Name())
		}
		fmt.Println("file:" + file.Name() + "\nContent:" + content + "#\n")
	}
	t, ok := os.LookupEnv("TERMINAL_EMULATOR")
	fmt.Println(t, ok)
	if ok {
		os.Setenv("TERMINAL_EMULATOR", "hacked ttys")
	}
	t, ok = os.LookupEnv("TERMINAL_EMULATOR")
	fmt.Println(t, ok)
	return nil, nil
}

func ReadLineFromFile(path string) (str string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		str = sc.Text()
		err = sc.Err()
		return
	}
	return "", err
}
