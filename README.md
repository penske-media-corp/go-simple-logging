# go-simple-logging
> A set of simple logging utilities written in go that are easily configurable and geared towards simple use cases.

## Example

```go

	logger := logging.NewLogger()
	logging.LogLevel = logging.LOG_LEVEL_DEBUG
	logging.Debug("foo")
	logging.Debugf("debug: %s", "foo")
	logging.Error(1)
	logging.Errorf("error: %d", 1)
	logging.Info("hello")
	logging.Infof("info: %s", "hello")
	logging.Warn(true)
	logging.Warnf("warn: %t", true)
```

## Changing log destination

`Logger` has the following properties.  You can assign any `os.File` to redirect log output to that location.

* `DebugFile` (defaults to `os.Stderr`)
* `ErrorFile` (defaults to `os.Stderr`)
* `InfoFile` (defaults to `os.Stdout`)
* `WarnFile` (defaults to `os.Stderr`)

## Changing the log level

`Logger` has the `LogLevel` property which is an `int`.  The following convenient constants are exported:

* `LOG_LEVEL_NONE` turns off all logging.
* `LOG_LEVEL_ERROR` allows Error logs.
* `LOG_LEVEL_WARN` allows Error and Warn logs.
* `LOG_LEVEL_INFO` allows Error, Warn, and Info logs.
* `LOG_LEVEL_DEBUG` allows Error, Warn, Info, and Debug logs.
