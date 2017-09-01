Loggr.go
========

A minimalist logger that writes sensible messages

[![Build Status](https://travis-ci.org/talltom/loggr.svg?branch=master)](https://travis-ci.org/talltom/loggr)

[![Coverage Status](https://coveralls.io/repos/github/talltom/loggr/badge.svg?branch=master)](https://coveralls.io/github/talltom/loggr?branch=master)

### Key features
- messages tagged by type (info, error, debug)
- optional verbose mode to toggle debug statements
- timestamps fixed to UTC and printed as ISO 8601 string
- extensible, accepts io.Writer object for output

### Design
Created so that Go programs could follow the existing logging structure (standardised timestamps) used in other related projects, and enable toggling of debug statements at run time to quickly adjust verbosity of output logs. Inspired somewhat by https://12factor.net/logs.

### Install
```go
go get github.com/talltom/loggr
```

### Use
**Example 1. Logging to Stdout**
```go
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
```

Result
```sh
2017-08-30T15:12:03Z - info: Info statement
2017-08-30T15:12:03Z - error: Error statement
2017-08-30T15:12:03Z - debug: Debug statement optional, requires verbose mode
```

**Example 2. Logging to File**
```go
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
```

**Example 3. Error handling**
Errors in loggr are returned to the user. Available error types are defined in loggr.go.
In this example no Writer object is set so the log cannot be written.
```go
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
```

### License
GNU GPLv3, see LICENSE.txt
