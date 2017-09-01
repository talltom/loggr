[![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](http://www.gnu.org/licenses/gpl-3.0)
[![Build Status](https://travis-ci.org/talltom/loggr.svg?branch=master)](https://travis-ci.org/talltom/loggr)
[![Coverage Status](https://coveralls.io/repos/github/talltom/loggr/badge.svg?branch=master)](https://coveralls.io/github/talltom/loggr?branch=master)
[![Documentation](https://img.shields.io/badge/Documentation-GoDoc-green.svg)](https://godoc.org/github.com/talltom/loggr/loggr)


Loggr.go
========

A minimalist logger that writes sensible messages

Latest release: [v0.0.1](https://github.com/talltom/loggr/releases/tag/v0.0.1)

Go documentation at [godoc.org](https://godoc.org/github.com/talltom/loggr/loggr)

### Key features
- messages tagged by type (info, error, debug)
- optional verbose mode to toggle debug statements
- timestamps fixed to UTC and printed as ISO 8601 string
- extensible, accepts io.Writer object for output

### Design
Created so that Go programs could follow the existing log structure of [urbanriskmap](https://github.com/urbanriskmap) projects, and enable toggling of debug statements at run time to quickly adjust verbosity of output logs. Inspired somewhat by [12factor.net/logs](https://12factor.net/logs).

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

Log destinations (e.g. files) are managed by the user.
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
