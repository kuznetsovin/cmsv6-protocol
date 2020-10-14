package store

import "fmt"

type DefaultStore struct {
}

func (s *DefaultStore) Save(p GeoPoint) error {
	fmt.Printf("%+v", p)
	return nil
}
