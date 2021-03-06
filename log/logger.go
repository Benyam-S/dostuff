package log

import (
	"fmt"
	"os"
	"sync"
	"time"
)

// Logger is a type that defines the logger
type Logger struct {
	mu   sync.Mutex
	Logs *LogContainer
	flag string // Can define that state of the logger wheather to log or not
}

// NewLogger is a function that returns a new logger
func NewLogger(logs *LogContainer, flag string) *Logger {
	logger := &Logger{Logs: logs}

	// Checking if the log file exist or not
	// If they don't exist it will create them
	_, err1 := os.Stat(logger.Logs.DebugLogFile)
	if err1 != nil {
		os.Create(logger.Logs.DebugLogFile)
	}

	_, err2 := os.Stat(logger.Logs.ErrorLogFile)
	if err2 != nil {
		os.Create(logger.Logs.ErrorLogFile)
	}

	logger.SetFlag(flag)
	return logger
}

// SetFlag is a method that set the logger's flag to given state
func (l *Logger) SetFlag(state string) {
	if state != Debug && state != Normal && state != None {
		state = None
	}

	l.flag = state
}

// Log is a method that will log the given statement to the selected log file
func (l *Logger) Log(stmt, logFile string) {

	// Checking the status of the logger
	if l.flag == None {
		return
	} else if l.flag == Debug {
		l.LogToParent(stmt)
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY, 0644)
	if err == nil {
		defer file.Close()

		stmt = fmt.Sprintf("[ %s ] %s", time.Now(), stmt)

		fmt.Fprintln(file, stmt)

	}
}

// LogToParent is a method that will log the given statement to the program starter
func (l *Logger) LogToParent(stmt string) {

	// Checking the status of the logger
	if l.flag == None {
		return
	}

	stmt = fmt.Sprintf("[ %s ] %s", time.Now(), stmt)
	fmt.Println(stmt)
}
