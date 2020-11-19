package graphics

import (
	"container/list"
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

	speeds := list.New()

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
				speeds.PushBack(server.CurrentSpeed)
				if speeds.Len() == w/2 {
					speeds.Remove(speeds.Front())
				}
				draw(cv, fw, fh, speeds)
				painterb.SetBounds(0, 0, w, h)
				painter.DrawImage(cv)
				app.Publish()
				app.Send(paint.Event{})
			}
		}
	}
}

func draw(cv *canvas.Canvas, w float64, h float64, speeds *list.List) {
	md := w
	if h > w {
		md = h
	}
	cv.SetFillStyle(240, 240, 240)
	cv.SetFont(nil, 50)
	cv.FillText(fmt.Sprintf("Data %s / %s (%d conns)", byteCountSI(server.TotalRead), byteCountSI(server.TotalWrite), server.ActiveConn), md*0.1, md*0.1)
	cv.FillText(fmt.Sprintf("Speed %s/s", byteCountSI(server.CurrentSpeed)), md*0.1, md*0.1+60)
	cv.FillText(server.ExternalIP, md*0.1, md*0.1+120)

	x := 0
	var maxV int64 = 0
	for e := speeds.Front(); e != nil; e = e.Next() {
		v := e.Value.(int64)
		cv.FillRect(float64(x), md*0.1+180, 2, float64(v)/10000+10)
		x += 2
		if v > maxV {
			maxV = v
		}
	}
	if maxV > 0 {
		cv.SetFillStyle(0, 240, 240)
		cv.FillRect(0, md*0.1+180+float64(maxV)/10000+10, w, 2)
	}
}

func byteCountSI(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(b)/float64(div), "kMGTPE"[exp])
}
