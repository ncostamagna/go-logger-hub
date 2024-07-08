package loghub

type (
	Logger interface {
		Debug(msg string)
		Info(msg string)
		Warn(err error)
		Error(err error)
		GetTrace() int
	}

	logger struct {
		debug []func(msg string)
		info  []func(msg string)
		warn  []func(err error)
		error []func(err error)
	}
)

func New(loggers ...Logger) Logger {

	l := &logger{}
	for _, log := range loggers {
		if log.GetTrace()&DEBUG != 0 {
			l.debug = append(l.debug, log.Debug)
		}

		if log.GetTrace()&INFO != 0 {
			l.info = append(l.info, log.Info)
		}

		if log.GetTrace()&WARNING != 0 {
			l.warn = append(l.warn, log.Warn)
		}

		if log.GetTrace()&ERROR != 0 {
			l.error = append(l.error, log.Error)
		}
	}

	return l
}

func (l logger) Debug(msg string) {
	for _, log := range l.debug {
		log(msg)
	}
}

func (l logger) Info(msg string) {
	for _, log := range l.info {
		log(msg)
	}
}

func (l logger) Warn(err error) {
	for _, log := range l.warn {
		log(err)
	}
}

func (l logger) Error(err error) {
	for _, log := range l.error {
		log(err)
	}
}

func (l logger) GetTrace() int {
	return 0
}
