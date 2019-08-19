package i2c

import (
	"github.com/liam-b/robocup-2019/logger"
)

const (
	ULTRASONIC_SENSOR_ADDRESS = 0x70

	ULTRASONIC_SENSOR_COMMAND_REGISTER = 0x00
	ULTRASONIC_SENSOR_SOFTWARE_REVISION_REGISTER = 0x00
	ULTRASONIC_SENSOR_DISTANCE_REGISTER = 0x03

	ULTRASONIC_SENSOR_RANGING_INCHES = 0x50
	ULTRASONIC_SENSOR_RANGING_CENTIMETERS = 0x51
	ULTRASONIC_SENSOR_RANGING_MICROSECONDS = 0x52
)

type UltrasonicSensor struct {
	Address uint8
	device Device

	cachedDistance int
}

func (sensor UltrasonicSensor) New() UltrasonicSensor {
	sensor.Address = ULTRASONIC_SENSOR_ADDRESS
	sensor.device = Device{Address: sensor.Address}.New()
	return sensor
}

func (sensor UltrasonicSensor) Setup() {
	err := sensor.device.WriteByte(ULTRASONIC_SENSOR_COMMAND_REGISTER, ULTRASONIC_SENSOR_RANGING_CENTIMETERS)
	if err != nil {
		logger.Print("ultrasonic sensor: failed to setup sensor")
	}
}

func (sensor *UltrasonicSensor) Update() {
	response, err := sensor.device.ReadByte(ULTRASONIC_SENSOR_SOFTWARE_REVISION_REGISTER)
	if err != nil {
		logger.Print("ultrasonic sensor: failed to communicate with sensor")
	} else {
		if response != 0xff {
			data, err := sensor.device.ReadWord(ULTRASONIC_SENSOR_DISTANCE_REGISTER)
			if err != nil {
				logger.Print("ultrasonic sensor: failed to read distance value")
			} else {
				sensor.cachedDistance = int(data)
				err := sensor.device.WriteByte(ULTRASONIC_SENSOR_COMMAND_REGISTER, ULTRASONIC_SENSOR_RANGING_CENTIMETERS)
				if err != nil {
					logger.Print("ultrasonic sensor: failed to send ranging command")
				}
			}
		}
	}
}

func (sensor UltrasonicSensor) Cleanup() {}

func (sensor UltrasonicSensor) Distance() int {
	return sensor.cachedDistance
}

func (sensor UltrasonicSensor) Destroy() {
	sensor.device.Destroy()
}