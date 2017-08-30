Loggr.go
========

A minimalist logger that writes to stdout

### Key features
- optional verbose mode to print debug statements
- timestamps in UTC printed as ISO 8601 string

### Install
- go get github.com/talltom/loggr

### Use
Example
```go
import "loggr"

// create structure with verbose debug printing enabled
log := loggr.Log {
  Verbose = true,
}

log.Info("Info statement")
log.Error("Error statement")
log.Debug("Debug statement optional, requires verbose mode")
```

Result
```sh
2017-08-30T15:12:03Z - info: Info statement
2017-08-30T15:12:03Z - error: Error statement
2017-08-30T15:12:03Z - debug: Debug statement optional, requires verbose mode
```
