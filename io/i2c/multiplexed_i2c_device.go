package i2c

type MultiplexedDevice struct {
	Address uint8
	Multiplexer *Multiplexer
	Channel uint
	device Device
}

func (i2c MultiplexedDevice) New() MultiplexedDevice {
	i2c.device = Device{Address: i2c.Address}.New()
	return i2c
}

func (i2c *MultiplexedDevice) GetBytes(buffer []uint8) (int, error) {
	i2c.updateMultiplexerChannel()
	return i2c.device.GetBytes(buffer)
}

func (i2c *MultiplexedDevice) SendBytes(buffer []uint8) error {
	i2c.updateMultiplexerChannel()
	return i2c.device.SendBytes(buffer)
}

func (i2c *MultiplexedDevice) ReadByte(register uint8) (uint8, error) {
	i2c.updateMultiplexerChannel()
	return i2c.device.ReadByte(register)
}

func (i2c *MultiplexedDevice) ReadWord(register uint8) (uint16, error) {
	i2c.updateMultiplexerChannel()
	return i2c.device.ReadWord(register)
}

func (i2c *MultiplexedDevice) WriteByte(register uint8, value uint8) error {
	i2c.updateMultiplexerChannel()
	return i2c.device.WriteByte(register, value)
}

func (i2c MultiplexedDevice) Destroy() {
	i2c.device.Destroy()
}

func (i2c *MultiplexedDevice) updateMultiplexerChannel() {
	if i2c.Multiplexer.Channel != i2c.Channel {
		i2c.Multiplexer.SetChannel(i2c.Channel)
	}
}