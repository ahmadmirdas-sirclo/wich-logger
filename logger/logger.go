package logger

import (
	"log"
)

func LogDebug(message string) {
	log.Printf("DEBUG - %v", message)
}

func LogInfo(message string) {
	log.Printf("INFO - %v", message)
}

func LogWarn(message string) {
	log.Printf("WARN - %v", message)
}

func LogError(message string) {
	log.Printf("ERROR - %v", message)
}
