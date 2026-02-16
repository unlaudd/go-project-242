package pkg

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func setupTestData(t *testing.T) string {
	tmpDir := t.TempDir()

	// Обычные файлы в корне
	file1 := filepath.Join(tmpDir, "file1.txt")
	os.WriteFile(file1, []byte("1234567890"), 0644) // 10 байт

	file2 := filepath.Join(tmpDir, "file2.txt")
	os.WriteFile(file2, []byte("12345678901234567890"), 0644) // 20 байт

	// Скрытые файлы в корне
	hiddenFile := filepath.Join(tmpDir, ".hidden")
	os.WriteFile(hiddenFile, []byte("123456789012345678901234567890"), 0644) // 30 байт

	// Директория первого уровня
	testdir := filepath.Join(tmpDir, "testdir")
	os.Mkdir(testdir, 0755)

	nested := filepath.Join(testdir, "nested.txt")
	os.WriteFile(nested, []byte("12345"), 0644) // 5 байт

	// Скрытый файл в директории
	hiddenInDir := filepath.Join(testdir, ".hidden_nested")
	os.WriteFile(hiddenInDir, []byte("1234567890"), 0644) // 10 байт

	// Вложенная директория (второй уровень)
	nestedDir := filepath.Join(testdir, "nesteddir")
	os.Mkdir(nestedDir, 0755)

	deepFile := filepath.Join(nestedDir, "deep.txt")
	os.WriteFile(deepFile, []byte("123456789012345678901234567890"), 0644) // 30 байт

	return tmpDir
}

func TestGetPathSize_File(t *testing.T) {
	tmpDir := setupTestData(t)
	path := filepath.Join(tmpDir, "file1.txt")

	result, err := GetPathSize(path, false, false, false)
	require.NoError(t, err)
	require.Contains(t, result, "10")
	require.Contains(t, result, "file1.txt")
}

func TestGetPathSize_Directory_FirstLevel(t *testing.T) {
	tmpDir := setupTestData(t)
	path := filepath.Join(tmpDir, "testdir")

	// Без рекурсии (только первый уровень)
	result, err := GetPathSize(path, false, false, false)
	require.NoError(t, err)
	require.Contains(t, result, "5") // Только nested.txt
}

func TestGetPathSize_Directory_Recursive(t *testing.T) {
	tmpDir := setupTestData(t)
	path := filepath.Join(tmpDir, "testdir")

	// С рекурсией (все уровни)
	result, err := GetPathSize(path, true, false, false)
	require.NoError(t, err)
	require.Contains(t, result, "35") // nested.txt (5) + deep.txt (30)
}

func TestGetPathSize_Directory_RecursiveWithAll(t *testing.T) {
	tmpDir := setupTestData(t)
	path := filepath.Join(tmpDir, "testdir")

	// С рекурсией и скрытыми файлами
	result, err := GetPathSize(path, true, false, true)
	require.NoError(t, err)
	require.Contains(t, result, "45") // nested.txt (5) + .hidden_nested (10) + deep.txt (30)
}

func TestGetPathSize_Directory_RecursiveRoot(t *testing.T) {
	tmpDir := setupTestData(t)

	// Рекурсивный подсчёт всей тестовой директории
	result, err := GetPathSize(tmpDir, true, false, false)
	require.NoError(t, err)
	// file1.txt (10) + file2.txt (20) + nested.txt (5) + deep.txt (30) = 65
	require.Contains(t, result, "65")
}

func TestGetPathSize_Directory_RecursiveRootWithAll(t *testing.T) {
	tmpDir := setupTestData(t)

	// Рекурсивный подсчёт со скрытыми файлами
	result, err := GetPathSize(tmpDir, true, false, true)
	require.NoError(t, err)
	// file1.txt (10) + file2.txt (20) + .hidden (30) + nested.txt (5) + .hidden_nested (10) + deep.txt (30) = 105
	require.Contains(t, result, "105")
}

func TestGetPathSize_HumanReadable(t *testing.T) {
	tmpDir := setupTestData(t)
	path := filepath.Join(tmpDir, "file2.txt")

	result, err := GetPathSize(path, false, true, false)
	require.NoError(t, err)
	require.Contains(t, result, "20B")
	require.Contains(t, result, "file2.txt")
}

func TestGetPathSize_NotExist(t *testing.T) {
	tmpDir := setupTestData(t)
	path := filepath.Join(tmpDir, "nonexistent.txt")

	result, err := GetPathSize(path, false, false, false)
	require.Error(t, err)
	require.Empty(t, result)
}

func TestGetPathSize_TabSeparator(t *testing.T) {
	tmpDir := setupTestData(t)
	path := filepath.Join(tmpDir, "file1.txt")

	result, err := GetPathSize(path, false, false, false)
	require.NoError(t, err)
	require.Contains(t, result, "\t")
}

func TestGetPathSize_HiddenFiles(t *testing.T) {
	tmpDir := setupTestData(t)

	// Без флага --all скрытые файлы игнорируются
	result, err := GetPathSize(tmpDir, false, false, false)
	require.NoError(t, err)
	require.Contains(t, result, "30") // file1.txt (10) + file2.txt (20)

	// С флагом --all скрытые файлы учитываются
	result, err = GetPathSize(tmpDir, false, false, true)
	require.NoError(t, err)
	require.Contains(t, result, "60") // file1.txt (10) + file2.txt (20) + .hidden (30)
}

func TestIsHidden(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		expected bool
	}{
		{"hidden file", ".hidden", true},
		{"hidden dir", ".git", true},
		{"normal file", "file.txt", false},
		{"normal dir", "src", false},
		{"dot in middle", "file.txt.bak", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isHidden(tt.filename)
			require.Equal(t, tt.expected, result)
		})
	}
}

// Тесты для FormatSize
func TestFormatSize_HumanFalse(t *testing.T) {
	tests := []struct {
		name     string
		size     int64
		expected string
	}{
		{"zero", 0, "0B"},
		{"bytes", 123, "123B"},
		{"kilobytes", 1234, "1234B"},
		{"megabytes", 123456, "123456B"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatSize(tt.size, false)
			require.Equal(t, tt.expected, result)
		})
	}
}

func TestFormatSize_HumanTrue(t *testing.T) {
	tests := []struct {
		name     string
		size     int64
		expected string
	}{
		{"zero", 0, "0B"},
		{"bytes", 123, "123B"},
		{"kilobytes", 1234, "1.2KB"},
		{"megabytes", 1234567, "1.2MB"},
		{"gigabytes", 1234567890, "1.1GB"},
		{"terabytes", 1234567890123, "1.1TB"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatSize(tt.size, true)
			require.Equal(t, tt.expected, result)
		})
	}
}
