package locfs

import (
	"embed"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//go:embed test/*
var testFS embed.FS

func TestFSPath(t *testing.T) {
	t.Run("test_file", func(t *testing.T) {
		a, err := FSPath(testFS, "test/dir/file.txt")
		require.NoError(t, err)
		expByte, err := testFS.ReadFile("test/dir/file.txt")
		require.NoError(t, err)
		actual, err := ioutil.ReadFile(a)
		assert.Equal(t, expByte, actual)
	})
	t.Run("test_dir", func(t *testing.T) {
		expByte, err := testFS.ReadFile("test/dir/file.txt")
		require.NoError(t, err)
		a, err := FSPath(testFS, "test/dir")
		require.NoError(t, err)
		actual, err := ioutil.ReadFile(filepath.Join(a, "file.txt"))
		require.NoError(t, err)
		require.NotNil(t, actual)
		assert.Equal(t, expByte, actual)
	})

	t.Run("test_dir", func(t *testing.T) {
		expByte, err := testFS.ReadFile("test/dir2/file2.txt")
		require.NoError(t, err)
		a := FSPathMust(testFS, "test/dir2")
		actual, err := ioutil.ReadFile(filepath.Join(a, "file2.txt"))
		require.NoError(t, err)
		require.NotNil(t, actual)
		assert.Equal(t, expByte, actual)
	})
}
