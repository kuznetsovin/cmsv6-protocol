package main

import (
	rpc2 "cmsv6-protocol/rpc"
	"cmsv6-protocol/server"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"net"
	"net/http"
	"testing"
)

func TestService(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)
	testAddr := ":7000"

	cmdBuf := make(server.CommandQueue, 1000)

	rpc := rpc2.NewRPC(cmdBuf)

	go func() {
		logrus.Error("RPC error in ", rpc.StartServer())
	}()

	srv := server.New(testAddr, "127.0.0.1", cmdBuf)
	go func() {
		logrus.Error("Srv error in ", srv.Start())
	}()

	conn, err := net.Dial("tcp", testAddr)
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

	go func() {
		resp, err := http.Get("http://localhost:8089/start-live?device_id=0900000")
		if assert.NoError(t, err) {
			return
		}
		defer resp.Body.Close()
	}()

	rcvLen, err = conn.Read(buf)
	if assert.NoError(t, err) {
		assert.NotEmpty(t, string(buf[:rcvLen]))
	}
}
