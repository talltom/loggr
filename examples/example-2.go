package main

import (
  "bufio"
  "github.com/talltom/loggr/loggr"
  "os"
)

func main(){

  // File
	f, err := os.Create("/tmp/logger.log")
	if err != nil {
		panic(err)
	}
	defer f.Close()

  // Writer
	w:= bufio.NewWriter(f)
	defer w.Flush()

  // Logger
  var log = loggr.Log {
  		Verbose: false,
  		Writer: w,
  }
  log.Info("Now writing to file")
  log.Debug("This will not be written as logger not in verbose mode")
}
