package main

import (
	"cmsv6-protocol/server"
	"github.com/sirupsen/logrus"
)

func main() {
	addr := ":6608"
	videoAddr := "192.168.0.117"

	logLevel := logrus.DebugLevel

	logrus.SetLevel(logLevel)

	cmdBuf := make(server.CommandQueue, 1000000)
	srv := server.New(addr, videoAddr, cmdBuf)

	logrus.Error(srv.Start())
}
