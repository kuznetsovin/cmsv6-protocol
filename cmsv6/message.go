package cmsv6

import (
	"errors"
	"strings"
	"time"
)

type Decoder interface {
	Decode(Message) error
}

type Encoder interface {
	Encode() string
}

type Message []string

func ParsePacket(packet string) (interface{}, error) {
	var (
		result Decoder
	)
	m, err := parseMsg(packet)
	if err != nil {
		return nil, err
	}

	switch m[2] {
	case "V101":
		result = &V101{}
	case "V141":
		result = &V141{}
	case "V114":
		result = &V114{}
	case "V100":
		result = &V100{}
	default:
		return result, errors.New("Unknown type.")
	}

	err = result.Decode(m)
	return result, err
}

func CreateResponse(reqHeader Header, respTime time.Time, extra []string) Encoder {
	return &C100{
		Header: Header{
			MessageID:    "dc0056",
			PacketNumber: reqHeader.PacketNumber,
			Type:         "C100",
			DeviceID:     reqHeader.DeviceID,
			Timestamp:    respTime,
		},
		RequestType:      reqHeader.Type,
		RequestTimestamp: reqHeader.Timestamp,
		ExtraFields:      extra,
	}
}

func parseMsg(msg string) ([]string, error) {
	if !(strings.HasPrefix(msg, "$$") && strings.HasSuffix(msg, "#")) {
		return nil, errors.New("Incorrect packet")
	}
	clearCmd := strings.Trim(msg, "$$#")

	return strings.Split(clearCmd, ","), nil
}
