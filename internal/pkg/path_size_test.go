package pkg

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetPathSize_File(t *testing.T) {
	// Получаем корень проекта
	wd, _ := os.Getwd()
	root := filepath.Join(wd, "..", "..")

	path := filepath.Join(root, "testdata", "file1.txt")

	result, err := GetPathSize(path, false, false, false)
	require.NoError(t, err)
	require.Contains(t, result, "10")
}
