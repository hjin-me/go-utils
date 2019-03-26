package config

import (
	"testing"
)

func TestGet(t *testing.T) {
	Set("A", "B")
	var v string
	err := Get("A", &v)
	if err != nil {
		t.Error(err)
	}
	err = Get("B", &v)
	if err == nil {
		t.Error("should trigger error, no key")
	}
	if v != "" {
		t.Error("v should be zero")
	}
	var n int
	err = Get("A", &n)
	if err == nil {
		t.Error("should trigger can set")
	}
}
