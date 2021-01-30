# SimplyLog

A simple logging package that was inspired by an old article written by Dave Cheney on [Golang logging](https://dave.cheney.net/2015/11/05/lets-talk-about-logging). To summarise, a lot of logging packges offer too much functionality, a lot of the 'logging levels' are left unused. This leaves us with `Info` and `Debug`:

* **Info** is utilised by the user of your program, it provides *information* about things which are useful to them.
* **Debug** is put in place to aid another developer/support engineer who is required to troubleshoot the software that has been deployed or mainly for the actual development process - these should be used frequently throughout a program.

## Usage

