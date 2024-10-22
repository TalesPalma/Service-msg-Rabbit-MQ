package logsservice

import (
	"fmt"
	"log"
	"os"
	"time"
)

func SaveLog(msg string) {
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	date := time.Now().Format("2006-01-02 15:04:05")
	msgSave := fmt.Sprintf("%s - %s\n", date, msg)

	if _, err := file.Write([]byte(msgSave)); err != nil {
		log.Fatal(err)
	}

}
