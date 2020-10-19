package cmsv6

import (
	"errors"
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

func CreateVideoRequest(date time.Time, deviceID string) Encoder {
	return &C508{
		Header: Header{
			MessageID: "dc0067",
			Type:      "C508",
			DeviceID:  deviceID,
			Timestamp: date,
		},
		UnknownUID: 44108896,
		MediaType:  1,
		Type:       0,
		AvType:     1,
		Channel:    0,
		SrvPort:    6602,
	}
}

type V100 struct {
	gpsInfo
	RequestType      string
	RequestTimestamp time.Time
	ExtraFields      []string
}

func (v *V100) Decode(msg Message) error {
	if len(msg) < 24 {
		return errors.New("Incorrect V100 message len.")
	}

	err := v.gpsInfo.decode(msg[:24])
	if err != nil {
		return err
	}

	v.RequestType = msg[25]

	v.RequestTimestamp, err = time.Parse(timestampFmt, msg[26])
	if err != nil {
		return fmt.Errorf("Incorrect v100 timestamp format: %v", err)
	}

	v.ExtraFields = msg[27:]
	return err
}
