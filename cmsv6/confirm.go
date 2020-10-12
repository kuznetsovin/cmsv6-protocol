package cmsv6

import (
	"fmt"
	"strings"
	"time"
)

type C100 struct {
	Header
	RequestType      string
	RequestTimestamp time.Time
	ExtraFields      []string
}

func (c *C100) Encode() string {
	return fmt.Sprintf("$$%s,%s,%s,%s,#", c.Header.Encode(), c.RequestType,
		c.RequestTimestamp.Format(timestampFmt), strings.Join(c.ExtraFields, ","))
}
