package logex

import (
	"bytes"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/httptest"
	"strings"
	"testing"
)

func TestWrapIris(t *testing.T) {
	app := iris.Default()
	var bf bytes.Buffer
	WrapIris(app, "testP")
	app.Get("/test", func(context context.Context) {
		context.Application().Logger().Error("iris test", "error")
		t.Log("request complete")
	})
	e := httptest.New(t, app)
	app.Logger().SetLevel("debug")
	app.Logger().SetOutput(&bf)
	e.GET("/test").Expect().Status(httptest.StatusOK)

	s := bf.String()
	if strings.Index(s, "[testP]") == -1 {
		t.Error("log miss", s)
	}
	if strings.Index(s, "[ERRO]iris testerror") == -1 {
		t.Error("log miss", s)
	}
	t.Log(s)
}
