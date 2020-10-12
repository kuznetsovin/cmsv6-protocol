package cmsv6

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestParsePacket(t *testing.T) {
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

func TestCreateResponse(t *testing.T) {
	v101 := "$$dc0242,1,V101,0900000,,200924 112940,V0000,37,57,432623999,-0,0,0,0.00,0,0000000000007383,0000000000000000,0.00,0.00,0.00,0,0.00,0,0|0.00|0|0|0|0|0|0|0|0.00|0|0,,V1.0.0.1,4108,,0,0,0,SZ88888,1,USER=root,13,1,37,14,0900000,V2018 0414,V6.1.48 20180122,,0,1,0,#"
	dt := time.Date(2020, time.September, 24, 11, 29, 41, 0, time.UTC)

	r, err := ParsePacket(v101)
	if assert.NoError(t, err) {
		v, ok := r.(*V101)
		if assert.Equal(t, ok, true) {
			r := CreateResponse(v.Header, dt, []string{"0", "1", "1"})
			assert.Equal(t, r, "$$dc0056,1,C100,0900000,,200924 112941,V101,200924 112940,0,1,1,#")
		}
	}
}
