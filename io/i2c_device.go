package io

import (
	"os"
	"syscall"
	"errors"
)

const (
	I2C_SLAVE = 0x0703
)

// var _ I2CDeviceShape = (*I2CDevice)(nil)
type I2CDevice struct {
	Address uint8
	Connected bool
	bus *os.File
}

func (i2c I2CDevice) New() I2CDevice {
	bus, _ := os.OpenFile("/dev/i2c-1", os.O_RDWR, 0600)
	ioctl(bus.Fd(), I2C_SLAVE, uintptr(i2c.Address))
	i2c.bus = bus

	// i2c.Connected = 

	return i2c
}

func (i2c *I2CDevice) read(buf []uint8) (int, error) {
	return i2c.bus.Read(buf)
}

func (i2c *I2CDevice) write(buf []uint8) (int, error) {
	return i2c.bus.Write(buf)
}

func (i2c *I2CDevice) GetBytes(buffer []uint8) (int, error) {
	value, err := i2c.read(buffer)
	if err != nil {
		return 0, errors.New("i2c: failed to get bytes")
	}

	return value, nil
}

func (i2c *I2CDevice) SendBytes(buffer []uint8) error {
	_, err := i2c.write(buffer)
	if err != nil {
		return errors.New("i2c: failed to send bytes")
	}

	return nil
}

func (i2c *I2CDevice) ReadByte(register uint8) (uint8, error) {
	err := i2c.SendBytes([]uint8{register})
	if err != nil {
		return 0, errors.New("i2c: failed to read byte")
	}

	buffer := make([]uint8, 1)
	_, err = i2c.GetBytes(buffer)
	if err != nil {
		return 0, errors.New("i2c: failed to read byte")
	}

	return buffer[0], nil
}

func (i2c *I2CDevice) ReadWord(register uint8) (uint16, error) {
	err := i2c.SendBytes([]uint8{register})
	if err != nil {
		return 0, errors.New("i2c: failed to read word")
	}

	buffer := make([]uint8, 2)
	_, err = i2c.GetBytes(buffer)
	if err != nil {
		return 0, errors.New("i2c: failed to read word")
	}

	result := uint16(buffer[0]) + uint16(buffer[1]) << 8
	return result, nil
}

func (i2c *I2CDevice) WriteByte(register uint8, value uint8) error {
	buffer := []uint8{register, value}
	err := i2c.SendBytes(buffer)
	if err != nil {
		return errors.New("i2c: failed to read word")
	}
	
	return nil
}

func (i2c *I2CDevice) Destroy() {
	i2c.bus.Close()
}

func ioctl(fd, cmd, arg uintptr) error {
	_, _, err := syscall.Syscall6(syscall.SYS_IOCTL, fd, cmd, arg, 0, 0, 0)
	if err != 0 {
		return err
	}
	return nil
}