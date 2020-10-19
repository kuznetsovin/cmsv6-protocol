package rpc

import (
	"cmsv6-protocol/server"
	"fmt"
	"net/http"
)

type RPC struct {
	commandQueue server.CommandQueue
}

func (rpc *RPC) StartServer() error {
	http.HandleFunc("/start-live", func(w http.ResponseWriter, r *http.Request) {
		devID := r.URL.Query().Get("device_id")
		if devID == "" {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = fmt.Fprintf(w, "Device not set")
			return
		}

		rpc.commandQueue <- server.CreateVideoRequestCommand(devID)

		_, _ = fmt.Fprintf(w, "Command send")
	})

	return http.ListenAndServe(":8089", nil)
}

func NewRPC(cmdBuf server.CommandQueue) RPC {
	return RPC{cmdBuf}
}
