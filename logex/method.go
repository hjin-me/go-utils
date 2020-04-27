package logex

import (
	"github.com/sirupsen/logrus"
)

func Tracef(format string, args ...interface{}) {
	defaultLogger.Tracef(format, args...)
}

func Debugf(format string, args ...interface{}) {
	defaultLogger.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	defaultLogger.Infof(format, args...)
}

func Printf(format string, args ...interface{}) {
	defaultLogger.Printf(format, args...)
}

func Warnf(format string, args ...interface{}) {
	defaultLogger.Warnf(format, args...)
}

func Warningf(format string, args ...interface{}) {
	defaultLogger.Warningf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	defaultLogger.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	defaultLogger.Fatalf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	defaultLogger.Panicf(format, args...)
}

func Trace(args ...interface{}) {
	defaultLogger.Trace(args...)
}

func Debug(args ...interface{}) {
	defaultLogger.Debug(args...)
}

func Info(args ...interface{}) {
	defaultLogger.Info(args...)
}

func Warn(args ...interface{}) {
	defaultLogger.Warn(args...)
}

func Warning(args ...interface{}) {
	defaultLogger.Warning(args...)
}

func Error(args ...interface{}) {
	defaultLogger.Error(args...)
}

func Fatal(args ...interface{}) {
	defaultLogger.Fatal(args...)
}

func Panic(args ...interface{}) {
	defaultLogger.Panic(args...)
}

func Traceln(args ...interface{}) {
	defaultLogger.Traceln(args...)
}

func Debugln(args ...interface{}) {
	defaultLogger.Debugln(args...)
}

func Infoln(args ...interface{}) {
	defaultLogger.Infoln(args...)
}

func Warnln(args ...interface{}) {
	defaultLogger.Warnln(args...)
}

func Warningln(args ...interface{}) {
	defaultLogger.Warningln(args...)
}

func Errorln(args ...interface{}) {
	defaultLogger.Errorln(args...)
}

func Fatalln(args ...interface{}) {
	defaultLogger.Fatalln(args...)
}

func Panicln(args ...interface{}) {
	defaultLogger.Panicln(args...)
}

func WithFields(fields logrus.Fields) *logrus.Entry {
	return defaultLogger.WithFields(fields)
}
func WithField(key string, value interface{}) *logrus.Entry {
	return defaultLogger.WithField(key, value)
}

func TraceField(traceId, module string) logrus.Fields {
	return logrus.Fields{
		"trace_id": traceId,
		"module":   module,
	}
}
