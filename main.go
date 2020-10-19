package main

import (
	rpc2 "cmsv6-protocol/rpc"
	"cmsv6-protocol/server"
	"github.com/sirupsen/logrus"
)

func main() {
	addr := ":6608"
	videoAddr := "192.168.0.117"

	logLevel := logrus.DebugLevel

	logrus.SetLevel(logLevel)

	cmdBuf := make(server.CommandQueue, 1000000)

	rpc := rpc2.NewRPC(cmdBuf)

	go func() {
		logrus.Error("RPC error in ", rpc.StartServer())
	}()

	srv := server.New(addr, videoAddr, cmdBuf)

	logrus.Error(srv.Start())
}
