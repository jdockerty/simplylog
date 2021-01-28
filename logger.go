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

// SetOutput allows the user to specify a specific output, such as a file.
// This defaults to stderr when using simplylog calling this.
func (l *Logger) SetOutput(output io.Writer) {
	l.Out = output
}