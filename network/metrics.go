package network

import (
	"container/list"
	"net"
	"sync/atomic"
	"time"
)

var (
	// TotalRead total bytes read
	TotalRead int64 = 0
	// TotalWrite total bytes written
	TotalWrite int64 = 0
	// ActiveConn active connections
	ActiveConn int64 = 0
	// CurrentSpeed is bytes/s
	CurrentSpeed int64 = 0
)

func init() {
	go WatchSpeed()
}

// WatchSpeed computes network speed
func WatchSpeed() {
	ticker := time.NewTicker(1 * time.Second)
	var write int64 = 0
	speeds := list.New()
	for {
		select {
		case <-ticker.C:
			newWrite := TotalWrite
			speeds.PushBack(newWrite - write)
			if speeds.Len() == 3 {
				var max int64 = 0
				for e := speeds.Front(); e != nil; e = e.Next() {
					if v := e.Value.(int64); v > max {
						max = v
					}
				}
				CurrentSpeed = max
				speeds.Remove(speeds.Front())
			}
			write = newWrite
		}
	}
}

// NewMetricConn wraps a conn
func NewMetricConn(conn net.Conn) MetricConn {
	var read int64 = 0
	var write int64 = 0
	mc := MetricConn{
		conn:       conn,
		readBytes:  &read,
		writeBytes: &write,
		quit:       make(chan struct{}),
	}
	go mc.Watch()
	return mc
}

// MetricConn is conn with metrics
type MetricConn struct {
	conn       net.Conn
	readBytes  *int64
	writeBytes *int64
	quit       chan struct{}
}

// Watch computes metrics
func (mc MetricConn) Watch() {
	ticker := time.NewTicker(250 * time.Millisecond)
	var read int64 = 0
	var write int64 = 0
	atomic.AddInt64(&ActiveConn, 1)
	for {
		select {
		case <-ticker.C:
			newRead := *mc.readBytes
			newWrite := *mc.readBytes
			atomic.AddInt64(&TotalRead, newRead-read)
			atomic.AddInt64(&TotalWrite, newWrite-write)
			read = newRead
			write = newWrite
		case <-mc.quit:
			ticker.Stop()
			atomic.AddInt64(&ActiveConn, -1)
			return
		}
	}
}

func (mc MetricConn) Read(b []byte) (int, error) {
	n, err := mc.conn.Read(b)
	*mc.readBytes += int64(n)
	return n, err
}

func (mc MetricConn) Write(b []byte) (int, error) {
	n, err := mc.conn.Write(b)
	*mc.writeBytes += int64(n)
	return n, err
}

func (mc MetricConn) Close() error {
	mc.quit <- struct{}{}
	return mc.conn.Close()
}

func (mc MetricConn) LocalAddr() net.Addr {
	return mc.conn.LocalAddr()
}

func (mc MetricConn) RemoteAddr() net.Addr {
	return mc.conn.RemoteAddr()
}

func (mc MetricConn) SetDeadline(t time.Time) error {
	return mc.conn.SetDeadline(t)
}

func (mc MetricConn) SetReadDeadline(t time.Time) error {
	return mc.conn.SetReadDeadline(t)
}

func (mc MetricConn) SetWriteDeadline(t time.Time) error {
	return mc.conn.SetWriteDeadline(t)
}
