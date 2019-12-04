package logex

import (
	"bytes"
	"errors"
	"log"
	"strings"
	"testing"
)

func TestWrap(t *testing.T) {

	var bf bytes.Buffer
	out := log.New(&bf, "", log.Llongfile)
	defaultLogger.SetOutput(out, out)

	logger := Wrap("aaaa", "mmmmm")
	logger.Info("x", "y")
	{
		s := bf.String()
		if strings.Index(s, "INFO") == -1 {
			t.Error("Not Output Level", s)
		}
		if strings.Index(s, `"req_id":"aaaa"`) == -1 {
			t.Error("req_id is not aaaa", s)
		}
		if strings.Index(s, `"module":"mmmmm"`) == -1 {
			t.Error("module name is not mmmmm", s)
		}
		if strings.Index(s, `"info":"[\"x\",\"y\"]"`) == -1 {
			t.Error("info is not [x,y]", s)
		}
		if strings.Index(s, "[default]") == -1 {
			t.Error("lost name")
		}
		if strings.Index(s, "\033[0;36m") == -1 {
			t.Error("lost color", s)
		}
		bf.Reset()
	}

	logger.Warning("6")
	{
		s := bf.String()
		if strings.Index(s, "WARN") == -1 {
			t.Error("Not Output Level", s)
		}
		if strings.Index(s, `"req_id":"aaaa"`) == -1 {
			t.Error("req_id is not aaaa", s)
		}
		if strings.Index(s, `"module":"mmmmm"`) == -1 {
			t.Error("module name is not mmmmm", s)
		}
		if strings.Index(s, `"info":"6"`) == -1 {
			t.Error("info is not 6", s)
		}
		bf.Reset()
	}
	logger.Fatal()
	{
		s := bf.String()
		if strings.Index(s, "FTAL") == -1 {
			t.Error("Not Output Level", s)
		}
		if strings.Index(s, `"req_id":"aaaa"`) == -1 {
			t.Error("req_id is not aaaa", s)
		}
		if strings.Index(s, `"module":"mmmmm"`) == -1 {
			t.Error("module name is not mmmmm", s)
		}
		if strings.Index(s, `"info"`) != -1 {
			t.Error("info should not exist", s)
		}
		bf.Reset()
	}

	{
		// normal
		var s string
		logger.Debug(errors.New("some err"))
		s = bf.String()
		if strings.Index(s, "wrap_test.go") == -1 {
			t.Error("depth is not right", s)
		}
		if strings.Index(s, "DBUG") == -1 {
			t.Error("Not Output Level", s)
		}
		if strings.Index(s, "some err") == -1 {
			t.Error("Not Output err.Error()", s)
		}
		bf.Reset()
	}
	{
		// statistics
		var s string
		logger.Stats(errors.New("some err"))
		s = bf.String()
		if strings.Index(s, "wrap_test.go") == -1 {
			t.Error("depth is not right", s)
		}
		if strings.Index(s, "STAT") == -1 {
			t.Error("Not Output Level", s)
		}
		if strings.Index(s, "some err") == -1 {
			t.Error("Not Output err.Error()", s)
		}
		bf.Reset()
	}
}
