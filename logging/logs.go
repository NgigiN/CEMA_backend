// This module provides a simple reusable and consistent logging utility for the application.
package logging

import (
	"log"
	"os"
	"time"
)

var Logger *log.Logger

// Initialize sets up the logger to write to standard output with a specific format
func Initialize() {
	Logger = log.New(os.Stdout, "INFO:", log.Ldate|log.Ltime|log.Lshortfile)
}

// Info logs an informational message with a timestamp
func Info(message string) {
	Logger.SetPrefix("INFO: ")
	Logger.Println(time.Now().Format("2006-01-02 15:04:05") + " " + message)
}

// Error logs an error message with a timestamp
func Error(message string) {
	Logger.SetPrefix("ERROR: ")
	Logger.Println(time.Now().Format("2006-01-02 15:04:05") + " " + message)
}

// Warning logs a warning message with a timestamp
func Fatal(message string) {
	Logger.SetPrefix("FATAL: ")
	Logger.Fatalln(time.Now().Format("2006-01-02 15:04:05") + " " + message)
}
