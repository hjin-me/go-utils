package logex

import (
	"os"
	"reflect"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

var defaultLogger *logrus.Entry

func init() {
	logIns := logrus.New()
	defaultLogger = logIns.WithField("_info_", "logex.Init() first")
}

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
	mutex.Unlock()
	p = true
}

var mutex sync.Mutex
var p = false

func init() {
	mutex.Lock()
}

func Ensure() *logrus.Entry {
	if !p {
		mutex.Lock()
		defer mutex.Unlock()
	}
	return defaultLogger
}
