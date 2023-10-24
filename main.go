package mylogger

import "log"

func LogInfo(message string) {
	log.Println("INFO - %v", message)
}

func LogWarning(message string) {
	log.Println("WARNING - %v", message)
}

func LogError(message string) {
	log.Println("ERROR - %v", message)
}
