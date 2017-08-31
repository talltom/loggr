Loggr.go
========

A minimalist logger that writes sensible messages

[![Build Status](https://travis-ci.org/talltom/loggr.svg?branch=master)](https://travis-ci.org/talltom/loggr)

[![Coverage Status](https://coveralls.io/repos/github/talltom/loggr/badge.svg?branch=master)](https://coveralls.io/github/talltom/loggr?branch=master)

### Key features
- optional verbose mode to print debug statements
- timestamps in UTC printed as ISO 8601 string
- extensible, accepts io.Writer object for output

### Install
- go get github.com/talltom/loggr

### Use
**Example 1. Logging to Stdout**
```go
import (
  "loggr"
  "os"
)

func main(){
  // create structure with verbose debug printing enabled
  log := loggr.Log {
    Verbose = true,
    Writer = os.Stdout,
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
import (
  "loggr"
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
	log.Writer = w
	defer w.Flush()

  // Logger
  var log = loggr.Log {
  		Verbose: true,
  		Writer: w,
  }
  log.Info("Now writing to file")
}
```

### License
- GNU GPLv3
