package logex

import (
	"bytes"
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
		if strings.Index(s, `"info":["x","y"]`) == -1 {
			t.Error("info is not [x,y]", s)
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
}
