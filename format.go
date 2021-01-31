package simplylog

import (
	"fmt"
	"strings"
)

// Format is used to provide configuration for different options.
type Format struct {
	Timestamp string
	Type      string
}

// SetType is used to determine whether the format is simply 'text' or 'json' for the log output type.
func (f *Format) SetType(formatType string) {

	chosenFormat := strings.ToLower(formatType)
	switch chosenFormat {

	case "text":
		f.Type = "text"
	case "json":
		f.Type = "json"

	default:
		fmt.Println("Invalid type used, defaulting to text.")
		f.Type = "text"
	}
}

// SetTimestamp will create the appropriate formatted time for future logs.
// If this function is not called, the default time layout is HH:MM:SS DD-MM-YYYY.
// Go uses a reference time of Mon Jan 2 15:04:05 -0700 MST 2006, so you should
// place your desired format using this specific reference time.
func (f *Format) SetTimestamp(timestampFormat string) {

	f.Timestamp = timestampFormat

}
