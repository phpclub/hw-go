package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {

	t.Run("Empty param case", func(t *testing.T) {
		testCmd := []string{}
		testEnv := make(map[string]string)

		returnCode := RunCmd(testCmd, testEnv)

		require.Equal(t, 1, returnCode)
	})

	t.Run("Simple case", func(t *testing.T) {
		testEnv := make(map[string]string)
		testCmd := make([]string, 0)
		testCmd = append(testCmd, "ls")

		returnCode := RunCmd(testCmd, testEnv)

		require.Equal(t, 0, returnCode)
	})

	t.Run("Extended case", func(t *testing.T) {
		testEnv := make(map[string]string)
		testCmd := make([]string, 0)
		testCmd = append(testCmd, "ls", "README.md")

		returnCode := RunCmd(testCmd, testEnv)

		require.Equal(t, 0, returnCode)
	})

	t.Run("With environment case", func(t *testing.T) {
		tempFile := "test.tmp"
		defer os.Remove(tempFile)
		testEnv := make(map[string]string)
		testCmd := make([]string, 0)
		testCmd = append(testCmd, "./testEnv.sh")
		testEnv["TEST"] = "ENV_OK"

		returnCode := RunCmd(testCmd, testEnv)

		require.Equal(t, 0, returnCode)
		b, err := ioutil.ReadFile(tempFile)
		require.NoError(t, err)
		require.Equal(t, "ENV_OK\n", string(b))
	})
}
