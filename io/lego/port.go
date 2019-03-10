package lego

const (
	PORT_PATH = "/sys/class/lego-port/"
)

type PortAddress string; const (
	PORT_S1 PortAddress = "spi0.1:S1"
	PORT_S2 PortAddress = "spi0.1:S2"
	PORT_S3 PortAddress = "spi0.1:S3"
	PORT_S4 PortAddress = "spi0.1:S4"

	PORT_MA PortAddress = "spi0.1:MA"
	PORT_MB PortAddress = "spi0.1:MB"
	PORT_MC PortAddress = "spi0.1:MC"
	PORT_MD PortAddress = "spi0.1:MD"
)