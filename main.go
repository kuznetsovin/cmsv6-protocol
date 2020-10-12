package main

import (
	"cmsv6-protocol/server"
	"github.com/sirupsen/logrus"
)

func main() {
	addr := ":6608"
	srv := server.New(addr)

	logrus.Error(srv.Start())
}
