package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type dbBackupData struct {
	Day   int8 `json:"dayBakFile"`
	Month int8 `json:"monthBakFile"`
	Year  int8 `json:"yearBakFile"`
}

type osBackupData struct {
	WeekNameBackupOS     string `json:"weekNameBackupOS"`
	AvailabilityBackupOS string `json:"availabilityBackupOS"`
}

// ================================= //
// ==== CHECK DATABASE BACKUPS ==== //
// ================================= //

// path - путь до каталога с подкаталогами бэкапов Day/, Month/ и Year/
func checkDB(path string) string {

	// Day		: Pyramid2_backup_Day_2025_09_23.bak
	// Month	: Pyramid2_backup_Month_2025_09_01.bak
	// Year  	: Pyramid2_backup_Day_2025_01_01.bak

	backupTypes := [3]string{"Day", "Month", "Year"}

	d := checkDBRelevance(path+backupTypes[0], backupTypes[0])
	m := checkDBRelevance(path+backupTypes[1], backupTypes[1])
	y := checkDBRelevance(path+backupTypes[2], backupTypes[2])

	backupData := dbBackupData{
		Day:   d,
		Month: m,
		Year:  y,
	}

	jsonData, err := json.Marshal(backupData)
	if err != nil {
		log.Println("Error! ", err)
	}

	fmt.Println(string(jsonData))

	return string(jsonData)

}

func checkDBRelevance(path string, backupType string) int8 {

	files, err := os.ReadDir(path)
	if err != nil {
		log.Println("Error! ", err)
	}

	for _, file := range files {

		//Pyramid2_backup_Day_2025_01_01.bak
		if strings.Split(file.Name(), "_")[0] == "Pyramid2" && strings.Split(file.Name(), "_")[1] == "backup" {

			fileInfo, err := file.Info()

			if err != nil {
				log.Println("Error! ", err)
			}

			modFileDate := fileInfo.ModTime().Local().Format("2006-01-02")
			currentDate := time.Now().Local().Format("2006-01-02")

			now := time.Now()

			currentMonthFirstDay := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location()).Format("2006-01-02")
			currentYearFirstDay := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, now.Location()).Format("2006-01-02")

			switch backupType {
			case "Day":
				if modFileDate == currentDate {
					return 1
				}
			case "Month":
				if modFileDate == currentMonthFirstDay {
					return 1
				}
			case "Year":
				if modFileDate == currentYearFirstDay {
					return 1
				}
			}
		}
	}

	return 0
}

// =========================== //
// ==== CHECK OS BACKUPS ==== //
// =========================== //

func checkOS(path, saveToFile string) {

	content := checkOSRelevnce(path, saveToFile)

	backupData := osBackupData{
		WeekNameBackupOS:     strings.Split(content, ":")[0],
		AvailabilityBackupOS: strings.Split(content, ":")[1],
	}

	jsonData, err := json.Marshal(backupData)
	if err != nil {
		log.Println("Error! ", err)
	}

	fmt.Println(string(jsonData))
}

func checkOSRelevnce(path string, dstFile string) string {
	// Проверяю создание бэкапа и записываю в файл
	if time.Now().Local().Weekday() == time.Saturday {
		files, err := os.ReadDir(path)
		if err != nil {
			log.Println("Error! ", err)
		}

		for _, file := range files {

			currentDate := time.Now().Format("2006-01-02")

			backupFileExist := strings.Contains(file.Name(), fmt.Sprintf("System Backup%s", currentDate))
			if backupFileExist {
				// если есть файл бэкапа в день бэкапа
				writeToFile(1, file.Name(), dstFile)

			} else {
				// что делать, если нет файла бэкапа в день бэкапа...
				writeToFile(0, file.Name(), dstFile)
			}
		}

	} else {
		return readFromFile(dstFile)
	}

	return readFromFile(dstFile)
}

// функция записи в файл
func writeToFile(state int, content, filename string) {
	openFile, err := os.Create(filename)
	if err != nil {
		log.Println("Error! ", err)
	}

	defer openFile.Close()

	writer := bufio.NewWriter(openFile)
	writer.WriteString(fmt.Sprintf("%d:%s", state, content))
	writer.Flush()
}

// функция чтения из файл и возвращение прочитанного контента
func readFromFile(filename string) string {
	openFile, err := os.Open(filename)
	if err != nil {
		log.Println("Error! ", err)
	}

	defer openFile.Close()

	reader := bufio.NewReader(openFile)

	stat, _ := openFile.Stat()

	buf := make([]byte, stat.Size())

	reader.Read(buf)
	return string(buf)
}
