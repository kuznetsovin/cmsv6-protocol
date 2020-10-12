package main

import (
	"cmsv6-protocol/server"
	"github.com/sirupsen/logrus"
)

func main() {
	addr := ":6608"
	logLevel := logrus.DebugLevel

	logrus.SetLevel(logLevel)
	srv := server.New(addr)

	logrus.Error(srv.Start())
}
