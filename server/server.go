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

type CommandQueue chan DeviceCommand

type Server struct {
	conn          string
	videoSrvIp    string
	db            store.Store
	devices       *DeviceRegistry
	commandsQueue CommandQueue
}

func (s *Server) Start() error {
	s.devices = NewDeviceRegistry()

	l, err := net.Listen("tcp", s.conn)
	if err != nil {
		logrus.Fatal("Decode server ", err)
	}
	defer l.Close()

	go func() {
		s.startCommandSender()
	}()

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

func (s *Server) startCommandSender() {
	for c := range s.commandsQueue {
		if c.Command == nil {
			logrus.Warn("Empty command")
			continue
		}

		cmd := ""
		switch m := c.Command.(type) {
		case *cmsv6.C100:
			cmd = m.Encode()
		case *cmsv6.C508:
			m.PacketNumber = 6
			m.SrvIp = s.videoSrvIp
			cmd = m.Encode()
		default:
			continue
		}
		if err := s.devices.SendCommand(c.DeviceID, cmd); err != nil {
			logrus.WithFields(logrus.Fields{
				"device": c.DeviceID,
				"msg":    cmd,
			}).Errorf("Send packet error %v", err)
		}

		logrus.WithFields(logrus.Fields{
			"msg": cmd,
		}).Debug("Send packet")
	}
}

func (s *Server) connHandler(c net.Conn) error {
	currentDeviceID := ""
	defer func() {
		if currentDeviceID != "" {
			s.devices.RemoveDevice(currentDeviceID)
		}
	}()

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
			currentDeviceID = m.Header.DeviceID
			s.devices.AddDevice(currentDeviceID, c)

			cmd := DeviceCommand{
				DeviceID: m.DeviceID,
				Command:  cmsv6.CreateResponse(m.Header, time.Now().UTC(), []string{"0", "1", "1"}),
			}
			s.commandsQueue <- cmd
		case *cmsv6.V141:
			cmd := DeviceCommand{
				DeviceID: m.DeviceID,
				Command: cmsv6.CreateResponse(m.Header, time.Now().UTC(),
					[]string{"0", "0", "0", "0", "", "", "0", "", "0", ""}),
			}
			s.commandsQueue <- cmd
		case *cmsv6.V114:
			p := store.GeoPoint{DeviceID: m.DeviceID, NavTime: m.Timestamp, Lat: m.Latitude, Lon: m.Longitude}
			if err := s.db.Save(p); err != nil {
				return fmt.Errorf("Error save geo data %v", err)
			}
		case *cmsv6.V100:
			logrus.Info("Received video response")
			// TODO: start video service
		default:
			logrus.Warn("Unknown type")
			continue
		}
	}
}

func New(conn, videoAddr string, cmdQueue CommandQueue) *Server {
	return &Server{conn: conn, videoSrvIp: videoAddr, db: store.NewStore(""), commandsQueue: cmdQueue}
}

func CreateVideoRequestCommand(deviceID string) DeviceCommand {
	ts := time.Now().UTC()

	return DeviceCommand{
		DeviceID: deviceID,
		Command:  cmsv6.CreateVideoRequest(ts, deviceID),
	}
}
