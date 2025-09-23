package main

import "fmt"

func getVersion() {
	fmt.Println("\ndeveloper: woundmee")
	fmt.Println("startDate: 23.09.2025")
	fmt.Print("version: 0.2\n\n")
}

func getHelper() {

	title := `  
  ____  ___         ____ _   _ _____ ____ _  __     ____    _    ____ _  ___   _ ____  ____  
 / ___|/ _ \       / ___| | | | ____/ ___| |/ /    | __ )  / \  / ___| |/ / | | |  _ \/ ___| 
| |  _| | | |_____| |   | |_| |  _|| |   | ' /_____|  _ \ / _ \| |   | ' /| | | | |_) \___ \ 
| |_| | |_| |_____| |___|  _  | |__| |___| . \_____| |_) / ___ \ |___| . \| |_| |  __/ ___) |
 \____|\___/       \____|_| |_|_____\____|_|\_\    |____/_/   \_\____|_|\_\\___/|_|   |____/ 
 `

	fmt.Println(title)
	fmt.Print("\nИспользование: goCheckBackups [-h, -d, -o, -v]. Используй -h или --help.\n\n")
	fmt.Println("-h\t", "Показать справку")
	fmt.Println("-d\t", "Проверка бэкапов БД")
	fmt.Println("-o\t", "Проверка бэкапов ОС")
	fmt.Println("-v\t", "Версия скрипта")

	fmt.Print("\n")
}
