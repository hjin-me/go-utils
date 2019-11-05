package logex

import (
	"bytes"
	"context"
	"errors"
	"io"
	"log"
	"strings"
	"testing"
)

func TestOutput(t *testing.T) {
	SetLogLevel(LDebug)
	var bf bytes.Buffer
	out := log.New(&bf, "", log.Llongfile)
	SetOutput(out, out)

	var s string
	Debug("123", "abc")
	s = bf.String()
	if strings.Index(s, "DBUG") == -1 {
		t.Error("Not Output Level", s)
	}
	if strings.Index(s, "123") == -1 {
		t.Error("Not Output 123", s)
	}
	bf.Reset()
	t.Log(bf.String())

	// Test Level
	SetLogLevel(LError)
	Debug("456", "xyz")
	s = bf.String()
	if strings.Index(s, "DBUG") > -1 {
		t.Error("Not Output Level", s)
	}
	if strings.Index(s, "123") > -1 {
		t.Error("Not Output 123", s)
	}
	bf = bytes.Buffer{}
}

type FatalWriter struct {
	io.Writer
}

func (f *FatalWriter) Write(p []byte) (n int, err error) {
	err = errors.New(string(p))
	return
}

func TestSetCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	SetLogLevel(LDebug)

	var fatalWriter = FatalWriter{}
	out := log.New(&fatalWriter, "", log.Llongfile)
	SetOutput(out, out)
	SetCancel(cancel)

	go func() {
		Debug("debug")
		Info("info")
	}()
	<-ctx.Done()
	if ctx.Err() != context.Canceled {
		t.Error(ctx.Err())
	}
	t.Log("success")
}
