package i2c

import (
	"github.com/liam-b/robocup-2019/logger"
	"math"
)

const (
	COLOR_SENSOR_ADDRESS = 0x29

	COLOR_SENSOR_ENABLE_REGISTER = 0x80
	COLOR_SENSOR_TIMING_REGISTER = 0x81

	COLOR_SENSOR_CLEAR_REGISTER = 0x94
	COLOR_SENSOR_RED_REGISTER = 0x96
	COLOR_SENSOR_GREEN_REGISTER = 0x98
	COLOR_SENSOR_BLUE_REGISTER = 0x9a

	// 5:17 ratio
	COLOR_SENSOR_CLEAR_SCALE = 0.1
	COLOR_SENSOR_COLOR_SCALE = 0.34
)

type ColorSensor struct {
	Address uint8
	Multiplexer *Multiplexer
	Channel uint
	device MultiplexedDevice

	cachedClearValue int
	cachedRedValues int
	cachedGreenValues int
	cachedBlueValues int
}

func (sensor ColorSensor) New() ColorSensor {
	sensor.Address = COLOR_SENSOR_ADDRESS
	sensor.device = MultiplexedDevice{Address: sensor.Address, Multiplexer: sensor.Multiplexer, Channel: sensor.Channel}.New()
	return sensor
}

func (sensor ColorSensor) Setup() {
	err := sensor.device.WriteByte(COLOR_SENSOR_ENABLE_REGISTER, 0x03)
	if err != nil {
		logger.Print("color sensor: failed to setup sensor")
		return
	}

	err = sensor.device.WriteByte(COLOR_SENSOR_TIMING_REGISTER, 0xf3)
	if err != nil {
		logger.Print("color sensor: failed to setup sensor")
	}
}

func (sensor *ColorSensor) Update() {
	sensor.cachedClearValue = sensor.getClearValue(COLOR_SENSOR_CLEAR_REGISTER)
	sensor.cachedRedValues = sensor.getColorValue(COLOR_SENSOR_RED_REGISTER)
	sensor.cachedGreenValues = sensor.getColorValue(COLOR_SENSOR_GREEN_REGISTER)
	sensor.cachedBlueValues = sensor.getColorValue(COLOR_SENSOR_BLUE_REGISTER)
}

func (sensor ColorSensor) Cleanup() {
	err := sensor.device.WriteByte(COLOR_SENSOR_ENABLE_REGISTER, 0x00)
	if err != nil {
		logger.Print("color sensor: failed to cleanup sensor")
	}
	sensor.device.Destroy()
}

func (sensor ColorSensor) Intensity() int {
	return sensor.cachedClearValue
}

func (sensor ColorSensor) RGB() (int, int, int) {
	red := sensor.cachedRedValues
	green := sensor.cachedGreenValues
	blue := sensor.cachedBlueValues

	return red, green, blue
}

func (sensor ColorSensor) getClearValue(register uint8) int {
	data, err := sensor.device.ReadWord(register)
	if err != nil {
		logger.Print("color sensor: failed to read clear value")
		return 0
	}

	value := float64(data) * COLOR_SENSOR_CLEAR_SCALE
	value = math.Max(0, math.Min(value, 100))
	return int(value)
}

func (sensor ColorSensor) getColorValue(register uint8) int {
	data, err := sensor.device.ReadWord(register)
	if err != nil {
		logger.Print("color sensor: failed to read color value")
		return 0
	}

	value := float64(data) * COLOR_SENSOR_COLOR_SCALE
	value = math.Max(0, math.Min(value, 100))
	return int(value)
}