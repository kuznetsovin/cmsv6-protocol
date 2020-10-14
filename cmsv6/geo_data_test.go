package cmsv6

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestV114_Decode(t *testing.T) {
	cmd := "$$dc0165,4,V114,0900000,,200924 112942,A0000,37,57,421385999,55,49,237689999,0.00,0,0F0EE331000D7383,0000000000000000,0.00,0.00,0.00,0,0.00,0,0|0.00|0|0|0|0|0|0|0|0.00|0|0,1#"
	m, err := parseMsg(cmd)
	if !assert.NoError(t, err) {
		return
	}

	v := V114{}
	if assert.NoError(t, v.Decode(m)) {
		assert.Equal(t, v.Type, "V114")
		assert.Equal(t, v.Ack, "1")
	}
}
