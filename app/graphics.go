package app

import (
	"fmt"

	"github.com/tfriedel6/canvas"
	"github.com/tfriedel6/canvas/backend/xmobilebackend"
	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
	"golang.org/x/mobile/gl"
)

// OnAppLaunch handles app launch and event loop
func OnAppLaunch(app app.App) {
	var cv, painter *canvas.Canvas
	var cvb *xmobilebackend.XMobileBackendOffscreen
	var painterb *xmobilebackend.XMobileBackend
	var glctx gl.Context
	var width int
	var height int
	for e := range app.Events() {
		switch e := app.Filter(e).(type) {
		case lifecycle.Event:
			switch e.Crosses(lifecycle.StageVisible) {
			case lifecycle.CrossOn:
				var err error
				glctx = e.DrawContext.(gl.Context)
				ctx, err := xmobilebackend.NewGLContext(glctx)
				if err != nil {
					fmt.Print(err)
					return
				}
				cvb, err = xmobilebackend.NewOffscreen(0, 0, false, ctx)
				if err != nil {
					fmt.Print(err)
					return
				}
				painterb, err = xmobilebackend.New(0, 0, 0, 0, ctx)
				if err != nil {
					fmt.Print(err)
					return
				}
				cv = canvas.New(cvb)
				cv.LoadFont(robotoFont)
				painter = canvas.New(painterb)
				app.Send(paint.Event{})
			case lifecycle.CrossOff:
				cvb.Delete()
				glctx = nil
			}
		case size.Event:
			width, height = e.WidthPx, e.HeightPx
		case paint.Event:
			if glctx != nil {
				fw, fh := float64(width), float64(height)
				cvb.SetSize(width, height)
				cv.SetFillStyle("#eee")
				cv.FillRect(0, 0, fw, fh)
				cv.SetFont(nil, 50)
				cv.SetFillStyle("#222")
				cv.SetTextAlign(canvas.Center)
				cv.SetTextBaseline(canvas.Middle)
				cv.FillText("view in browser", fw/2, fh/2)
				painterb.SetBounds(0, 0, width, height)
				painter.DrawImage(cv)
				app.Publish()
				app.Send(paint.Event{})
			}
		}
	}
}
