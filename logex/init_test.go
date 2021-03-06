package logex

import (
	"testing"
	"time"
)

func TestEnsure(t *testing.T) {
	t.Run("wait 3 second", func(t *testing.T) {
		ch := make(chan interface{})
		go func() {
			ch <- Ensure()
		}()

		select {
		case <-ch:
			t.Error("ensure should not trigger")
		case <-time.After(3 * time.Second):
			t.Log("success")
		}
	})
	t.Run("ensure first", func(t *testing.T) {
		ch := make(chan interface{})
		defer close(ch)
		go func() {
			ch <- Ensure()
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			Init(nil, false)
			Logger.Debug("debug log")
		}()

		select {
		case <-ch:
			t.Log("ensure trigger first")
		case <-time.After(2 * time.Second):
			t.Error("ensure not trigger")
		}
	})
	t.Run("debug log", func(t *testing.T) {
		Init(nil, false)
		Logger.Debug("some debug log")
	})
	t.Run("product log", func(t *testing.T) {
		Init(nil, true)
		Logger.Debug("some debug log")
	})
	t.Run("double init", func(t *testing.T) {
		Init(nil, false)
		Init(nil, false)
	})
}
