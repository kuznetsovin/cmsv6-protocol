package cmsv6

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestC508_Decode(t *testing.T) {
	c := C508{
		Header: Header{
			MessageID:    "dc0067",
			PacketNumber: 6,
			Type:         "C508",
			DeviceID:     "0900000",
			Timestamp:    time.Date(2020, time.September, 24, 11, 31, 37, 0, time.UTC),
		},
		UnknownUID: 44108896,
		MediaType:  1,
		Type:       0,
		AvType:     1,
		Channel:    0,
		SrvIp:      "192.168.32.117",
		SrvPort:    6602,
	}

	assert.Equal(t, c.Decode(), "$$dc0067,6,C508,0900000,,200924 113137,44108896,1,0,1,0,192.168.32.117,6602#")
}

//$$dc0192,6,V100,0900000,,200924 113137,A0000,37,57,432905999,55,49,240083999,0.00,0,0F0EE331000D7383,0000000000000000,0.00,0.00,0.00,0,0.00,0,0|0.00|0|0|0|0|0|0|0|0.00|0|0,,C508,200924 113137,0,1,1,,0#
