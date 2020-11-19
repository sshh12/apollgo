package main

import (
	"fmt"
	"log"
	"math"
	"net"
	"time"

	"github.com/nadoo/glider/rule"
	"github.com/tfriedel6/canvas"
	"github.com/tfriedel6/canvas/backend/xmobilebackend"
	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
	"golang.org/x/mobile/gl"
)

func main() {
	go startServer()
	app.Main(func(a app.App) {
		var cv, painter *canvas.Canvas
		var cvb *xmobilebackend.XMobileBackendOffscreen
		var painterb *xmobilebackend.XMobileBackend
		var w, h int

		var glctx gl.Context
		for e := range a.Events() {
			switch e := a.Filter(e).(type) {
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
					a.Send(paint.Event{})
				case lifecycle.CrossOff:
					cvb.Delete()
					glctx = nil
				}
			case size.Event:
				w, h = e.WidthPx, e.HeightPx
			case paint.Event:
				if glctx != nil {
					fw, fh := float64(w), float64(h)
					color := math.Sin(float64(time.Now().UnixNano())*0.000000002)*0.3 + 0.7

					cvb.SetSize(w, h)

					cv.SetFillStyle("#000")
					cv.FillRect(0, 0, fw, fh)

					cv.SetFillStyle(color*0.2, color*0.2, color*0.8)
					cv.FillRect(fw*0.25, fh*0.25, fw*0.5, fh*0.5)

					cv.SetFillStyle(255, 0, 0)
					cv.SetFont(nil, 50)
					cv.FillText(fmt.Sprintf("Data %d / %d", read, write), fw*0.25, fh*0.2)

					painterb.SetBounds(0, 0, w, h)
					painter.DrawImage(cv)

					a.Publish()
					a.Send(paint.Event{})
				}
			}
		}
	})
}

var read = 0
var write = 0

type wrapperConn struct {
	conn net.Conn
}

func (wc wrapperConn) Read(b []byte) (int, error) {
	n, err := wc.conn.Read(b)
	read += n
	return n, err
}

func (wc wrapperConn) Write(b []byte) (int, error) {
	n, err := wc.conn.Write(b)
	write += n
	return n, err
}

func (wc wrapperConn) Close() error {
	return wc.conn.Close()
}

func (wc wrapperConn) LocalAddr() net.Addr {
	return wc.conn.LocalAddr()
}

func (wc wrapperConn) RemoteAddr() net.Addr {
	return wc.conn.RemoteAddr()
}

func (wc wrapperConn) SetDeadline(t time.Time) error {
	return wc.conn.SetDeadline(t)
}

func (wc wrapperConn) SetReadDeadline(t time.Time) error {
	return wc.conn.SetReadDeadline(t)
}

func (wc wrapperConn) SetWriteDeadline(t time.Time) error {
	return wc.conn.SetWriteDeadline(t)
}

func startServer() {

	pxy := rule.NewProxy(
		[]string{},
		&rule.StrategyConfig{},
		[]*rule.Config{},
	)

	// ipset manager
	// ipsetM, _ := ipset.NewManager([]*rule.Config{})

	// check and setup dns server
	// if config.DNS != "" {
	// 	d, err := dns.NewServer(config.DNS, pxy, &config.DNSConfig)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	// rules
	// 	for _, r := range []*rule.Config{} {
	// 		for _, domain := range r.Domain {
	// 			if len(r.DNSServers) > 0 {
	// 				d.SetServers(domain, r.DNSServers)
	// 			}
	// 		}
	// 	}

	// 	// add a handler to update proxy rules when a domain resolved
	// 	d.AddHandler(pxy.AddDomainIP)
	// 	if ipsetM != nil {
	// 		d.AddHandler(ipsetM.AddDomainIP)
	// 	}

	// 	d.Start()

	// 	// custom resolver
	// 	net.DefaultResolver = &net.Resolver{
	// 		PreferGo: true,
	// 		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
	// 			d := net.Dialer{Timeout: time.Second * 3}
	// 			return d.DialContext(ctx, "udp", config.DNS)
	// 		},
	// 	}
	// }

	// enable checkers
	pxy.Check()

	local, err := NewMixedServer("tcp://:8443", pxy)
	if err != nil {
		fmt.Println(err)
		return
	}
	go local.ListenAndServe()

}
