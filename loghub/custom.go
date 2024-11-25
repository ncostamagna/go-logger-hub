package loghub

type (
	CustomLogger struct {
		Debug func(msg string)
		Info  func(msg string)
		Warn  func(err error)
		Error func(err error)
	}

	customLogger struct {
		logger CustomLogger
		trace  int
	}
)

func NewCustomLogger(cl CustomLogger, trace int) Logger {
	return &customLogger{
		logger: cl,
		trace:  trace,
	}
}

func (c customLogger) Debug(msg string) {
	c.logger.Debug(msg)
}

func (c customLogger) Info(msg string) {
	c.logger.Info(msg)
}

func (c customLogger) Warn(err error) {
	c.logger.Warn(err)
}

func (c customLogger) Error(err error) {
	c.logger.Error(err)
}

func (c customLogger) GetTrace() int {
	return c.trace
}
