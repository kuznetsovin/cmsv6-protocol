package main

import (
	"cmsv6-protocol/server"
	"cmsv6-protocol/store"
	"github.com/sirupsen/logrus"
)

func main() {
	addr := ":6608"
	logLevel := logrus.DebugLevel

	logrus.SetLevel(logLevel)
	db := store.NewStore("")
	cmdBuf := make(server.CommandQueue, 1000000)
	srv := server.New(addr, db, cmdBuf)

	logrus.Error(srv.Start())
}
