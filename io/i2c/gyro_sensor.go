package i2c

import (
	"github.com/liam-b/robocup-2019/logger"
)

const (
	GRYRO_SENSOR_ADDRESS = 0x68
	GYRO_SENSOR_ROTATION_REGISTER = 0x48
)

type GyroSensor struct {
	Address uint8
	device Device

	cachedRotation int
}

func (sensor GyroSensor) New() GyroSensor {
	sensor.Address = GRYRO_SENSOR_ADDRESS
	sensor.device = Device{Address: sensor.Address}.New()
	return sensor
}

func (sensor GyroSensor) Setup() {}

func (sensor GyroSensor) Update() {
	sensor.cachedRotation += sensor.getRotation()
}

func (sensor GyroSensor) Rotation() int {
	return sensor.cachedRotation
}

func (sensor GyroSensor) Destroy() {
	sensor.device.Destroy()
}

func (sensor GyroSensor) getRotation() int {
	valueLow, err := sensor.device.ReadByte(GYRO_SENSOR_ROTATION_REGISTER)
	if err != nil {
		logger.Error("gyro sensor: failed to read rotation")
		return 0
	}

	valueHigh, err := sensor.device.ReadByte(GYRO_SENSOR_ROTATION_REGISTER - 1)
	if err != nil {
		logger.Error("gyro sensor: failed to read rotation")
		return 0
	}

	return int(valueLow) + int(valueHigh) << 8
}