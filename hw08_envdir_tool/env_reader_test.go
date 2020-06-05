package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	testDir := path + "/testdata/env"
	expectedMap := make(map[string]string)

	t.Run("simple case", func(t *testing.T) {
		expectedMap["BAR"] = "bar"
		expectedMap["FOO"] = "   foo\nwith new line"
		expectedMap["HELLO"] = "\"hello\""
		expectedMap["UNSET"] = ""

		env, ok := ReadDir(testDir)

		require.NoError(t, ok)
		for k, val := range expectedMap {
			require.Equal(t, val, env[k])
		}
	})

	t.Run("Wrong dir case", func(t *testing.T) {
		testWrongDir := "/dev/null"

		env, ok := ReadDir(testWrongDir)

		require.Error(t, ok)
		require.Len(t, env, 0)
	})

	t.Run("Empty dir case", func(t *testing.T) {
		testEmptyDir := testDir + "/tmp"
		err := os.Mkdir(testEmptyDir, 0777)
		if err != nil {
			panic(err)
		}
		defer os.Remove(testEmptyDir)

		env, ok := ReadDir(testEmptyDir)

		require.Nil(t, ok)
		require.Len(t, env, 0)
	})
}
