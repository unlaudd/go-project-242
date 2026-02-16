package tests

import (
	"os"
	"path/filepath"
	"testing"

	"hexlet-path-size/pkg"

	"github.com/stretchr/testify/require"
)

func setupTestData(t *testing.T) string {
	tmpDir := t.TempDir()

	// Обычные файлы в корне (10 + 20 = 30 байт)
	file1 := filepath.Join(tmpDir, "file1.txt")
	os.WriteFile(file1, []byte("1234567890"), 0644) // 10 байт

	file2 := filepath.Join(tmpDir, "file2.txt")
	os.WriteFile(file2, []byte("12345678901234567890"), 0644) // 20 байт

	// Скрытый файл в корне (30 байт)
	hiddenFile := filepath.Join(tmpDir, ".hidden")
	os.WriteFile(hiddenFile, []byte("123456789012345678901234567890"), 0644) // 30 байт

	// Директория с файлами
	testdir := filepath.Join(tmpDir, "testdir")
	os.Mkdir(testdir, 0755)

	nested := filepath.Join(testdir, "nested.txt")
	os.WriteFile(nested, []byte("12345"), 0644) // 5 байт

	// Скрытый файл в директории
	hiddenInDir := filepath.Join(testdir, ".hidden_nested")
	os.WriteFile(hiddenInDir, []byte("1234567890"), 0644) // 10 байт

	// Вложенная директория (для рекурсивных тестов)
	nestedDir := filepath.Join(testdir, "nesteddir")
	os.Mkdir(nestedDir, 0755)

	deepFile := filepath.Join(nestedDir, "deep.txt")
	os.WriteFile(deepFile, []byte("123456789012345678901234567890"), 0644) // 30 байт

	return tmpDir
}

func TestGetPathSize_File(t *testing.T) {
	tmpDir := setupTestData(t)
	path := filepath.Join(tmpDir, "file1.txt")

	result, err := pkg.GetPathSize(path, false, false, false)
	require.NoError(t, err)
	require.Contains(t, result, "10")
}

func TestGetPathSize_Directory(t *testing.T) {
	tmpDir := setupTestData(t)
	path := filepath.Join(tmpDir, "testdir")

	result, err := pkg.GetPathSize(path, false, false, false)
	require.NoError(t, err)
	require.Contains(t, result, "5")
}

func TestGetPathSize_Recursive(t *testing.T) {
	tmpDir := setupTestData(t)
	path := filepath.Join(tmpDir, "testdir")

	result, err := pkg.GetPathSize(path, true, false, false)
	require.NoError(t, err)
	require.Contains(t, result, "35") // nested.txt (5) + deep.txt (30)
}

func TestGetPathSize_Hidden(t *testing.T) {
	tmpDir := setupTestData(t)

	// Без флага --all (скрытые файлы игнорируются)
	// file1.txt (10) + file2.txt (20) = 30 байт
	result, err := pkg.GetPathSize(tmpDir, false, false, false)
	require.NoError(t, err)
	require.Contains(t, result, "30")

	// С флагом --all (скрытые файлы учитываются)
	// file1.txt (10) + file2.txt (20) + .hidden (30) = 60 байт
	result, err = pkg.GetPathSize(tmpDir, false, false, true)
	require.NoError(t, err)
	require.Contains(t, result, "60")
}

func TestGetPathSize_HumanReadable(t *testing.T) {
	tmpDir := setupTestData(t)
	path := filepath.Join(tmpDir, "file2.txt")

	result, err := pkg.GetPathSize(path, false, true, false)
	require.NoError(t, err)
	require.Contains(t, result, "20B")
}

func TestGetPathSize_NotExist(t *testing.T) {
	tmpDir := setupTestData(t)
	path := filepath.Join(tmpDir, "nonexistent.txt")

	result, err := pkg.GetPathSize(path, false, false, false)
	require.Error(t, err)
	require.Empty(t, result)
}
