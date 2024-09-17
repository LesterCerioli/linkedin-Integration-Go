package utils
import (
    "encoding/json"
    "os"
    "time"
)

type LogEntry struct {
    Timestamp time.Time `json:"timestamp"`
    Status    string    `json:"status"`
    Message   string    `json:"message"`
}

func LogEvent(status, message string) {
    logEntry := LogEntry{
        Timestamp: time.Now(),
        Status:    status,
        Message:   message,
    }
    logFile, _ := os.OpenFile("events.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    defer logFile.Close()
    logEntryJSON, _ := json.Marshal(logEntry)
    logFile.Write(logEntryJSON)
    logFile.Write([]byte("\n"))
}