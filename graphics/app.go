package graphics

import (
	"container/list"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/sshh12/apollgo/server"

	"github.com/tfriedel6/canvas"
	"github.com/tfriedel6/canvas/backend/xmobilebackend"
	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
	"golang.org/x/mobile/gl"
)

type appState struct {
	width   int
	height  int
	read    *int64
	write   *int64
	speed   *int64
	conns   *int64
	ip      string
	dlSpeed int64
	upSpeed int64
	speeds  *list.List
	lock    sync.Mutex
}

// OnAppLaunch handles app launch and event loop
func OnAppLaunch(app app.App) {
	var cv, painter *canvas.Canvas
	var cvb *xmobilebackend.XMobileBackendOffscreen
	var painterb *xmobilebackend.XMobileBackend
	var glctx gl.Context
	state := &appState{
		read:   &server.TotalRead,
		write:  &server.TotalWrite,
		speed:  &server.CurrentSpeed,
		conns:  &server.ActiveConn,
		speeds: list.New(),
		lock:   sync.Mutex{},
	}
	go func() {
		tick := time.NewTicker(100 * time.Millisecond)
		for {
			select {
			case <-tick.C:
				state.lock.Lock()
				state.speeds.PushBack(server.CurrentSpeed)
				if state.speeds.Len() == state.width/2 {
					state.speeds.Remove(state.speeds.Front())
				}
				state.lock.Unlock()
			}
		}
	}()
	go func() {
		tick := time.NewTicker(15 * time.Minute)
		for ; true; <-tick.C {
			ip, err := server.ExternalIP()
			if err == nil {
				state.ip = ip
			}
			dl, up, err := server.RunSpeedTest()
			if err != nil {
				log.Println(err)
				continue
			}
			state.dlSpeed = int64(dl)
			state.upSpeed = int64(up)
		}
	}()
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
			state.width, state.height = e.WidthPx, e.HeightPx
		case paint.Event:
			doPaint(app, glctx, cvb, cv, painter, painterb, state)
		}
	}
}

func doPaint(app app.App, glctx gl.Context, cvb *xmobilebackend.XMobileBackendOffscreen, cv *canvas.Canvas, painter *canvas.Canvas, painterb *xmobilebackend.XMobileBackend, state *appState) {
	if glctx == nil || state.width == 0 || state.height == 0 {
		return
	}
	fw, fh := float64(state.width), float64(state.height)
	cvb.SetSize(state.width, state.height)
	cv.SetFillStyle("#000")
	cv.FillRect(0, 0, fw, fh)
	draw(cv, fw, fh, state)
	painterb.SetBounds(0, 0, state.width, state.height)
	painter.DrawImage(cv)
	app.Publish()
	app.Send(paint.Event{})
}

func draw(cv *canvas.Canvas, w float64, h float64, state *appState) {
	cx := 60.0
	cy := 140.0
	state.lock.Lock()
	defer state.lock.Unlock()
	cv.SetFillStyle(240, 240, 240)
	cv.SetFont(nil, 50)
	cv.FillText(fmt.Sprintf("Data %s / %s (%d conns)", byteCountSI(*state.read), byteCountSI(*state.write), *state.conns), cx, cy)
	cv.FillText(fmt.Sprintf("Transfer %s/s", byteCountSI(*state.speed)), cx, cy+60)
	cv.FillText(fmt.Sprintf("SpeedTest %s/s %s/s", byteCountSI(state.dlSpeed), byteCountSI(state.upSpeed)), cx, cy+120)
	cv.FillText(state.ip, cx, cy+180)
	x := 0
	var maxV int64 = 0
	for e := state.speeds.Front(); e != nil; e = e.Next() {
		v := e.Value.(int64)
		cv.FillRect(float64(x), cy+240, 2, float64(v)/10000+10)
		x += 2
		if v > maxV {
			maxV = v
		}
	}
	if maxV > 0 {
		cv.SetFillStyle(0, 240, 240)
		cv.FillRect(0, cy+240+float64(maxV)/10000+10, w, 2)
	}
	if state.dlSpeed > 0 {
		cv.SetFillStyle(240, 0, 240)
		cv.FillRect(0, cy+240+float64(state.dlSpeed)/10000+10, w, 2)
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
