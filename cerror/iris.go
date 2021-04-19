package cerror

import (
	"errors"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/hjin-me/go-utils/v2/logex"
	"github.com/kataras/iris/v12"
)

func ParseValidateError(err error) string {
	var vErrs validator.ValidationErrors
	if errors.As(err, &vErrs) {
		errFields := make([]string, len(vErrs))
		for i, e := range vErrs {
			errFields[i] = e.Field()
		}
		return "request json validate failed. fields [" + strings.Join(errFields, ", ") + "] are illegal."
	}
	return err.Error()
}

func ResponseError(ctx iris.Context, err error) {
	var cErr Error
	if errors.As(err, &cErr) {
		ctx.StatusCode(cErr.StatusCode())
		var vErrs validator.ValidationErrors
		if errors.As(cErr.Unwrap(), &vErrs) {
			errFields := make([]string, len(vErrs))
			for i, e := range vErrs {
				logex.Debug(e.Tag(), " ", e.ActualTag())
				errFields[i] = e.Field()
			}
			cErr.Msg = "request json validate failed. fields [" + strings.Join(errFields, ", ") + "] are illegal."
		}
		_, e := ctx.JSON(cErr)
		if e != nil {
			logex.Warningf("response error failed, %v", err)
		}
	} else {
		ctx.StatusCode(http.StatusInternalServerError)
		body := struct {
			Code uint   `json:"err_code"`
			Msg  string `json:"err_msg"`
		}{
			Code: internalErrCode,
			Msg:  err.Error(),
		}
		_, e := ctx.JSON(body)
		if e != nil {
			logex.Warningf("response error failed, %v", err)
		}
	}

}

func ResponseSuccess(ctx iris.Context, data interface{}) {
	ctx.StatusCode(http.StatusOK)
	resp := map[string]interface{}{"err_code": successCode, "err_msg": ""}
	if data != nil {
		resp["data"] = data
	}
	_, err := ctx.JSON(resp)
	if err != nil {
		logex.Warningf("response error failed, %v", err)
	}
}

func DefineInternalServerErrCode(c uint) {
	internalErrCode = c
}

func DefineSuccessCode(code ErrCode) {
	successCode = code
}
