package cmsv6

import (
	"errors"
)

type V101 struct {
	gpsInfo
	AuthInfo []string
}

func (v *V101) Decode(msg Message) error {
	if len(msg) < 24 {
		return errors.New("Incorrect message len.")
	}

	if err := v.gpsInfo.decode(msg[:24]); err != nil {
		return err
	}

	v.AuthInfo = msg[24:]
	return nil
}

type V141 struct {
	V114
}

func (v *V141) Decode(msg Message) error {
	return v.V114.Decode(msg)
}
