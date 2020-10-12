package server

import (
	"cmsv6-protocol/cmsv6"
	"github.com/sirupsen/logrus"
	"net"
	"time"
)

func connHandler(c net.Conn) {
	defer c.Close()
	for {
		buf := make([]byte, 1024)
		readLen, err := c.Read(buf)
		if err != nil {
			logrus.Error("Received error ", err)
			return
		}

		rawMsg := string(buf[:readLen])
		msg, err := cmsv6.ParsePacket(rawMsg)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"msg": rawMsg,
			}).Error("Received error ", err)
			return
		}

		switch m := msg.(type) {
		case *cmsv6.V101:
			logrus.WithFields(logrus.Fields{
				"type": "V101",
				"msg":  rawMsg,
			}).Info("Received packet")
			r := cmsv6.CreateResponse(m.Header, time.Now().UTC(), []string{"0", "1", "1"})
			_, err = c.Write([]byte(r))
			if err != nil {
				logrus.Error("Send error ", err)
				return
			}
			logrus.WithFields(logrus.Fields{
				"type": "C100",
				"msg":  r,
			}).Info("Send response packet")
		default:
			logrus.Error("Unknown type")
			return
		}

	}
}
