package cmsv6

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestParsePacketV101(t *testing.T) {
	v101 := "$$dc0242,1,V101,0900000,,200924 112940,V0000,37,57,432623999,-0,0,0,0.00,0,0000000000007383,0000000000000000,0.00,0.00,0.00,0,0.00,0,0|0.00|0|0|0|0|0|0|0|0.00|0|0,,V1.0.0.1,4108,,0,0,0,SZ88888,1,USER=root,13,1,37,14,0900000,V2018 0414,V6.1.48 20180122,,0,1,0,#"

	r, err := ParsePacket(v101)
	if assert.NoError(t, err) {
		v, ok := r.(*V101)
		if assert.Equal(t, ok, true) {
			assert.Equal(t, v.Type, "V101")
			assert.Equal(t, v.Timestamp, time.Date(2020, time.September, 24, 11, 29,
				40, 0, time.UTC))
		}
	}
}

func TestParsePacketV141(t *testing.T) {
	v141 := "$$dc0146,3,V141,0900000,,200924 112940,V0000,-0,0,0,-0,0,0,0.00,0,0000000000007383,0000000000000000,0.00,0.00,0.00,0,0.00,0,0|0.00|0|0|0|0|0|0|0|0.00|0|0,#"

	r, err := ParsePacket(v141)
	if assert.NoError(t, err) {
		v, ok := r.(*V141)
		if assert.Equal(t, ok, true) {
			assert.Equal(t, v.Type, "V141")
			assert.Equal(t, v.Ack, "")
		}
	}
}

func TestParsePacketV114(t *testing.T) {
	v114 := "$$dc0165,4,V114,0900000,,200924 112942,A0000,37,57,421385999,55,49,237689999,0.00,0,0F0EE331000D7383,0000000000000000,0.00,0.00,0.00,0,0.00,0,0|0.00|0|0|0|0|0|0|0|0.00|0|0,1#"

	r, err := ParsePacket(v114)
	if assert.NoError(t, err) {
		v, ok := r.(*V114)
		if assert.Equal(t, ok, true) {
			assert.Equal(t, v.Type, "V114")
			assert.Equal(t, v.Ack, "1")
		}
	}
}

func TestCreateResponse(t *testing.T) {
	v101 := "$$dc0242,1,V101,0900000,,200924 112940,V0000,37,57,432623999,-0,0,0,0.00,0,0000000000007383,0000000000000000,0.00,0.00,0.00,0,0.00,0,0|0.00|0|0|0|0|0|0|0|0.00|0|0,,V1.0.0.1,4108,,0,0,0,SZ88888,1,USER=root,13,1,37,14,0900000,V2018 0414,V6.1.48 20180122,,0,1,0,#"
	dt := time.Date(2020, time.September, 24, 11, 29, 41, 0, time.UTC)

	r, err := ParsePacket(v101)
	if assert.NoError(t, err) {
		v, ok := r.(*V101)
		if assert.Equal(t, ok, true) {
			r := CreateResponse(v.Header, dt, []string{"0", "1", "1"})
			assert.Equal(t, r.Encode(), "$$dc0056,1,C100,0900000,,200924 112941,V101,200924 112940,0,1,1,#")
		}
	}
}
