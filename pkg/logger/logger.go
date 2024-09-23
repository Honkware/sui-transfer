package logger

import (
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func init() {
	log = logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:    false, // Disable timestamps if you don't need them
		DisableTimestamp: false, // Enable timestamps in a clean format
		DisableColors:    false, // Enable colors for better readability
		ForceColors:      true,  // Force colors even if output is not a TTY
		QuoteEmptyFields: true,  // Quote empty fields for better clarity
	})
	log.SetLevel(logrus.InfoLevel)
}

// Exported function to get the global logger
func GetLogger() *logrus.Logger {
	return log
}
