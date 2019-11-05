// Custom Writer use SetOutput
// Custom Level use SetLogLevel
// handle log fatal use SetCancel
package logex

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"time"
)

type Level uint

const (
	LNone Level = iota
	LFatal
	LError
	LWarning
	LInfo
	LDebug
	LMax
)

func (l Level) String() string {
	switch l {
	case LFatal:
		return "FTAL"
	case LError:
		return "ERRO"
	case LWarning:
		return "WARN"
	case LInfo:
		return "INFO"
	case LDebug:
		return "DBUG"
	}
	return ""
}

var defaultLogger *loggerIns = &loggerIns{}

type Formatter func(v ...interface{}) ([]byte, error)

type loggerIns struct {
	color     bool
	name      string
	level     Level
	out       *log.Logger
	errOut    *log.Logger
	cancel    context.CancelFunc
	depth     int
	formatter Formatter
}

func (l *loggerIns) SetName(n string) {
	l.name = n
}
func (l *loggerIns) SetLogLevel(level Level) {
	l.level = level
}
func (l *loggerIns) SetOutput(out, err *log.Logger) {
	l.out = out
	l.errOut = err
}
func (l *loggerIns) SetCancel(fn context.CancelFunc) {
	l.cancel = fn
}
func (l *loggerIns) SetFormatter(f Formatter) {
	l.formatter = f
}
func (l *loggerIns) DepthIncrease(delta int) {
	l.depth = l.depth + 1
}
func (l *loggerIns) Clone() *loggerIns {
	return New(l.level, l.out, l.errOut)
}
func defaultFormatter(v ...interface{}) ([]byte, error) {
	switch len(v) {
	case 0:
		return nil, nil
	case 1:
		b, err := json.Marshal(v[0])
		return b, err
	default:
		b, err := json.Marshal(v)
		return b, err
	}
}
func colorFormatter(level Level, s string) string {
	// 开发环境懒得考虑性能
	r := regexp.MustCompile("(FTAL|ERRO|WARN|INFO|DBUG)")
	pre := ""
	switch level {
	case LFatal:
		pre = "\033[0;41m"
	case LError:
		pre = "\033[0;31m"
	case LWarning:
		pre = "\033[0;33m"
	case LInfo:
		pre = "\033[0;36m"
	case LDebug:
		pre = "\033[0;32m"
	}
	return r.ReplaceAllString(s, pre+"$1"+"\033[0m")
}
func (l *loggerIns) output(level Level, callDepth int, v ...interface{}) {
	if level > l.level {
		return
	}
	var err error
	t := time.Now().Format("2006-01-02T15:04:05.999-07:00")
	body, err := l.formatter(v...)
	msg := fmt.Sprintf("[%s][%s][%s]%s", l.name, t,
		level.String(),
		body)
	if l.color {
		msg = colorFormatter(level, msg)
	}
	switch level {
	case LFatal:
		fallthrough
	case LError:
		err = l.errOut.Output(callDepth, msg)
	case LWarning:
		fallthrough
	case LInfo:
		fallthrough
	case LDebug:
		err = l.out.Output(callDepth, msg)
	}
	if err != nil {
		fmt.Printf("[%s][%s][FATAL]%v", l.name, t, err)
		l.cancel()
	}
}
func (l *loggerIns) outputf(level Level, format string, v ...interface{}) {
	l.output(level, l.depth+1, fmt.Sprintf(format, v...))
}

// Fatalf is equivalent to Printf() for FATAL-level log.
func (l *loggerIns) Fatalf(format string, v ...interface{}) {
	l.outputf(LFatal, format, v...)
}

// Fatal is equivalent to Print() for FATAL-level log.
func (l *loggerIns) Fatal(v ...interface{}) {
	l.output(LFatal, l.depth, v...)
}

// Errorf is equivalent to Printf() for Error-level log.
func (l *loggerIns) Errorf(format string, v ...interface{}) {
	l.outputf(LError, format, v...)
}

// Error is equivalent to Print() for Error-level log.
func (l *loggerIns) Error(v ...interface{}) {
	l.output(LError, l.depth, v...)
}

// Warningf is equivalent to Printf() for WARNING-level log.
func (l *loggerIns) Warningf(format string, v ...interface{}) {
	l.outputf(LWarning, format, v...)
}

// Waring is equivalent to Print() for WARING-level log.
func (l *loggerIns) Warning(v ...interface{}) {
	l.output(LWarning, l.depth, v...)
}

// Infof is equivalent to Printf() for Info-level log.
func (l *loggerIns) Infof(format string, v ...interface{}) {
	l.outputf(LInfo, format, v...)
}

// Info is equivalent to Print() for Info-level log.
func (l *loggerIns) Info(v ...interface{}) {
	l.output(LInfo, l.depth, v...)
}

// Debugf is equivalent to Printf() for DEBUG-level log.
func (l *loggerIns) Debugf(format string, v ...interface{}) {
	l.outputf(LDebug, format, v...)
}

// Debug is equivalent to Print() for DEBUG-level log.
func (l *loggerIns) Debug(v ...interface{}) {
	l.output(LDebug, l.depth, v...)
}

func New(level Level, out, errOut *log.Logger) *loggerIns {
	logger := loggerIns{}
	logger.SetLogLevel(level)
	logger.SetOutput(out, errOut)
	logger.depth = 3
	logger.SetFormatter(defaultFormatter)
	return &logger
}

func init() {
	var flag int
	var level Level
	var color bool

	if os.Getenv("LOG_MODE") != "production" {
		color = true
		level = LDebug
		flag = log.Llongfile
		defer func() {
			Infof("Mode debug, color=[%v], logLevel=[%d]", true, LDebug)
		}()
	} else {
		flag = 0
		level = LInfo
	}
	loggerOut := log.New(os.Stdout, "", flag)
	loggerErr := log.New(os.Stderr, "", flag)
	defaultLogger = New(level, loggerOut, loggerErr)
	defaultLogger.color = color
	defaultLogger.SetOutput(loggerOut, loggerErr)
	defaultLogger.SetName("default")
	//defaultLogger.DepthIncrease(1)
	defaultLogger.SetCancel(func() {
		panic(errors.New("write log failed, but no context cancel"))
	})
}
func SetOutput(out, err *log.Logger) {
	defaultLogger.SetOutput(out, err)
}
func SetName(n string) {
	defaultLogger.SetName(n)
}

func SetLogLevel(level Level) {
	defaultLogger.SetLogLevel(level)
}

// When Log write failed, it will call cancel.
// context.Context should be Done. others goroutine should finish their jobs and exit safety.
func SetCancel(fn context.CancelFunc) {
	defaultLogger.SetCancel(fn)
}

// Fatalf is equivalent to Printf() for FATAL-level log.
func Fatalf(format string, v ...interface{}) {
	defaultLogger.Fatalf(format, v...)
}

// Fatal is equivalent to Print() for FATAL-level log.
func Fatal(v ...interface{}) {
	defaultLogger.Fatal(v...)
}

// Errorf is equivalent to Printf() for Error-level log.
func Errorf(format string, v ...interface{}) {
	defaultLogger.Errorf(format, v...)
}

// Error is equivalent to Print() for Error-level log.
func Error(v ...interface{}) {
	defaultLogger.Error(v...)
}

// Warningf is equivalent to Printf() for WARNING-level log.
func Warningf(format string, v ...interface{}) {
	defaultLogger.Warningf(format, v...)
}

// Waring is equivalent to Print() for WARING-level log.
func Warning(v ...interface{}) {
	defaultLogger.Warning(v...)
}

// Infof is equivalent to Printf() for Info-level log.
func Infof(format string, v ...interface{}) {
	defaultLogger.Infof(format, v...)
}

// Info is equivalent to Print() for Info-level log.
func Info(v ...interface{}) {
	defaultLogger.Info(v...)
}

// Debugf is equivalent to Printf() for DEBUG-level log.
func Debugf(format string, v ...interface{}) {
	defaultLogger.Debugf(format, v...)
}

// Debug is equivalent to Print() for DEBUG-level log.
func Debug(v ...interface{}) {
	defaultLogger.Debug(v...)
}
