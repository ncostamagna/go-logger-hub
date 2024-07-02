package main

import (
	"errors"
	"github.com/ncostamagna/go-logger-hub/loghub"
	"log"
	"os"
)

func main() {
	logger := loghub.New(loghub.NewNativeLogger(
		log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile),
		15,
	))

	logger.Debug("my message Debug")
	logger.Info("my message Info")
	logger.Warn(errors.New("my warn"))
	logger.Error(errors.New("my error"))

}
