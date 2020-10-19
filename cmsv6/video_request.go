package cmsv6

import (
	"fmt"
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

func (c *C508) Decode() string {
	return fmt.Sprintf("$$%s,%d,%d,%d,%d,%d,%s,%d#", c.Header.Encode(), c.UnknownUID, c.MediaType, c.Type,
		c.AvType, c.Channel, c.SrvIp, c.SrvPort)
}
