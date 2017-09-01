package main

import (
  "github.com/talltom/loggr/loggr"
)

func main(){

  // Logger
  var log = loggr.Log {
  		Verbose: false,
  		// Writer: w, -> forgot to set valid writer object!
  }
  err := log.Info("Now writing to file")
  if err != nil {
    if err == loggr.ErrInvalidWriter {
      // Forgot to set writer
    } else {
      // Some other error occurred
    }
  }
}
