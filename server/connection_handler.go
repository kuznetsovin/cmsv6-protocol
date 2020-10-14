package server

import (
	"cmsv6-protocol/cmsv6"
	"cmsv6-protocol/store"
	"github.com/sirupsen/logrus"
	"io"
	"net"
	"time"
)

func connHandler(c net.Conn, db store.Store) {
	var (
		response string
	)
	defer c.Close()
	for {
		buf := make([]byte, 1024)
		readLen, err := c.Read(buf)
		switch err {
		case nil:
			break
		case io.EOF:
			continue
		default:
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
		case *cmsv6.V114:
			p := store.GeoPoint{DeviceID: m.DeviceID, NavTime: m.Timestamp, Lat: m.Latitude, Lon: m.Longitude}
			if err := db.Save(p); err != nil {
				logrus.Error("Error save geo data ", err)
			}
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
