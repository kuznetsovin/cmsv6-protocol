package cmsv6

import "time"

type C100 struct {
	Header
	RequestType      string
	RequestTimestamp time.Time
	ExtraFields      []string
}
