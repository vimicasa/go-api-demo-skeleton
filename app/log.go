package app

import (
	"errors"
	"os"

	"github.com/mattn/go-isatty"
	"github.com/sirupsen/logrus"
)

var (
	green   = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	white   = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	yellow  = string([]byte{27, 91, 57, 55, 59, 52, 51, 109})
	red     = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	blue    = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	magenta = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	cyan    = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	reset   = string([]byte{27, 91, 48, 109})
)

var IsTerm bool

func init() {
	IsTerm = isatty.IsTerminal(os.Stdout.Fd())
}

// InitLog use for initial log module
func InitLog() error {

	var err error

	// init logger
	LogAccess = logrus.New()
	LogError = logrus.New()

	if !IsTerm {
		LogAccess.SetFormatter(&logrus.JSONFormatter{})
		LogError.SetFormatter(&logrus.JSONFormatter{})
	} else {
		LogAccess.Formatter = &logrus.TextFormatter{
			TimestampFormat: "2006/01/02 - 15:04:05",
			FullTimestamp:   true,
		}

		LogError.Formatter = &logrus.TextFormatter{
			TimestampFormat: "2006/01/02 - 15:04:05",
			FullTimestamp:   true,
		}
	}

	// set logger
	if err = SetLogLevel(LogAccess, AppConf.Log.AccessLevel); err != nil {
		return errors.New("Set access log level error: " + err.Error())
	}

	if err = SetLogLevel(LogError, AppConf.Log.ErrorLevel); err != nil {
		return errors.New("Set error log level error: " + err.Error())
	}

	if err = SetLogOut(LogAccess, AppConf.Log.AccessLog); err != nil {
		return errors.New("Set access log path error: " + err.Error())
	}

	if err = SetLogOut(LogError, AppConf.Log.ErrorLog); err != nil {
		return errors.New("Set error log path error: " + err.Error())
	}

	return nil
}

// SetLogOut provide log stdout and stderr output
func SetLogOut(log *logrus.Logger, outString string) error {
	switch outString {
	case "stdout":
		log.Out = os.Stdout
	case "stderr":
		log.Out = os.Stderr
	default:
		f, err := os.OpenFile(outString, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600)

		if err != nil {
			return err
		}

		log.Out = f
	}

	return nil
}

// SetLogLevel is define log level what you want
// log level: panic, fatal, error, warn, info and debug
func SetLogLevel(log *logrus.Logger, levelString string) error {
	level, err := logrus.ParseLevel(levelString)

	if err != nil {
		return err
	}

	log.Level = level

	return nil
}
