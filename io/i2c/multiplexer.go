package i2c

import "github.com/liam-b/robocup-2019/logger"

const (
	I2C_MULTIPLEXER_ADDRESS = 0x71
)

type Multiplexer struct {
	Address uint8
	Channel uint
	device Device
}

func (i2c Multiplexer) New() Multiplexer {
	i2c.Address = I2C_MULTIPLEXER_ADDRESS
	i2c.device = Device{Address: i2c.Address}.New()
	return i2c
}

func (i2c *Multiplexer) Setup() {
	i2c.SetChannel(0)
}

func (i2c *Multiplexer) SetChannel(index uint) {
	i2c.Channel = index
	channel := uint(1 << index)
	err := i2c.device.SendBytes([]uint8{uint8(channel)})
	if err != nil {
		logger.Error("multiplexer: failed to set channel")
	}
}

func (i2c Multiplexer) Destroy() {
	i2c.device.Destroy()
}