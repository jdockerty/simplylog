package simplylog

import (
	"io"
	"os"
	"sync"
)

// Logger is the core component of simplylog.
type Logger struct {

	// Format provides configuration options for a few simple values that are used
	// within the package.
	Format Format

	// Verbose is used to set whether the logging output should be set to 'Debug' or not.
	// As such, setting verbose to true is akin to having the logging level set as 'Debug',
	// with false being 'Informational'. This follows Dave Cheney's dicussion on Golang logging.
	Verbose bool

	// Out is the output
	Out io.Writer

	m sync.Mutex
}

// New will create a new logger with the default settings.
func New() *Logger {
	return &Logger{
		Format: Format{
			Timestamp: "15:04:05 02/01/2006", // Default Go reference time
			Type:      "text",
		},
		Out:     os.Stderr,
		Verbose: false,
	}
}

// SetOutput allows the user to specify a specific output, such as a file.
// This defaults to stderr when using simplylog calling this.
func (l *Logger) SetOutput(output io.Writer) {
	l.m.Lock()
	l.Out = output
	l.m.Unlock()
}
