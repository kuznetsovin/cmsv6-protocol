package device_manager

import (
	"errors"
	"net"
	"sync"
)

type DeviceRegistry struct {
	registry map[string]net.Conn
	m        sync.RWMutex
}

func (r *DeviceRegistry) AddDevice(deviceId string, c net.Conn) {
	r.m.RLock()
	r.registry[deviceId] = c
	r.m.RUnlock()
}

func (r *DeviceRegistry) SendCommand(deviceId, command string) error {
	devConn, ok := r.registry[deviceId]
	if !ok {
		return errors.New("Device not found in registry")
	}

	_, err := devConn.Write([]byte(command))

	return err
}

func (r *DeviceRegistry) RemoveDevice(deviceId string) {
	r.m.RLock()
	delete(r.registry, deviceId)
	r.m.RUnlock()
}

func NewDeviceRegistry() DeviceRegistry {
	result := DeviceRegistry{}

	result.registry = make(map[string]net.Conn)

	return result
}
