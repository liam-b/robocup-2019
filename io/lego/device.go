package lego

import (
	"errors"
	"strings"
)

type DeviceType int

const (
	Motor DeviceType = iota
)

var PATHS = map[DeviceType]string{
	Motor: "/sys/class/tacho-motor/",
}

type Device struct {
	Port PortAddress
	Type DeviceType

	path string
}

func (device Device) New() Device {
	return device
}

func (device Device) Setup() error {
	path := PATHS[device.Type]
	ports, err := ListFiles(PATHS[device.Type])

	if err != nil {
		return errors.New("lego: failed to read ports")
	}

	for _, port := range ports {
		address, err := ReadFile(path + port + "/address")
		if err != nil {
			return errors.New("lego: failed to read port " + port)
		}

		if strings.Contains(address, string(device.Port)) {
			device.path = path + port + "/"
			break
		}
	}

	if device.path == "" {
		return errors.New("lego: failed to find device in port " + string(device.Port))
	}
	return nil
}
