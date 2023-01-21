# go-log

A simple logging library in go.

### Example Usage

```go
package main

import (
	"errors"

	"github.com/brandonc/go-log"
)

func main() {
	// Reads the LOG environment variable to determine log level
	logger := log.NewLoggerFromEnv()

	// This version would log errors to stderr
	// logger := log.DefaultLogger

	logger.Trace("This is a trace message")
	logger.Debug("This is a debug message")
	logger.Info("This is an info message")
	logger.Warn("This is a warning")
	logger.Errorf("This is an error: %v", errors.New("some error"))
}
```

### Example Output:

```
$ go run main.go
2023-01-20T06:39:30.116-07:00 [ERROR] This is an error: some error
```

```
$ LOG=debug go run main.go
2023-01-20T06:41:06.112-07:00 [DEBUG] This is a debug message
2023-01-20T06:41:06.112-07:00 [INFO]  This is an info message
2023-01-20T06:41:06.112-07:00 [WARN]  This is a warning
2023-01-20T06:41:06.112-07:00 [ERROR] This is an error: some error
```
