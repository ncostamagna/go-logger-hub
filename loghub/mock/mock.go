package mock

type LoghubTransport struct {
	DebugMock func(msg string)
	InfoMock  func(msg string)
	WarnMock  func(err error)
	ErrorMock func(err error)	
	GetTraceMock func() int
}

func (t LoghubTransport) Debug(msg string) {
	if t.DebugMock == nil {
		return
	}
	t.DebugMock(msg)
}

func (t LoghubTransport) Info(msg string) {
	if t.InfoMock == nil {
		return
	}
	t.InfoMock(msg)
}

func (t LoghubTransport) Warn(err error) {
	if t.WarnMock == nil {
		return
	}
	t.WarnMock(err)
}

func (t LoghubTransport) Error(err error) {
	if t.ErrorMock == nil {
		return
	}
	t.ErrorMock(err)
}

func (t LoghubTransport) GetTrace() int {
	if t.GetTraceMock == nil {
		return 0
	}
	return t.GetTraceMock()
}
