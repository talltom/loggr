package main

import (
  "github.com/talltom/loggr/loggr"
  "os"
)

func main(){
  // create structure with verbose debug printing enabled
  log := loggr.Log {
    Verbose: true,
    Writer: os.Stdout,
  }

  log.Info("Info statement")
  log.Error("Error statement")
  log.Debug("Debug statement optional, requires verbose mode")
}
