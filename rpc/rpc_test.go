package rpc

import (
	"cmsv6-protocol/server"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestRPC(t *testing.T) {
	queue := make(server.CommandQueue, 10)
	rpc := NewRPC(queue)

	go func() {
		err := rpc.StartServer()
		if assert.NoError(t, err) {
			return
		}
	}()

	resp, err := http.Get("http://localhost:8089/start-live?device_id=0900000")
	if assert.NoError(t, err) {
		return
	}
	defer resp.Body.Close()

	if body, err := ioutil.ReadAll(resp.Body); assert.NoError(t, err) {
		assert.Equal(t, body, []byte("Command send"))
		c := <-rpc.commandQueue
		assert.Equal(t, c.DeviceID, "0900000")
	}
}
