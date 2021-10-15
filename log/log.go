package log

// Debug is a constant that indicates the logger is in debug mode
const Debug = "Debug"

// Normal is a constant that indicates the logger is normally logging in the log files
const Normal = "Normal"

// None is a constant that indicates the logger isn't logging
const None = "None"

// DebugLogFile is a constant that indicates the log will be written on 'DebugLogFile'
const DebugLogFile = "DebugLogFile"

// ErrorLogFile is a constant that indicates the log will be written on 'ErrorLogFile'
const ErrorLogFile = "ErrorLogFile"

// LogContainer is a type that defines all the available logs
type LogContainer struct {
	DebugLogFile string // The 'DebugLogFile' is used for logging any transaction that is done in normal mode
	ErrorLogFile string // The 'ErrorLogFile' is used for logging any error occuring while performing transaction
}
