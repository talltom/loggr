// Minimalist logger
package loggr

import (
  "fmt"
  "time"
  "io"
  )

// Logger construct with logging level
type Log struct {
  Verbose bool // Optionally print debug values
  Writer io.Writer // Log destination
}

// Write function
// TODO - testing - make this return a value?
func (l Log) write (level string, s string){
  fmt.Fprintf(l.Writer, "%s - %s: %s\n", time.Now().UTC().Format(time.RFC3339), level, s)
}

// Handle Log.Info(msg) statements
func (l Log) Info(s string) {
  // Always write info strings
  l.write("info", s)
}

// Handle Log.Error(msg) statements
func (l Log) Error(s string) {
  // Always write error strings
  l.write("error", s)
}

// Handle Log.Debug(msg) statements
func (l Log) Debug(s string) {
  // Only write if debug is true
  if (l.Verbose == true){
    l.write("debug", s)
  }
}
