# go-check-backup [![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?logo=go)](https://golang.org/doc/devel/release) 

**go-check-backup** — это утилита для мониторинга и проверки актуальности резервных копий, созданных SQL-серверами и Veeam Agent (Windows Server). Интегрируется с системами мониторинга (например, Zabbix) через JSON-вывод.

## ✨ Возможности

- **Проверка бэкапов БД** (SQL Server): анализирует наличие свежих `.bak` файлов в иерархии каталогов `Day/`, `Month/`, `Year/` по шаблону имени.
- **Проверка бэкапов ОС** (Veeam): определяет дату создания последнего полного бэкапа, записывает состояние в файл и выводит JSON.
- **Гибкий вывод**: JSON для систем мониторинга (Zabbix и др.) и текстовый файл с состоянием для ОС.
- **Простота использования**: единый бинарный файл без зависимостей.

## 📦 Установка

### Требования
- Go версии **1.25** или выше (только для сборки).

### Сборка из исходников
```sh
git clone https://github.com/woundmee/go-check-backup.git
cd go-check-backup
go build -o go-check-backup
```

## 🚀 Использование
Утилита имеет простой CLI-интерфейс. Для получения справки используйте флаг `-h`.

### ✔ **Проверка резервных копий баз данных (SQL Server)**
```sh
./go-check-backup -d /path/to/backup/root
```

Структура каталогов и имена файлов должны строго соответствовать:

```
/path/to/backup/root/
├── Day/
│   └── Pyramid2_backup_Day_2025_09_24.bak
├── Month/
│   └── Pyramid2_backup_Month_2025_09_01.bak
└── Year/
    └── Pyramid2_backup_Year_2025_01_01.bak
```

Шаблон имени: `{Prefix}_backup_{Type}_{YYYY}_{MM}_{DD}.bak`
- `{Prefix}` — по умолчанию Pyramid2
- `{Type}` — Day, Month или Year

Пример JSON-вывода:
```json
{
    "dayBakFile": 1,
    "monthBakFile": 1,
    "yearBakFile": 0
}
```
где, 1 - указание наличия и актуальности бекапа.
<br>

### ✔ **Проверка резервных копий ОС (Veeam Agent)**
```sh
./go-check-backup -o /path/to/veeam/backups
```
Утилита анализирует файлы с именем, содержащим дату в формате `System BackupYYYY-MM-DDThhmmss` (стандартное именование Veeam).
Результат сохраняется в файл `os_backup_state.txt` рядом с бинарным файлом в формате:

```sh
1:backupName
# Где 1 — бэкап актуален, 0 — нет.
```

Пример JSON-вывода
```json
{
    "weekNameBackupOS": "System Backup2026-02-21T010013.vib",
    "availabilityBackupOS": "1"
}
```

## 🤝 Как помочь проекту
- Сообщайте об ошибках через [Issues](https://github.com/woundmee/go-check-backup/issues).
- Предлагайте улучшения через Pull Requests.
