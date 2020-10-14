package server

import (
	"cmsv6-protocol/store"
	"github.com/sirupsen/logrus"
	"net"
)

type Server struct {
	conn string
	db   store.Store
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
			go connHandler(c, s.db)
		}
	}
}

func New(conn string, db store.Store) *Server {
	return &Server{conn: conn, db: db}
}
