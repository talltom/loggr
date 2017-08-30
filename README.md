Loggr.go
========

A minimalist logger that writes to stdout

### Key features
- verbosity controlled by different log states (e.g. "debug", "info", "error")
- timestamps UTC printed as ISO 8601 string

### Install

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
> "Info statements always printed"
> "Error statements always printed"
> "Debug statement optional, requires verbose mode"
