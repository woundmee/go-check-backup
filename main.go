package main

import (
	"flag"
)

func main() {

	var (
		help     = flag.Bool("h", false, "Показать справку")
		database = flag.String("d", "", "Проверка бэкапов БД")
		os       = flag.String("o", "", "Проверка бэкапов ОС")
		version  = flag.Bool("v", false, "Версия скрипта")
	)

	flag.Parse()

	switch {
	case *help:
		getHelper()
	case *database != "":
		checkDB(*database)
	case *os != "":
		checkOS(*os, "os_backup_state.txt")
	case *version:
		getVersion()
	}
}
