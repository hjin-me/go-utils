package cerror

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHttpResponseSuccess(t *testing.T) {
	sCode := ErrCode{123000, 200}
	DefineSuccessCode(sCode)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Log("receive request")
		HttpResponseSuccess(w, nil)
	}))
	req, err := http.Post(ts.URL, "application/json", bytes.NewBuffer(nil))
	if err != nil {
		t.Error(err)
		return
	}
	if req.StatusCode != http.StatusOK {
		t.Error(req.StatusCode)
	}
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(string(b), fmt.Sprintf(":%d", sCode.c)) {
		t.Error(string(b))
	}
}
func TestHttpResponseError(t *testing.T) {

	sCode := ErrCode{123002, 500}
	errMsg := "test error"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Log("receive request")
		HttpResponseError(w, New(sCode, errMsg, nil))
	}))
	req, err := http.Post(ts.URL, "application/json", bytes.NewBuffer(nil))
	if err != nil {
		t.Error(err)
		return
	}
	if req.StatusCode != sCode.s {
		t.Error(req.StatusCode, "!=", sCode.s)
	}
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(string(b), fmt.Sprintf(":%d", sCode.c)) {
		t.Error(string(b))
	}
	if !strings.Contains(string(b), errMsg) {
		t.Error(string(b))
	}
}