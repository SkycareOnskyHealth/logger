package logger

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	// DefautlLogBucket default directory to log
	DefautlLogBucket = "log"
	// DefautlLogLocation default location to log
	DefautlLogLocation = "./logs"
)

// Init config log
func Init(logLevel int, logLocation string, logBucket string) (*zerolog.Logger, *os.File, error) {
	var f *os.File
	if logBucket == "" {
		logBucket = DefautlLogBucket
	}
	if logLocation == "" {
		logLocation = DefautlLogLocation
	}
	// Init logger
	currentTime := time.Now()
	filePath := fmt.Sprintf("%s/%s-%s.log", logLocation, logBucket, currentTime.Format("02-01-2006"))
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("not exists")
			// path/to/whatever does *not* exist
			err = os.Mkdir(logLocation, 0755)
			if err != nil {
				return nil, f, err
			}

		} else {
			fmt.Println("else")
			// Schrodinger: file may or may not exist. See err for details.

			// Therefore, do *NOT* use !os.IsNotExist(err) to test for file existence
			return nil, nil, errors.New("file error")
		}
	}
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, f, err
	}

	logger, err := Config(logLevel, f)
	if err != nil {
		return nil, f, err
	}
	return logger, f, nil
}

// Config - custom time format for logger of empty string to use default
func Config(lvl int, file *os.File) (*zerolog.Logger, error) {
	var logLevel zerolog.Level
	//! File
	if file != nil {
		log.Logger = log.Output(file)
	}
	//!
	switch lvl {
	case -1:
		logLevel = zerolog.TraceLevel
	case 0:
		logLevel = zerolog.DebugLevel
	case 1:
		logLevel = zerolog.InfoLevel
	case 2:
		logLevel = zerolog.WarnLevel
	case 3:
		logLevel = zerolog.ErrorLevel
	case 4:
		logLevel = zerolog.FatalLevel
	case 5:
		logLevel = zerolog.PanicLevel
	default:
		logLevel = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(logLevel)
	log.Logger = log.With().Caller().Logger()

	return &log.Logger, nil
}
