package simplylog

import (
	"io"
	"os"
)

// Logger is the core component of simplylog.
type Logger struct {
	Out io.Writer
	// Format Format
	Verbose bool	
}


// New will create a new logger with the default settings.
func New() *Logger {
	return &Logger{
		Out: os.Stderr,
		Verbose: false,
	}
}