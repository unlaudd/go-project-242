package pkg

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetPathSize_File(t *testing.T) {
	path := filepath.Join("..", "..", "testdata", "file1.txt")

	result, err := GetPathSize(path, false, false, false)
	require.NoError(t, err)
	require.Contains(t, result, "10")
	require.Contains(t, result, "file1.txt")
}

func TestGetPathSize_Directory(t *testing.T) {
	path := filepath.Join("..", "..", "testdata", "testdir")

	result, err := GetPathSize(path, false, false, false)
	require.NoError(t, err)
	// Директория содержит один файл размером 5 байт
	require.Contains(t, result, "5")
	require.Contains(t, result, "testdir")
}

func TestGetPathSize_HumanReadable(t *testing.T) {
	path := filepath.Join("..", "..", "testdata", "file2.txt")

	result, err := GetPathSize(path, false, true, false)
	require.NoError(t, err)
	require.Contains(t, result, "20B")
	require.Contains(t, result, "file2.txt")
}

func TestGetPathSize_NotExist(t *testing.T) {
	path := filepath.Join("..", "..", "testdata", "nonexistent.txt")

	result, err := GetPathSize(path, false, false, false)
	require.Error(t, err)
	require.Empty(t, result)
}

func TestGetPathSize_TabSeparator(t *testing.T) {
	path := filepath.Join("..", "..", "testdata", "file1.txt")

	result, err := GetPathSize(path, false, false, false)
	require.NoError(t, err)
	require.Contains(t, result, "\t")
}
