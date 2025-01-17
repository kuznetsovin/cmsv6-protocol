package server

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
	"time"
)

func TestServer_Start(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)

	srv := ":6608"
	queue := make(CommandQueue, 1000000)
	s := New(srv, "", queue)
	go func() {
		assert.NoError(t, s.Start())
	}()
	time.Sleep(2 * time.Second)

	conn, err := net.Dial("tcp", srv)
	if !assert.NoError(t, err) {
		return
	}
	defer conn.Close()

	testReq := []byte("$$dc0242,1,V101,0900000,,200924 112940,V0000,37,57,432623999,-0,0,0,0.00,0,0000000000007383,0000000000000000,0.00,0.00,0.00,0,0.00,0,0|0.00|0|0|0|0|0|0|0|0.00|0|0,,V1.0.0.1,4108,,0,0,0,SZ88888,1,USER=root,13,1,37,14,0900000,V2018 0414,V6.1.48 20180122,,0,1,0,#")
	_, _ = conn.Write(testReq)

	buf := make([]byte, 1024)
	rcvLen, err := conn.Read(buf)
	if assert.NoError(t, err) {
		assert.NotEmpty(t, string(buf[:rcvLen]))
	}

	testReq = []byte("$$dc0146,3,V141,0900000,,200924 112940,V0000,-0,0,0,-0,0,0,0.00,0,0000000000007383,0000000000000000,0.00,0.00,0.00,0,0.00,0,0|0.00|0|0|0|0|0|0|0|0.00|0|0,#")
	_, _ = conn.Write(testReq)

	rcvLen, err = conn.Read(buf)
	if assert.NoError(t, err) {
		assert.NotEmpty(t, string(buf[:rcvLen]))
	}

	testReq = []byte("$$dc0165,4,V114,0900000,,200924 112942,A0000,37,57,421385999,55,49,237689999,0.00,0,0F0EE331000D7383,0000000000000000,0.00,0.00,0.00,0,0.00,0,0|0.00|0|0|0|0|0|0|0|0.00|0|0,1#")
	_, _ = conn.Write(testReq)

	testReq = []byte("$$dc0192,6,V100,0900000,,200924 113137,A0000,37,57,432905999,55,49,240083999,0.00,0,0F0EE331000D7383,0000000000000000,0.00,0.00,0.00,0,0.00,0,0|0.00|0|0|0|0|0|0|0|0.00|0|0,,C508,200924 113137,0,1,1,,0#")
	_, _ = conn.Write(testReq)
}
