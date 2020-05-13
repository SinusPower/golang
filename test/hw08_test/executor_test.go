package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	t.Run("return codes", func(t *testing.T) {
		f, err := os.Create("testdata/test")
		if err != nil {
			t.Fatal(err)
		}

		_, err = f.Write([]byte("File contents."))
		if err != nil {
			t.Fatal(err)
		}
		f.Close()

		cmd := []string{"cat", "testdata/test"}
		env := Environment{}
		returnCode := RunCmd(cmd, env)
		require.Equal(t, 0, returnCode)

		cmd = []string{"cat", "testdata/notExist"}
		returnCode = RunCmd(cmd, env)
		require.Equal(t, 1, returnCode)

		err = os.Remove("testdata/test")
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("empty cmd", func(t *testing.T) {
		cmd := []string{}
		env := Environment{}
		returnCode := RunCmd(cmd, env)
		require.Equal(t, 0, returnCode)
	})
}
