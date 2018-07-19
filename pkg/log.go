package pkg

import (
	"errors"
	"os"

	"github.com/sirupsen/logrus"
)

var (
	AppLog *logrus.Logger
)

// InitLog use for initial log module
func InitLog() error {

	var err error

	AppLog = logrus.New()

	AppLog.Formatter = &logrus.TextFormatter{
		TimestampFormat: "2006/01/02 - 15:04:05",
		ForceColors:     true,
		FullTimestamp:   true,
	}

	// set logger
	if err := SetLogLevel(AppLog, Conf.Log.Level); err != nil {
		return errors.New("Set log level error: " + err.Error())
	}

	if err = SetLogOut(AppLog, Conf.Log.File); err != nil {
		return errors.New("Set log path error: " + err.Error())
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
		f, err := os.OpenFile(outString, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)

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
