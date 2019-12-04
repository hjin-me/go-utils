package cerror

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/httptest"
	"testing"
)

func TestResponseSuccess(t *testing.T) {
	sCode := ErrCode{123000, 200}
	DefineSuccessCode(sCode)
	app := iris.New()
	app.Get("/", func(context context.Context) {
		ResponseSuccess(context, nil)
	})
	expect := httptest.New(t, app)
	expect.GET("/").Expect().Status(sCode.s).
		Body().Contains(fmt.Sprintf(":%d", sCode.c))
}
func TestResponseError(t *testing.T) {

	sCode := ErrCode{123001, 400}
	app := iris.New()
	errMsg := "errMsgerrMsgerrMsgerrMsg"
	app.Get("/", func(context context.Context) {
		ResponseError(context, New(sCode, errMsg, nil))
	})
	expect := httptest.New(t, app)
	expect.GET("/").Expect().Status(sCode.s).
		Body().Contains(fmt.Sprintf(":%d", sCode.c)).Contains(errMsg)
}
