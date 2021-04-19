package cerror

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/hjin-me/go-utils/v2/logex"
)

func HttpResponseError(w http.ResponseWriter, err error) {
	var cErr Error
	if errors.As(err, &cErr) {
		w.WriteHeader(cErr.StatusCode())
		var vErrs validator.ValidationErrors
		if errors.As(cErr.Unwrap(), &vErrs) {
			errFields := make([]string, len(vErrs))
			for i, e := range vErrs {
				errFields[i] = e.Field()
			}
			cErr.Msg = "request json validate failed. fields [" + strings.Join(errFields, ", ") + "] are illegal."
		}
		b, e := json.Marshal(cErr)
		if e != nil {
			logex.Warningf("json marshal failed, %v", cErr)
		}
		_, e = w.Write(b)
		if e != nil {
			logex.Warningf("response error failed, %v", cErr)
		}
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		body := struct {
			Code uint   `json:"err_code"`
			Msg  string `json:"err_msg"`
		}{
			Code: internalErrCode,
			Msg:  err.Error(),
		}
		b, e := json.Marshal(body)
		if e != nil {
			logex.Warningf("json marshal failed, %v", cErr)
		}
		_, e = w.Write(b)
		if e != nil {
			logex.Warningf("response error failed, %v", cErr)
		}
	}
}

func HttpResponseSuccess(w http.ResponseWriter, data interface{}) {
	w.WriteHeader(http.StatusOK)
	resp := map[string]interface{}{"err_code": successCode, "err_msg": ""}
	if data != nil {
		resp["data"] = data
	}
	b, e := json.Marshal(resp)
	if e != nil {
		logex.Warningf("json marshal failed, %v", e)
	}
	_, e = w.Write(b)
	if e != nil {
		logex.Warningf("response error failed, %v", e)
	}
}

func HttpResponseSuccessWithCode(w http.ResponseWriter, errCode int, data interface{}) {
	w.WriteHeader(http.StatusOK)
	resp := map[string]interface{}{"err_code": errCode, "err_msg": ""}
	if data != nil {
		resp["data"] = data
	}
	b, e := json.Marshal(resp)
	if e != nil {
		logex.Warningf("json marshal failed, %v", e)
	}
	_, e = w.Write(b)
	if e != nil {
		logex.Warningf("response error failed, %v", e)
	}
}
