package loghub

type (
	Logger interface {
		Debug(msg string)
		Info(msg string)
		Warn(err error)
		Error(err error)
	}

	logger struct {
		loggers []Logger
	}
)

func New(loggers ...Logger) Logger {
	return &logger{
		loggers: loggers,
	}
}

func (l logger) Debug(msg string) {
	for _, log := range l.loggers {
		log.Debug(msg)
	}
}

func (l logger) Info(msg string) {
	for _, log := range l.loggers {
		log.Info(msg)
	}
}

func (l logger) Warn(err error) {
	for _, log := range l.loggers {
		log.Warn(err)
	}
}

func (l logger) Error(err error) {
	for _, log := range l.loggers {
		log.Error(err)
	}
}
