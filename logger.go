package simplylog

import (
	"io"
	"os"
)

// Logger is the core component of simplylog.
type Logger struct {
	Out io.Writer
	// Format Format

	// Verbose is used to set whether the logging output should be set to 'Debug' or not.
	// As such, setting verbose to true is akin to having the logging level set as 'Debug', 
	// with false being 'Informational'. This follows Dave Cheney's dicussion on Golang logging.
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