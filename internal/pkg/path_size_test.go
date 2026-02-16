package pkg

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

// getTestDataPath возвращает абсолютный путь к файлу в testdata
func getTestDataPath(parts ...string) string {
	wd, _ := os.Getwd()
	// Поднимаемся на 2 уровня вверх: internal/pkg -> internal -> root
	return filepath.Join(append([]string{wd, "..", "..", "testdata"}, parts...)...)
}

func TestGetPathSize_File(t *testing.T) {
	path := getTestDataPath("file1.txt")

	result, err := GetPathSize(path, false, false, false)
	require.NoError(t, err)
	require.Contains(t, result, "10")
	require.Contains(t, result, "file1.txt")
}

func TestGetPathSize_Directory(t *testing.T) {
	path := getTestDataPath("testdir")

	result, err := GetPathSize(path, false, false, false)
	require.NoError(t, err)
	require.Contains(t, result, "5")
	require.Contains(t, result, "testdir")
}

func TestGetPathSize_HumanReadable(t *testing.T) {
	path := getTestDataPath("file2.txt")

	result, err := GetPathSize(path, false, true, false)
	require.NoError(t, err)
	require.Contains(t, result, "20B")
	require.Contains(t, result, "file2.txt")
}

func TestGetPathSize_NotExist(t *testing.T) {
	path := getTestDataPath("nonexistent.txt")

	result, err := GetPathSize(path, false, false, false)
	require.Error(t, err)
	require.Empty(t, result)
}

func TestGetPathSize_TabSeparator(t *testing.T) {
	path := getTestDataPath("file1.txt")

	result, err := GetPathSize(path, false, false, false)
	require.NoError(t, err)
	require.Contains(t, result, "\t")
}
