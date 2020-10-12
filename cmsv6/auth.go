package cmsv6

import (
	"errors"
	"strings"
)

type V101 struct {
	Header
	CommonGPS
	UnknownFields []string
}

func (v *V101) Decode(command string) error {
	msg := strings.Split(command, ",")

	if len(msg) < 13 {
		return errors.New("Incorrect message len.")
	}
	err := v.Header.Init(msg[:6])
	if err != nil {
		return err
	}

	err = v.CommonGPS.Init(msg[6:13])
	if err != nil {
		return err
	}

	v.UnknownFields = msg[13:]
	return nil
}
