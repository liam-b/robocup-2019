package lego

import (
	"errors"
	"strings"
	"strconv"
)

type DeviceType int; const (
	MotorDeviceType DeviceType = iota
)

var PATHS = map[DeviceType]string{
	MotorDeviceType: "/sys/class/tacho-motor/",
}

type Device struct {
	Port PortAddress
	Type DeviceType

	path string
}

func (device Device) New() Device {
	return device
}

func (device *Device) Setup() error {
	path := PATHS[device.Type]
	devices, err := ListFiles(PATHS[device.Type])

	if err != nil {
		return errors.New("lego: failed to read ports")
	}

	for _, dev := range devices {
		address, err := ReadFile(path + dev + "/address")
		if err != nil {
			return errors.New("lego: failed to read address from port " + dev)
		}

		if strings.Contains(address, string(device.Port)) {
			device.path = path + dev + "/"
			break
		}
	}

	if device.path == "" {
		return errors.New("lego: failed to find device in port " + string(device.Port))
	}
	return nil
}

func (device Device) SetAttribute(attribute string, value string) error {
	return WriteFile(device.path + attribute, value)
}

func (device Device) GetAttribute(attribute string) (string, error) {
	return ReadFile(device.path + attribute)
}