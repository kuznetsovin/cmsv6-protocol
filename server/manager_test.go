package server

import (
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
	"time"
)

func TestDeviceRegistry(t *testing.T) {
	d := NewDeviceRegistry()
	testAddr := ":3333"
	deviceID := "1"

	go func() {
		l, err := net.Listen("tcp", testAddr)
		if !assert.NoError(t, err) {
			return
		}
		defer l.Close()

		for {
			if c, err := l.Accept(); !assert.NoError(t, err) {
				return
			} else {
				go func(conn net.Conn) {
					d.AddDevice(deviceID, conn)
					buf := make([]byte, 128)
					l, _ := conn.Read(buf)
					assert.NoError(t, d.SendCommand(deviceID, string(buf[:l])))
				}(c)
			}
		}
	}()
	time.Sleep(1 * time.Second)

	conn, err := net.Dial("tcp", testAddr)
	if !assert.NoError(t, err) {
		return
	}
	defer conn.Close()

	testData := []byte("test")
	if _, err = conn.Write(testData); !assert.NoError(t, err) {
		return
	}

	buf := make([]byte, 128)
	l, _ := conn.Read(buf)
	assert.Equal(t, buf[:l], testData)

	d.RemoveDevice(deviceID)
	assert.Equal(t, len(d.registry), 0)
}
