package cmsv6

import "errors"

type gpsInfo struct {
	Header
	CommonGPS
}

func (g *gpsInfo) decode(fields []string) error {
	if len(fields) < 24 {
		return errors.New("Incorrect message len.")
	}
	if err := g.Header.Decode(fields[:6]); err != nil {
		return err
	}

	if err := g.CommonGPS.Decode(fields[6:24]); err != nil {
		return err
	}

	return nil
}

type V114 struct {
	gpsInfo
	Ack string
}

func (v *V114) Decode(msg Message) error {
	lastIdx := len(msg) - 1
	if err := v.gpsInfo.decode(msg[:lastIdx]); err != nil {
		return err
	}
	v.Ack = msg[lastIdx]
	return nil
}
