// Package loggr provides a minimalist logger that writes sensible messages
package loggr

import (
	"errors"
	"fmt"
	"io"
	"time"
)

// Define an error
var ErrInvalidWriter = errors.New("loggr: invalid writer")

// Logger construct with logging level
type Log struct {
	Verbose bool      // Optionally print debug values
	Writer  io.Writer // Log destination
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

// Handle Log.Info(msg) statements
func (l Log) Info(s string) (err error) {
	// Always write info strings
	return l.write("info", s)
}

// Handle Log.Error(msg) statements
func (l Log) Error(s string) (err error) {
	// Always write error strings
	return l.write("error", s)
}

// Handle Log.Debug(msg) statements
func (l Log) Debug(s string) (err error) {
	// Only write if debug is true
	if l.Verbose == true {
		return l.write("debug", s)
	}
	return
}
