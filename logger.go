package simplylog

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

// Logger is the core component of simplylog. This wraps the various concepts together.
type Logger struct {

	// Verbose is used to set whether the logging output should be verbose or not.
	// As such, setting verbose to true is akin to having the logging level set as 'Debug',
	// with false being 'Informational'. This follows Dave Cheney's dicussion on Golang logging.
	Verbose bool

	// Format provides configuration options for a few simple values that are used
	// within the package.
	Format Format

	// Out is the output
	Out io.Writer

	// Message is the string you wish to print, plus a timestamp which is generated for you.
	message msg

	// Mutex is mainly used when writing output to files.
	m sync.Mutex
}

// msg is used to encompass the string that is passed by the user. It combines a timestamp with a message
// and will be outputed in a different format dependent on what format type is selected.
type msg struct {
	Timestamp string `json:"timestamp"`
	Content   string `json:"msg"`
}

func getMsg(content, formatType, timestampFormat, level string) string {

	now := time.Now().Format(timestampFormat)

	if formatType == "json" {
		msg := &msg{
			Timestamp: now,
			Content:   content,
		}

		jsonMsg, err := json.Marshal(msg)
		if err != nil {
			fmt.Println("Unable to marshal as JSON, defaulting to text")
			getMsg(content, "text", timestampFormat, level)
		}
		return string(jsonMsg)
	}

	msg := fmt.Sprintf("[%s] %s, %s", level, now, content)
	return msg

}

// New will create a new logger with the default settings.
func New() *Logger {
	return &Logger{
		Format: Format{
			Timestamp: "15:04:05 02/01/2006", // Default Go reference time
			Type:      "text",
		},
		Out:     os.Stdout,
		Verbose: false,
	}
}

// SetOutput allows the user to specify a specific output, such as a file.
// This defaults to 'stdout' without using this function.
func (l *Logger) SetOutput(output io.Writer) {
	l.m.Lock()
	l.Out = output
	l.m.Unlock()
}

// Info is used to print out logs that are considered 'informational', aiding anyone
// who is using the program.
func (l *Logger) Info(args ...interface{}) {

	msgContent := fmt.Sprint(args...)
	msg := getMsg(msgContent, l.Format.Type, l.Format.Timestamp, "INFO")

	sendLog(l, msg)
}

// Infof is used to print a formatted string to the log, the same concept apply as per the Info function.
func (l *Logger) Infof(format string, args ...interface{}) {

	msgContent := fmt.Sprintf(format, args...)
	msg := getMsg(msgContent, l.Format.Type, l.Format.Timestamp, "INFO")

	sendLog(l, msg)
}

// Debug is used to print out logs that are used by other engineers, as you might expect
// for debugging purposes when trying to understand what has occurred in a system.
func (l *Logger) Debug(args ...interface{}) {

	if l.Verbose {
		msgContent := fmt.Sprint(args...)
		msg := getMsg(msgContent, l.Format.Type, l.Format.Timestamp, "DEBUG")

		sendLog(l, msg)
	}

}

// Debugf is used to print a formatted string to the log, the same concept apply as per the Debug function.
func (l *Logger) Debugf(format string, args ...interface{}) {

	if l.Verbose {
		msgContent := fmt.Sprintf(format, args...)
		msg := getMsg(msgContent, l.Format.Type, l.Format.Timestamp, "DEBUG")

		sendLog(l, msg)
	}

}

// Check where the output should be sent to, prior to writing it.
func sendLog(logger *Logger, msg string) {
	if logger.Out != os.Stdout {
		w := bufio.NewWriter(logger.Out)
		logger.m.Lock()

		w.Write([]byte(msg + "\n"))
		w.Flush()

		logger.m.Unlock()
	} else {
		fmt.Println(msg)
	}
}
