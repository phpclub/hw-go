package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// Удаляем файлы сгенеренные тестами.
func cleanUp(file string) {
	err := os.Remove(file)
	if err != nil {
		panic(err)
	}
}

// Получим размер файла для тестов.
func getFileSize(file string) int64 {
	iFile, err := os.OpenFile(file, os.O_RDONLY, 0)
	if err != nil {
		panic(err)
	}
	defer iFile.Close()
	iFileStat, err := iFile.Stat()
	if err != nil {
		panic(err)
	}
	return iFileStat.Size()
}

func TestCopy(t *testing.T) {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	testInputFile := path + "/testdata/input.txt"
	testOutputFile := path + "/test_output.txt"
	testOutputSeekFile := path + "/testdata/out_offset100_limit1000.txt"

	t.Run("limit 0", func(t *testing.T) {
		ok := Copy(testInputFile, testOutputFile, 0, 0)
		require.NoError(t, ok)
		require.Equal(t, getFileSize(testInputFile), getFileSize(testOutputFile))
		cleanUp(testOutputFile)
	})

	t.Run("limit 10", func(t *testing.T) {
		ok := Copy(testInputFile, testOutputFile, 0, 10)
		require.NoError(t, ok)
		require.Equal(t, int64(10), getFileSize(testOutputFile))
		cleanUp(testOutputFile)
	})

	t.Run("limit 10000", func(t *testing.T) {
		ok := Copy(testInputFile, testOutputFile, 0, 10000)
		require.NoError(t, ok)
		require.Equal(t, getFileSize(testInputFile), getFileSize(testOutputFile))
		cleanUp(testOutputFile)
	})

	t.Run("offset 100 limit 1000", func(t *testing.T) {
		ok := Copy(testInputFile, testOutputFile, 100, 1000)
		require.NoError(t, ok)
		require.Equal(t, int64(1000), getFileSize(testOutputFile))
		//Сравним testOutputSeekFile и testOutputFile
		outSeek, err := ioutil.ReadFile(testOutputSeekFile)
		if err != nil {
			panic(err)
		}
		outFile, err := ioutil.ReadFile(testOutputFile)
		if err != nil {
			panic(err)
		}
		require.Equal(t, outSeek, outFile)
		cleanUp(testOutputFile)
	})

	t.Run("offset 7000", func(t *testing.T) {
		ok := Copy(testInputFile, testOutputFile, 7000, 0)
		require.Error(t, ok)
		require.EqualError(t, ok, "offset exceeds file size")
	})

	t.Run("Unsupported File", func(t *testing.T) {
		ok := Copy("/dev/urandom", testOutputFile, 7000, 0)
		require.Error(t, ok)
		require.EqualError(t, ok, "unsupported file")
	})
}
