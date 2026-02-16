package pkg

import (
	"fmt"
	"os"
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
	sizeStr := formatSize(size, human)

	return fmt.Sprintf("%s\t%s", sizeStr, path), nil
}

// formatSize форматирует размер в человекочитаемый вид
func formatSize(size int64, human bool) string {
	if !human {
		return fmt.Sprintf("%d", size)
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
