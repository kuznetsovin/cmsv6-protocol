package cmsv6

import (
	"errors"
	"strings"
)

type Decoder interface {
	Decode(Message) error
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
	default:
		return result, errors.New("Unknown type.")
	}

	err = result.Decode(m)
	return result, err
}

func parseMsg(msg string) ([]string, error) {
	if !(strings.HasPrefix(msg, "$$") && strings.HasSuffix(msg, "#")) {
		return nil, errors.New("Incorrect packet")
	}
	clearCmd := strings.Trim(msg, "$$#")

	return strings.Split(clearCmd, ","), nil
}
