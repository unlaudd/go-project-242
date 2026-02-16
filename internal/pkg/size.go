package pkg

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// GetPathSize возвращает размер файла или директории в формате "<размер> <путь>"
func GetPathSize(path string, recursive, human, all bool) (string, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return "", err
	}

	var size int64

	if !info.IsDir() {
		// Это файл
		size = info.Size()
	} else {
		// Это директория
		if recursive {
			size = getDirSizeRecursive(path, all)
		} else {
			size = getDirSizeFirstLevel(path, all)
		}
	}

	// Форматируем размер
	sizeStr := FormatSize(size, human)

	return fmt.Sprintf("%s\t%s", sizeStr, path), nil
}

// getDirSizeFirstLevel считает размер файлов только первого уровня
func getDirSizeFirstLevel(path string, all bool) int64 {
	entries, err := os.ReadDir(path)
	if err != nil {
		return 0
	}

	var size int64
	for _, entry := range entries {
		// Пропускаем скрытые файлы, если флаг --all не установлен
		if !all && isHidden(entry.Name()) {
			continue
		}

		entryInfo, err := entry.Info()
		if err != nil {
			continue
		}
		if !entryInfo.IsDir() {
			size += entryInfo.Size()
		}
	}

	return size
}

// getDirSizeRecursive рекурсивно считает размер всех файлов в директории
func getDirSizeRecursive(path string, all bool) int64 {
	var size int64

	err := filepath.WalkDir(path, func(currentPath string, d os.DirEntry, err error) error {
		if err != nil {
			return nil // Продолжаем обход даже при ошибках
		}

		// Пропускаем скрытые файлы/директории, если флаг --all не установлен
		if !all && isHidden(d.Name()) {
			if d.IsDir() {
				return filepath.SkipDir // Пропускаем всю скрытую директорию
			}
			return nil // Пропускаем скрытый файл
		}

		// Считаем только файлы (не директории)
		if !d.IsDir() {
			info, err := d.Info()
			if err == nil {
				size += info.Size()
			}
		}

		return nil
	})

	if err != nil {
		return 0
	}

	return size
}

// isHidden проверяет, является ли файл или директория скрытыми
func isHidden(name string) bool {
	return strings.HasPrefix(name, ".")
}

// FormatSize форматирует размер в человекочитаемый вид
func FormatSize(size int64, human bool) string {
	if !human {
		return fmt.Sprintf("%dB", size)
	}

	const unit = 1024
	if size < unit {
		return fmt.Sprintf("%dB", size)
	}

	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}

	units := []string{"KB", "MB", "GB", "TB", "PB", "EB"}
	return fmt.Sprintf("%.1f%s", float64(size)/float64(div), units[exp-1])
}
