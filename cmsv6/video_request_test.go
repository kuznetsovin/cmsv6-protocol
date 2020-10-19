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

	assert.Equal(t, c.Encode(), "$$dc0067,6,C508,0900000,,200924 113137,44108896,1,0,1,0,192.168.32.117,6602#")
}

func TestCreateVideoRequest(t *testing.T) {
	dt := time.Date(2020, time.September, 24, 11, 31, 37, 0, time.UTC)
	r := CreateVideoRequest(dt, "0900000")

	c := r.(*C508)
	c.PacketNumber = 6
	c.SrvIp = "192.168.32.117"
	assert.Equal(t, c.Encode(),
		"$$dc0067,6,C508,0900000,,200924 113137,44108896,1,0,1,0,192.168.32.117,6602#")
}

func TestV100_Decode(t *testing.T) {
	cmd := "$$dc0192,6,V100,0900000,,200924 113137,A0000,37,57,432905999,55,49,240083999,0.00,0,0F0EE331000D7383,0000000000000000,0.00,0.00,0.00,0,0.00,0,0|0.00|0|0|0|0|0|0|0|0.00|0|0,,C508,200924 113137,0,1,1,,0#"
	m, err := parseMsg(cmd)
	if !assert.NoError(t, err) {
		return
	}

	v := V100{}
	if assert.NoError(t, v.Decode(m)) {
		assert.Equal(t, v.MessageID, "dc0192")
		assert.Equal(t, v.Type, "V100")
		assert.Equal(t, v.Timestamp, time.Date(2020, time.September, 24, 11, 31, 37, 0, time.UTC))
		assert.Equal(t, v.Latitude, 37.962025166638895)
		assert.Equal(t, v.RequestType, "C508")
		assert.Equal(t, v.RequestTimestamp, time.Date(2020, time.September, 24, 11, 31, 37, 0, time.UTC))
		assert.Equal(t, len(v.ExtraFields), 5)
	}
}
