// Package loggr provides a minimalist logger that writes sensible messages
package loggr

import (
	"errors"
	"fmt"
	"io"
	"time"
)

// Define an error with the Writer object
var ErrInvalidWriter = errors.New("loggr: invalid writer")

// Logger construct
type Log struct {
	Verbose bool      // Toggle debug messages on/off
	Writer  io.Writer // Output
}

// Write function
func (l Log) write(level string, s string) (err error) {
	// Catch unassigned writer
	if l.Writer == nil {
		return ErrInvalidWriter
	} else {
		_, err := fmt.Fprintf(l.Writer, "%s - %s: %s\n", time.Now().UTC().Format(time.RFC3339), level, s)
		return err
	}
}

// Write info message to log
func (l Log) Info(s string) (err error) {
	// Always write info strings
	return l.write("info", s)
}

// Write error message to log
func (l Log) Error(s string) (err error) {
	// Always write error strings
	return l.write("error", s)
}

// Write debug message to log
func (l Log) Debug(s string) (err error) {
	// Only write if debug is true
	if l.Verbose == true {
		return l.write("debug", s)
	}
	return
}
