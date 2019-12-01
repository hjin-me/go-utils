package cerror

import (
	"encoding/json"
	"errors"
	"github.com/hjin-me/go-utils/logex"
	"github.com/kataras/iris/v12"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"strings"
)

func HttpResponseError(w http.ResponseWriter, err Error) {
	w.WriteHeader(err.StatusCode())
	var vErrs validator.ValidationErrors
	if errors.As(err.Unwrap(), &vErrs) {
		errFields := make([]string, len(vErrs))
		for i, e := range vErrs {
			errFields[i] = e.Field()
		}
		err.Msg = "request json validate failed. fields [" + strings.Join(errFields, ", ") + "] are illegal."
	}
	b, e := json.Marshal(err)
	if e != nil {
		logex.Warningf("json marshal failed, %v", err)
	}
	_, e = w.Write(b)
	if e != nil {
		logex.Warningf("response error failed, %v", err)
	}
}

func HttpResponseSuccess(w http.ResponseWriter, data interface{}) {
	w.WriteHeader(http.StatusOK)
	resp := iris.Map{"err_code": successCode, "err_msg": ""}
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
	resp := iris.Map{"err_code": errCode, "err_msg": ""}
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
