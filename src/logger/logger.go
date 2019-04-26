package logger

import (
	"context"

	ctx "github.com/agustin-sarasua/user-system/src/context"
	"github.com/sirupsen/logrus"

	"io"
	"os"
	"strings"
)

const (
	defaultLoggerLevel = "INFO"
)

// Logger is an instance of Logger.
var Logger *logrus.Logger

func init() {
	Logger = logrus.New()
	Logger.Out = os.Stdout

	configLoggerLevel := os.Getenv("LOGGER_LEVEL")
	if configLoggerLevel == "" {
		configLoggerLevel = defaultLoggerLevel
	}

	level, error := logrus.ParseLevel(configLoggerLevel)
	if error != nil {
		panic(error)
	}

	Logger.Level = level
	Logger.Formatter = &MercuryFormatter{DisableColors: true}
}

func addFields(tags ...string) logrus.Fields {
	fields := make(logrus.Fields)

	for _, value := range tags {
		values := strings.Split(value, ":")

		fields[strings.TrimSpace(values[0])] = strings.TrimSpace(values[1])
	}

	return fields
}

// Info logs a message and the related tags, with an INFO level.
func Info(msg string, c context.Context, tags ...string) {
	if c == nil {
		Logger.WithFields(addFields(tags...)).Info(msg)
	} else {
		Logger.WithFields(addFields(tags...)).WithField("RequestID", ctx.UUID(c)).Info(msg)
	}
}

// Debug logs a message and the related tags, with an DEBUG level.
func Debug(msg string, tags ...string) {
	Logger.WithFields(addFields(tags...)).Debug(msg)
}

// Error logs a message and the related tags, with an ERROR level.
func Error(msg string, err error, tags ...string) {
	Logger.WithFields(addFields(tags...)).Errorf("%s - ERROR: %v", msg, err)
}

// GetOut returns the output writer.
func GetOut() io.Writer {
	return Logger.Out
}

// SetLevel changes the log level at runtime.
func SetLevel(lvl string) {
	level, err := logrus.ParseLevel(lvl)
	if err == nil {
		Logger.SetLevel(level)
	}
}
