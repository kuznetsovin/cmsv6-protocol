package cmsv6

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestV101_Decode(t *testing.T) {
	cmd := "$$dc0242,1,V101,0900000,,200924 112940,V0000,37,57,432623999,-0,0,0,0.00,0,0000000000007383,0000000000000000,0.00,0.00,0.00,0,0.00,0,0|0.00|0|0|0|0|0|0|0|0.00|0|0,,V1.0.0.1,4108,,0,0,0,SZ88888,1,USER=root,13,1,37,14,0900000,V2018 0414,V6.1.48 20180122,,0,1,0,#"
	m, err := parseMsg(cmd)
	if !assert.NoError(t, err) {
		return
	}

	v := V101{}
	if assert.NoError(t, v.Decode(m)) {
		assert.Equal(t, v.MessageID, "dc0242")
		assert.Equal(t, v.Type, "V101")
		assert.Equal(t, v.Timestamp, time.Date(2020, time.September, 24, 11, 29, 40, 0, time.UTC))
		assert.Equal(t, v.Latitude, 37.96201733330556)
		assert.Equal(t, len(v.AuthInfo), 22)

	}
}

func TestV141_Decode(t *testing.T) {
	cmd := "$$dc0146,3,V141,0900000,,200924 112940,V0000,-0,0,0,-0,0,0,0.00,0,0000000000007383,0000000000000000,0.00,0.00,0.00,0,0.00,0,0|0.00|0|0|0|0|0|0|0|0.00|0|0,#"
	m, err := parseMsg(cmd)
	if !assert.NoError(t, err) {
		return
	}

	v := V141{}
	if assert.NoError(t, v.Decode(m)) {
		assert.Equal(t, v.MessageID, "dc0146")
		assert.Equal(t, v.Type, "V141")
		assert.Equal(t, len(v.UnknownFields), 11)
		assert.Equal(t, v.Ack, "")
	}
}
