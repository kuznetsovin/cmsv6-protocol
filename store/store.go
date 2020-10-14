package store

import (
	"time"
)

type Store interface {
	Save(point GeoPoint) error
}

type GeoPoint struct {
	DeviceID string
	NavTime  time.Time
	Lat      float64
	Lon      float64
}

func NewStore(conn string) Store {
	var (
		result Store
	)

	if conn == "" {
		result = &DefaultStore{}
	}

	return result
}
