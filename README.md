# SimplyLog

A simple, lightweight logging package that was inspired by an old article written by Dave Cheney on [Golang logging](https://dave.cheney.net/2015/11/05/lets-talk-about-logging). To summarise, a lot of logging packges offer too much functionality, a bunch of the 'logging levels' are left unused. This leaves us with `Info` and `Debug`:

* **Info** is utilised by the user of your program, it provides *information* about things which are useful to them.
* **Debug** is put in place to aid another developer/support engineer who is required to troubleshoot the software that has been deployed or to aid yourself in the actual development process - these should be used frequently throughout a program.


This package is an attempt to remedy this, whilst also being a great exercise in implementing logging for myself.

## Usage

Firstly, you should install the package using `go get`

    go get -u github.com/jdockerty/simplylog

A simple example of how this package works can be seen below
```go
package main

import (
    "github.com/jdockerty/simplylog"
)

func main() {
	logger := simplylog.New()	

	logger.Info("My informational message!")

	// The package defaults to only printing out 'Info'.
	logger.Debug("This will not be printed yet")

	// Setting verbose to be true enables the printing of 'Debug' messages
	logger.Verbose = true

	logger.Debug("This will be printed, verbose is true!")
}
```
This works much in the same way you would expect, although it gets straight to the point without any complex setup. If you wish to see `Debug` output then set the `Verbose` field to true, otherwise you do not have to change it.

A more configured example, would be something along the lines of

```go
package main

import (
	"github.com/jdockerty/simplylog"
	"os"
)

func main() {

	logger := simplylog.New()

	logger.Verbose = true
	logger.Format.SetType("JSON")
	logger.Format.SetTimestamp("2006/1/2 15:04")

	f, err := os.OpenFile("program-output.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		// Handle error
	}

	logger.SetOutput(f)

	logger.Info("My informational message!")

	logger.Debug("Debug output for my development process.")

}
```

This uses an output file, customised timestamp format, and setting the format type to `JSON`, rather than plain text output that you would be accustom to in the terminal. Setting the output to `JSON` is more compatible when another program is required to ingest the logs. When using a custom timestamp format, remember that the format should be in the specified Go reference time, otherwise you may see unexpected results.