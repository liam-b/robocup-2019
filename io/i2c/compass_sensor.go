package i2c

import (
	"github.com/liam-b/robocup-2019/logger"
	"time"
)

const (
	COMPASS_SENSOR_ADDRESS = 0x13

	COMPASS_SENSOR_POWER_REGISTER = 0x4b
	COMPASS_SENSOR_POWER_ON = 0x01

	COMPASS_SENSOR_ENABLE_REGISTER = 0x4c
	COMPASS_SENSOR_ENABLE_ACTIVATE = 0x00

	COMPASS_SENSOR_ROTATION_REGISTER = 0x46
)

type CompassSensor struct {
	Address uint8
	device Device
}

func (sensor CompassSensor) New() CompassSensor {
	sensor.Address = COMPASS_SENSOR_ADDRESS
	sensor.device = Device{Address: sensor.Address}.New()
	return sensor
}

func (sensor CompassSensor) Setup() {
	err := sensor.device.WriteByte(COMPASS_SENSOR_POWER_REGISTER, COMPASS_SENSOR_POWER_ON)
	if err != nil {
		logger.Error("compass sensor: failed to setup sensor")
		return
	}

	time.Sleep(time.Millisecond * 100)
	err = sensor.device.WriteByte(COMPASS_SENSOR_ENABLE_REGISTER, COMPASS_SENSOR_ENABLE_ACTIVATE)
	if err != nil {
		logger.Error("compass sensor: failed to setup sensor")
		return
	}
}

func (sensor CompassSensor) Rotation() int {
	valueLow, err := sensor.device.ReadByte(COMPASS_SENSOR_ROTATION_REGISTER)
	if err != nil {
		logger.Error("compass sensor: failed to read rotation")
		return 0
	}

	// valueHigh := int(sensor.device.ReadByte(COMPASS_SENSOR_ROTATION_REGISTER + 1))

	// return int(valueLow >> 1 + valueHigh << 7)
	return int(valueLow)
}

func (sensor CompassSensor) Destroy() {
	sensor.device.Destroy()
}