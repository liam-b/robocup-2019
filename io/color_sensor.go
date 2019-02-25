package io

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
	i2cDevice MultiplexedI2CDevice
}

func (sensor ColorSensor) New() ColorSensor {
	sensor.Address = COLOR_SENSOR_ADDRESS
	sensor.i2cDevice = MultiplexedI2CDevice{Address: sensor.Address, Multiplexer: sensor.Multiplexer, Channel: sensor.Channel}.New()
	return sensor
}

func (sensor ColorSensor) Setup() {
	err := sensor.i2cDevice.WriteByte(COLOR_SENSOR_ENABLE_REGISTER, 0x03)
	if err != nil {
		logger.Error("color sensor: failed to setup sensor")
	}

	err = sensor.i2cDevice.WriteByte(COLOR_SENSOR_TIMING_REGISTER, 0xf2)
	if err != nil {
		logger.Error("color sensor: failed to setup sensor")
	}
}

func (sensor ColorSensor) Intensity() int {
	return sensor.readClearValue(COLOR_SENSOR_CLEAR_REGISTER)
}

func (sensor ColorSensor) RGB() (int, int, int) {
	red := sensor.readColorValue(COLOR_SENSOR_RED_REGISTER)
	green := sensor.readColorValue(COLOR_SENSOR_GREEN_REGISTER)
	blue := sensor.readColorValue(COLOR_SENSOR_BLUE_REGISTER)

	return int(red), int(green), int(blue)
}

func (sensor ColorSensor) Destroy() {
	sensor.i2cDevice.Destroy()
}

func (sensor ColorSensor) readClearValue(register uint8) int {
	data, err := sensor.i2cDevice.ReadWord(register)
	if err != nil {
		logger.Error("color sensor: failed to read clear value")
	}

	value := float64(data) * COLOR_SENSOR_CLEAR_SCALE
	value = math.Max(0, math.Min(value, 100))
	return int(value)
}

func (sensor ColorSensor) readColorValue(register uint8) int {
	data, err := sensor.i2cDevice.ReadWord(register)
	if err != nil {
		logger.Error("color sensor: failed to read color value")
	}

	value := float64(data) * COLOR_SENSOR_COLOR_SCALE
	value = math.Max(0, math.Min(value, 100))
	return int(value)
}