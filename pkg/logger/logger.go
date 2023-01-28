package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetReportCaller(true)
}

// Default logger
var Default = Group("default")

// Group sets a specific group of traces
func Group(group string) *logrus.Entry {
	h, _ := os.Hostname()

	return logrus.WithFields(logrus.Fields{
		"group":    group,
		"hostname": h,
	})
}
