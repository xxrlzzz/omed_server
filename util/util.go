package util

import (
	"log"
	"os"
)

func TryOpen(filePath string) *os.File {
	_, err := os.Stat(filePath)
	if os.IsPermission(err) {
		log.Printf("permission deny log file %s", filePath)
		return nil
	} else if os.IsNotExist(err) {
		//err := os.MkdirAll(filePath, os.ModePerm)
		//if err != nil {
		//	log.Printf("fail to make dir for log file %s", filePath)
		//	return nil
		//}
	}
	logFile, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Printf("fail to open log file %s", filePath)
	}
	return logFile
}
