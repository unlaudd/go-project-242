package pkg

import (
	"fmt"
	"os"
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
		// Это директория — суммируем размеры файлов первого уровня
		entries, err := os.ReadDir(path)
		if err != nil {
			return "", err
		}

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
	}

	// Форматируем размер
	sizeStr := FormatSize(size, human)

	return fmt.Sprintf("%s\t%s", sizeStr, path), nil
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

	units := "KMGTPE"
	return fmt.Sprintf("%.1f%c", float64(size)/float64(div), units[exp])
}
