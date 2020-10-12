package server

import (
	"cmsv6-protocol/cmsv6"
	"github.com/sirupsen/logrus"
	"net"
	"time"
)

func connHandler(c net.Conn) {
	var (
		response string
	)
	defer c.Close()
	for {
		buf := make([]byte, 1024)
		readLen, err := c.Read(buf)
		if err != nil {
			logrus.Error("Received error ", err)
			return
		}

		rawMsg := string(buf[:readLen])
		logrus.WithFields(logrus.Fields{
			"msg": rawMsg,
		}).Debug("Received packet")

		msg, err := cmsv6.ParsePacket(rawMsg)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"msg": rawMsg,
			}).Error("Received error ", err)
			return
		}

		switch m := msg.(type) {
		case *cmsv6.V101:
			response = cmsv6.CreateResponse(m.Header, time.Now().UTC(), []string{"0", "1", "1"})
		case *cmsv6.V141:
			response = cmsv6.CreateResponse(m.Header, time.Now().UTC(), []string{"0", "0", "0", "0", "", "", "0", "", "0", ""})
		default:
			logrus.Error("Unknown type")
			continue
		}

		if response != "" {
			_, err = c.Write([]byte(response))
			if err != nil {
				logrus.Error("Send response error ", err)
				return
			}
			logrus.WithFields(logrus.Fields{
				"msg": response,
			}).Debug("Send response packet")
		}

	}
}
