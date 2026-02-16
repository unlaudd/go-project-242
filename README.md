### Hexlet tests and linter status:
[![Build & Test](https://github.com/unlaudd/go-project-242/actions/workflows/build.yml/badge.svg)](https://github.com/unlaudd/go-project-242/actions/workflows/build.yml)
[![Actions Status](https://github.com/unlaudd/go-project-242/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/unlaudd/go-project-242/actions)

CLI-утилита для подсчёта размера файлов и директорий.

### Аксинема
[Asciicast](https://asciinema.org/a/K9lO6dmHTogs1O3N)

## Сборка

```bash
go build -o bin/hexlet-path-size cmd/hexlet-path-size/main.go
```

## Использование

### Обычный вывод
```bash
./bin/hexlet-path-size path/to/file
```

### Человекочитаемый формат
```bash
./bin/hexlet-path-size -H path/to/directory
```

### С учётом скрытых файлов
```bash
./bin/hexlet-path-size -a path/to/directory
```

### Рекурсивный подсчёт
```bash
./bin/hexlet-path-size -r path/to/directory
```

### Все флаги вместе
```bash
./bin/hexlet-path-size -H -a -r path/to/directory
```

## Флаги

|    Флаг     | Короткая форма |                 Описание                |
|-------------|----------------|-----------------------------------------|
| --recursive |      -r        | recursive size of directories           |
| --human     |      -H        | human-readable sizes (auto-select unit) |
| --all       |      -a        | include hidden files and directories    |
| --help      |      -h        | show help                               |