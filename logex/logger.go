package logex

type Logger interface {
	Fatalf(format string, v ...interface{})

	// Fatal is equivalent to Print() for FATAL-level log.
	Fatal(v ...interface{})

	// Errorf is equivalent to Printf() for Error-level log.
	Errorf(format string, v ...interface{})

	// Error is equivalent to Print() for Error-level log.
	Error(v ...interface{})

	// Warningf is equivalent to Printf() for WARNING-level log.
	Warningf(format string, v ...interface{})
	// Waring is equivalent to Print() for WARING-level log.
	Warning(v ...interface{})

	// Infof is equivalent to Printf() for Info-level log.
	Infof(format string, v ...interface{})

	// Info is equivalent to Print() for Info-level log.
	Info(v ...interface{})

	// Debugf is equivalent to Printf() for DEBUG-level log.
	Debugf(format string, v ...interface{})

	// Debug is equivalent to Print() for DEBUG-level log.
	Debug(v ...interface{})
}
