package cmsv6

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestC100_Encode(t *testing.T) {
	c := C100{
		Header: Header{
			MessageID:    "dc0056",
			PacketNumber: 1,
			Type:         "C100",
			DeviceID:     "0900000",
			Timestamp:    time.Date(2020, time.September, 24, 11, 29, 41, 0, time.UTC),
		},
		RequestType:      "V101",
		RequestTimestamp: time.Date(2020, time.September, 24, 11, 29, 40, 0, time.UTC),
		ExtraFields:      []string{"0", "1", "1"},
	}

	assert.Equal(t, c.Encode(), "$$dc0056,1,C100,0900000,,200924 112941,V101,200924 112940,0,1,1,#")
}
