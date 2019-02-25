package io

// var _ I2CDeviceShape = (*MultiplexedI2CDevice)(nil)
type MultiplexedI2CDevice struct {
	Address uint8
	Multiplexer *Multiplexer
	Channel uint
	i2cDevice I2CDevice
}

func (i2c MultiplexedI2CDevice) New() MultiplexedI2CDevice {
	i2c.i2cDevice = I2CDevice{Address: i2c.Address}.New()
	return i2c
}

func (i2c *MultiplexedI2CDevice) GetBytes(buffer []uint8) (int, error) {
	i2c.updateMultiplexerChannel()
	return i2c.i2cDevice.GetBytes(buffer)
}

func (i2c *MultiplexedI2CDevice) SendBytes(buffer []uint8) error {
	i2c.updateMultiplexerChannel()
	return i2c.i2cDevice.SendBytes(buffer)
}

func (i2c *MultiplexedI2CDevice) ReadByte(register uint8) (uint8, error) {
	i2c.updateMultiplexerChannel()
	return i2c.i2cDevice.ReadByte(register)
}

func (i2c *MultiplexedI2CDevice) ReadWord(register uint8) (uint16, error) {
	i2c.updateMultiplexerChannel()
	return i2c.i2cDevice.ReadWord(register)
}

func (i2c *MultiplexedI2CDevice) WriteByte(register uint8, value uint8) error {
	i2c.updateMultiplexerChannel()
	return i2c.i2cDevice.WriteByte(register, value)
}

func (i2c MultiplexedI2CDevice) Destroy() {
	i2c.i2cDevice.Destroy()
}

func (i2c *MultiplexedI2CDevice) updateMultiplexerChannel() {
	if i2c.Multiplexer.Channel != i2c.Channel {
		i2c.Multiplexer.SetChannel(i2c.Channel)
	}
}