package utils

// LoggingLevel is custom define the edgemicro logging level
type LoggingLevel string

// ResponseMessage is a message for response
type ResponseMessage string

const (
	// InfoLevel for edgemicro log
	InfoLevel LoggingLevel = "info"
	// WarnLevel for edgemicro log
	WarnLevel LoggingLevel = "warn"
	// ErrorLevel for edgemicro log
	ErrorLevel LoggingLevel = "error"

	// ResSuccess message
	ResSuccess ResponseMessage = "success"
	// ResError message
	ResError ResponseMessage = "error"
)
