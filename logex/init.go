package logex

import (
	"os"
	"reflect"
	"time"

	"github.com/sirupsen/logrus"
)

var defaultLogger *logrus.Entry

// prod will output json
func Init(fields logrus.Fields, prod bool) {
	logIns := logrus.New()
	logIns.SetOutput(os.Stdout)
	if prod {
		logIns.SetFormatter(&logrus.JSONFormatter{TimestampFormat: time.RFC3339Nano})
	}
	for s, i := range fields {
		if reflect.ValueOf(i).IsZero() {
			delete(fields, s)
		}
	}
	defaultLogger = logIns.WithFields(fields)
}
