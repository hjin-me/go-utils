package cerror

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	t.Run("custom error", func(t *testing.T) {

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
	})
	t.Run("normal error", func(t *testing.T) {
		DefineInternalServerErrCode(123002)
		errMsg := "test error"
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			t.Log("receive request")
			HttpResponseError(w, errors.New(errMsg))
		}))
		req, err := http.Post(ts.URL, "application/json", bytes.NewBuffer(nil))
		require.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, req.StatusCode)
		b, err := ioutil.ReadAll(req.Body)
		require.NoError(t, err)
		if !strings.Contains(string(b), fmt.Sprintf(":%d", 123002)) {
			t.Error(string(b))
		}
		if !strings.Contains(string(b), errMsg) {
			t.Error(string(b))
		}
	})
}
