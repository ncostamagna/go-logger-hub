package loghub_test

import (
	"errors"
	"log"
	"os"
	"testing"

	"github.com/ncostamagna/go-logger-hub/loghub"
)

func BenchmarkLoghub_Logg(b *testing.B) {
	logger := loghub.New(loghub.NewNativeLogger(
		log.New(os.Stdin, "", log.LstdFlags|log.Lshortfile),
		loghub.ERROR+loghub.WARNING,
	))

	msg := "my message"
	err := errors.New("my error")
	for i := 0; i < b.N; i++ {
		logger.Debug(msg)
		logger.Info(msg)
		logger.Warn(err)
		logger.Error(err)
	}
}
