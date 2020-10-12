package cmsv6

import (
	"errors"
)

type V101 struct {
	Header
	CommonGPS
	UnknownFields []string
}

func (v *V101) Decode(msg Message) error {
	if len(msg) < 13 {
		return errors.New("Incorrect message len.")
	}
	if err := v.Header.Decode(msg[:6]); err != nil {
		return err
	}

	if err := v.CommonGPS.Decode(msg[6:13]); err != nil {
		return err
	}

	v.UnknownFields = msg[13:]
	return nil
}
