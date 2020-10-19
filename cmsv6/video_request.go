package cmsv6

import (
	"fmt"
	"time"
)

//44108896,1,0,1,0,192.168.32.117,6602#
type C508 struct {
	Header
	UnknownUID int
	MediaType  int
	Type       int
	AvType     int
	Channel    int
	SrvIp      string
	SrvPort    int
}

func (c *C508) Encode() string {
	return fmt.Sprintf("$$%s,%d,%d,%d,%d,%d,%s,%d#", c.Header.Encode(), c.UnknownUID, c.MediaType, c.Type,
		c.AvType, c.Channel, c.SrvIp, c.SrvPort)
}

func CreateVideoRequest(pkgNum int, date time.Time, deviceID, ip string) string {
	c := C508{
		Header: Header{
			MessageID:    "dc0067",
			PacketNumber: pkgNum,
			Type:         "C508",
			DeviceID:     deviceID,
			Timestamp:    date,
		},
		UnknownUID: 44108896,
		MediaType:  1,
		Type:       0,
		AvType:     1,
		Channel:    0,
		SrvIp:      ip,
		SrvPort:    6602,
	}
	return c.Encode()
}
