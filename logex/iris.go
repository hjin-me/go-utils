package logex

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"github.com/kataras/pio"
)

func WrapIris(app *iris.Application, projectName string) {
	app.Logger().SetPrefix("[" + projectName + "]").SetTimeFormat("[2006-01-02T15:04:05.999-07:00]")
	app.Logger().Printer.Hijack(logHijacker)
}
func logHijacker(ctx *pio.Ctx) {
	l, ok := ctx.Value.(*golog.Log)
	if !ok {
		ctx.Next()
		return
	}
	line := ""
	if t := l.FormatTime(); t != "" {
		line += t
	}
	lineLevel := golog.GetTextForLevel(l.Level, ctx.Printer.IsTerminal)
	if lineLevel != "" {
		line += lineLevel
	}

	line += l.Message

	var b []byte
	if pref := l.Logger.Prefix; len(pref) > 0 {
		b = append(pref, []byte(line)...)
	} else {
		b = []byte(line)
	}

	ctx.Store(b, nil)
	ctx.Next()
}
