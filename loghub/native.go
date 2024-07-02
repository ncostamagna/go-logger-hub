package loghub

import (
	"fmt"
	"log"
	"os"
)

type (
	NativeLogger interface {
		Printf(format string, v ...any)
		Output(calldepth int, s string) error
	}

	nativeLogger struct {
		logger NativeLogger
		trace  int
	}
)

func NewNativeLogger(nativeLog NativeLogger, trace int) Logger {
	if nativeLog == nil {
		nativeLog = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	}
	return &nativeLogger{
		logger: nativeLog,
		trace:  trace,
	}
}

func (n nativeLogger) Debug(msg string) {
	if n.trace&DEBUG != 0 {
		n.logger.Output(3, fmt.Sprintf("DEBUG: %s\n", msg))
	}
}

func (n nativeLogger) Info(msg string) {
	if n.trace&INFO != 0 {
		n.logger.Output(3, fmt.Sprintf("INFO: %s\n", msg))
	}
}

func (n nativeLogger) Warn(err error) {
	if n.trace&WARNING != 0 {
		n.logger.Output(3, fmt.Sprintf("WARNING: %v\n", err))
	}
}

func (n nativeLogger) Error(err error) {
	if n.trace&ERROR != 0 {
		n.logger.Output(3, fmt.Sprintf("ERROR: %v\n", err))
	}
}
