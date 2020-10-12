package server

import (
	"github.com/sirupsen/logrus"
	"net"
)

type Server struct {
	conn string
}

func (s *Server) Start() error {
	l, err := net.Listen("tcp", s.conn)
	if err != nil {
		logrus.Fatal("Decode server ", err)
	}
	defer l.Close()

	logrus.Infof("Starting server on %s...", s.conn)
	for {
		if c, err := l.Accept(); err != nil {
			logrus.Errorf("Connection error %v", err)
		} else {
			go connHandler(c)
		}
	}
}

func New(conn string) *Server {
	return &Server{conn: conn}
}
