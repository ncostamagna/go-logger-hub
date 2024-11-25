package main

import (
	"errors"
	"github.com/ncostamagna/go-logger-hub/loghub"
	"log"
	"os"
)

func main() {

	// Custom Logger
	customLogger := loghub.CustomLogger{
		Debug: func(msg string) {
			log.Println("DEBUG custom: ", msg)
		},
		Info: func(msg string) {
			log.Println("INFO custom: ", msg)
		},
		Warn: func(err error) {
			log.Println("WARN custom: ", err)
		},
		Error: func(err error) {
			log.Println("ERROR custom: ", err)
		},
	}


	logger := loghub.New(
		loghub.NewNativeLogger(log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile),15),	
		loghub.NewCustomLogger(customLogger, 15),
	)

	logger.Debug("my message Debug")
	logger.Info("my message Info")
	logger.Warn(errors.New("my warn"))
	logger.Error(errors.New("my error"))

}
