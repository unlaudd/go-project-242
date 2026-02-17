package code

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func setupTestData(t *testing.T) string {
	tmpDir := t.TempDir()

	file1 := filepath.Join(tmpDir, "file1.txt")
	err := os.WriteFile(file1, []byte("1234567890"), 0644)
	require.NoError(t, err)

	file2 := filepath.Join(tmpDir, "file2.txt")
	err = os.WriteFile(file2, []byte("12345678901234567890"), 0644)
	require.NoError(t, err)

	hiddenFile := filepath.Join(tmpDir, ".hidden")
	err = os.WriteFile(hiddenFile, []byte("123456789012345678901234567890"), 0644)
	require.NoError(t, err)

	testdir := filepath.Join(tmpDir, "testdir")
	err = os.Mkdir(testdir, 0755)
	require.NoError(t, err)

	nested := filepath.Join(testdir, "nested.txt")
	err = os.WriteFile(nested, []byte("12345"), 0644)
	require.NoError(t, err)

	hiddenInDir := filepath.Join(testdir, ".hidden_nested")
	err = os.WriteFile(hiddenInDir, []byte("1234567890"), 0644)
	require.NoError(t, err)

	nestedDir := filepath.Join(testdir, "nesteddir")
	err = os.Mkdir(nestedDir, 0755)
	require.NoError(t, err)

	deepFile := filepath.Join(nestedDir, "deep.txt")
	err = os.WriteFile(deepFile, []byte("123456789012345678901234567890"), 0644)
	require.NoError(t, err)

	emptyFile := filepath.Join(tmpDir, "empty.txt")
	err = os.WriteFile(emptyFile, []byte(""), 0644)
	require.NoError(t, err)

	return tmpDir
}

func TestGetPathSize_File(t *testing.T) {
	tmpDir := setupTestData(t)
	path := filepath.Join(tmpDir, "file1.txt")

	result, err := GetPathSize(path, false, false, false)
	require.NoError(t, err)
	require.Equal(t, "10B", result)
}

func TestGetPathSize_Directory(t *testing.T) {
	tmpDir := setupTestData(t)
	path := filepath.Join(tmpDir, "testdir")

	result, err := GetPathSize(path, false, false, false)
	require.NoError(t, err)
	require.Equal(t, "5B", result)
}

func TestGetPathSize_Recursive(t *testing.T) {
	tmpDir := setupTestData(t)
	path := filepath.Join(tmpDir, "testdir")

	result, err := GetPathSize(path, true, false, false)
	require.NoError(t, err)
	require.Equal(t, "35B", result)
}

func TestGetPathSize_Hidden(t *testing.T) {
	tmpDir := setupTestData(t)

	result, err := GetPathSize(tmpDir, false, false, false)
	require.NoError(t, err)
	require.Equal(t, "30B", result)

	result, err = GetPathSize(tmpDir, false, false, true)
	require.NoError(t, err)
	require.Equal(t, "60B", result)
}

func TestGetPathSize_HumanReadable(t *testing.T) {
	tmpDir := setupTestData(t)
	path := filepath.Join(tmpDir, "file2.txt")

	result, err := GetPathSize(path, false, true, false)
	require.NoError(t, err)
	require.Equal(t, "20B", result)
}

func TestGetPathSize_NotExist(t *testing.T) {
	tmpDir := setupTestData(t)
	path := filepath.Join(tmpDir, "nonexistent.txt")

	result, err := GetPathSize(path, false, false, false)
	require.Error(t, err)
	require.Empty(t, result)
}

func TestGetPathSize_EmptyFile(t *testing.T) {
	tmpDir := setupTestData(t)
	path := filepath.Join(tmpDir, "empty.txt")

	result, err := GetPathSize(path, false, false, false)
	require.NoError(t, err)
	require.Equal(t, "0B", result)
}

func TestGetPathSize_EmptyDir(t *testing.T) {
	tmpDir := setupTestData(t)
	emptyDir := filepath.Join(tmpDir, "emptydir")
	err := os.Mkdir(emptyDir, 0755)
	require.NoError(t, err)

	result, err := GetPathSize(emptyDir, false, false, false)
	require.NoError(t, err)
	require.Equal(t, "0B", result)
}
