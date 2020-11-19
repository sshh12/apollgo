package graphics

import (
	"fmt"
	"log"

	"github.com/sshh12/apollgo/server"

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
	var w, h int

	var glctx gl.Context
	for e := range app.Events() {
		switch e := app.Filter(e).(type) {
		case lifecycle.Event:
			switch e.Crosses(lifecycle.StageVisible) {
			case lifecycle.CrossOn:
				var err error
				glctx = e.DrawContext.(gl.Context)
				ctx, err := xmobilebackend.NewGLContext(glctx)
				if err != nil {
					log.Fatal(err)
				}
				cvb, err = xmobilebackend.NewOffscreen(0, 0, false, ctx)
				if err != nil {
					log.Fatalln(err)
				}
				painterb, err = xmobilebackend.New(0, 0, 0, 0, ctx)
				if err != nil {
					log.Fatalln(err)
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
			w, h = e.WidthPx, e.HeightPx
		case paint.Event:
			if glctx != nil {
				fw, fh := float64(w), float64(h)

				cvb.SetSize(w, h)

				cv.SetFillStyle("#000")
				cv.FillRect(0, 0, fw, fh)

				cv.SetFillStyle(255, 0, 0)
				cv.SetFont(nil, 50)
				cv.FillText(fmt.Sprintf("Data %d / %d", server.TotalRead, server.TotalWrite), fw*0.1, fh*0.1)

				painterb.SetBounds(0, 0, w, h)
				painter.DrawImage(cv)

				app.Publish()
				app.Send(paint.Event{})
			}
		}
	}
}
