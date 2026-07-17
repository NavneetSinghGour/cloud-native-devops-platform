package logger

import (
	"log"
	"os"
)

var Logger = log.New(
	os.Stdout,
	"",
	log.Ldate|log.Ltime,
)

func Info(message string) {
	Logger.Printf("[INFO] %s", message)
}

func Warn(message string) {
	Logger.Printf("[WARN] %s", message)
}

func Error(message string) {
	Logger.Printf("[ERROR] %s", message)
}

func Fatal(message string) {
	Logger.Fatalf("[FATAL] %s", message)
}
