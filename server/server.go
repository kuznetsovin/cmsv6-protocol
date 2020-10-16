package server

import (
	"cmsv6-protocol/cmsv6"
	"cmsv6-protocol/store"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net"
	"time"
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
			go func(conn net.Conn) {
				defer c.Close()
				if err := s.connHandler(conn); err != nil {
					logrus.Error(err)
				}
			}(c)
		}
	}
}

func (s *Server) connHandler(c net.Conn) error {
	var (
		response string
	)
	for {
		buf := make([]byte, 1024)
		readLen, err := c.Read(buf)
		switch err {
		case nil:
			break
		case io.EOF:
			continue
		default:
			return fmt.Errorf("Received error %v", err)
		}

		rawMsg := string(buf[:readLen])
		logrus.WithFields(logrus.Fields{
			"msg": rawMsg,
		}).Debug("Received packet")

		msg, err := cmsv6.ParsePacket(rawMsg)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"msg": rawMsg,
			}).Warn("Incorrect packet")
			return fmt.Errorf("Parse packet error %v", err)
		}

		switch m := msg.(type) {
		case *cmsv6.V101:
			response = cmsv6.CreateResponse(m.Header, time.Now().UTC(), []string{"0", "1", "1"})
		case *cmsv6.V141:
			response = cmsv6.CreateResponse(m.Header, time.Now().UTC(), []string{"0", "0", "0", "0", "", "", "0", "", "0", ""})
		case *cmsv6.V114:
			p := store.GeoPoint{DeviceID: m.DeviceID, NavTime: m.Timestamp, Lat: m.Latitude, Lon: m.Longitude}
			if err := s.db.Save(p); err != nil {
				return fmt.Errorf("Error save geo data %v", err)
			}
		default:
			logrus.Warn("Unknown type")
			continue
		}

		if response != "" {
			_, err = c.Write([]byte(response))
			if err != nil {
				return fmt.Errorf("Send response error %v", err)
			}
			logrus.WithFields(logrus.Fields{
				"msg": response,
			}).Debug("Send response packet")
		}
	}
}

func New(conn string, db store.Store) *Server {
	return &Server{conn: conn, db: db}
}
