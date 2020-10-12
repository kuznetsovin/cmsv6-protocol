package server

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net"
	"testing"
	"time"
)

func TestServer_Start(t *testing.T) {
	logrus.SetOutput(ioutil.Discard)
	srv := ":6608"
	s := New(srv)
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
}
