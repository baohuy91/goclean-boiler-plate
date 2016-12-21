package infrastructure

import "github.com/Sirupsen/logrus"

func NewLogger() Logger {
	return &loggerImpl{}
}

type Logger interface {
	Printf(format string, args ...interface{})
	LogWithFields(fields map[string]interface{}, err string)
}

type loggerImpl struct {
}

func (l *loggerImpl) Printf(format string, args ...interface{}) {
	logrus.Printf(format, args...)
}
func (l *loggerImpl) LogWithFields(fields map[string]interface{}, message string) {
	logrus.WithFields(fields).Println(message)
}
