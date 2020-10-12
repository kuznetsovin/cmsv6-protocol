package cmsv6

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

var (
	timestampFmt = "060102 150405"
)

type Header struct {
	MessageID    string
	PacketNumber int
	Type         string
	DeviceID     string
	Timestamp    time.Time
}

func (h *Header) Decode(fields []string) error {
	var (
		err error
	)
	if len(fields) != 6 {
		return errors.New("Incorrect header slice size.")
	}

	h.MessageID = fields[0]
	h.PacketNumber, err = strconv.Atoi(fields[1])
	if err != nil {
		return fmt.Errorf("Incorrect packet number format: %v", err)
	}

	h.Type = fields[2]
	h.DeviceID = fields[3]

	h.Timestamp, err = time.Parse(timestampFmt, fields[5])
	if err != nil {
		return fmt.Errorf("Incorrect timestamp format: %v", err)
	}

	return err
}

func (h *Header) Encode() string {
	return fmt.Sprintf("%s,%d,%s,%s,,%s", h.MessageID, h.PacketNumber, h.Type, h.
		DeviceID, h.Timestamp.Format(timestampFmt))
}

type CommonGPS struct {
	State     string
	Latitude  float64
	Longitude float64
}

func (c *CommonGPS) Decode(fields []string) error {
	var (
		err error
	)

	if len(fields) != 7 {
		return errors.New("Incorrect gps data slice size.")
	}

	c.State = fields[0]
	c.Latitude, err = sliceToGeoCoord(fields[1:4])
	if err != nil {
		return fmt.Errorf("Incorrect latitude format: %v", err)
	}
	c.Longitude, err = sliceToGeoCoord(fields[4:7])
	if err != nil {
		return fmt.Errorf("Incorrect longitude format: %v", err)
	}
	return nil
}

func sliceToGeoCoord(s []string) (float64, error) {
	var (
		result float64
	)
	if len(s) != 3 {
		return result, errors.New("Incorrect coord slice size.")
	}

	degree, err := strconv.ParseFloat(s[0], 64)
	if err != nil {
		return result, fmt.Errorf("Incorrect degree format: %v", err)
	}
	result += degree

	m, err := strconv.ParseFloat(s[1], 64)
	if err != nil {
		return result, fmt.Errorf("Incorrect minute format: %v", err)
	}
	result += m / 60

	sec, err := strconv.ParseFloat(s[2], 64)
	if err != nil {
		return result, fmt.Errorf("Incorrect second format: %v", err)
	}
	result += sec / 36000000000

	return result, err
}
