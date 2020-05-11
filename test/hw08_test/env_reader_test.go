package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	t.Run("directory not exists", func(t *testing.T) {
		env, err := ReadDir("testdata/bad")
		require.Nil(t, env, "env is not nil")
		require.Equal(t, err, ErrCanNotOpenDir)
	})

	t.Run("read files", func(t *testing.T) {
		expected := Environment{
			"BAR": `bar`,
			"FOO": `   foo
with new line`,
			"HELLO": `"hello"`,
			"UNSET": ``,
		}
		actual, err := ReadDir("testdata/env")
		require.Nil(t, err, "err is not nil")
		require.Equal(t, expected, actual, "result map not match required map")
	})
}
