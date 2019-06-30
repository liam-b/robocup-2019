package lego

import (
	"github.com/liam-b/robocup-2019/logger"

	"strings"
	"strconv"
	"time"
)

type TachoMotor struct {
	Port PortAddress
	device Device

	buffer map[string]string
	cache map[string]string
	command string
}

func (device TachoMotor) New() TachoMotor {
	device.device = Device{Port: device.Port, Type: MotorDeviceType}.New()
	device.buffer = map[string]string{"speed_sp": "", "position": "", "position_sp": "", "stop_action": ""}
	device.cache = map[string]string{"position": "", "state": "", "speed": ""}
	device.command = ""

	return device
}

func (device *TachoMotor) Setup() {
	err := device.device.Setup()
	if err != nil {
		device.handleError("failed to setup device")
	} else {
		device.SetCommand("reset")
		time.Sleep(time.Millisecond * 100)
	}
}

func (device *TachoMotor) Update() {
	device.setBufferAttributes()
	device.getCacheAttributes()
}

func (device *TachoMotor) SetCommand(command string) {
	device.command = command
}

func (device TachoMotor) SetSpeed(speed int) {
	device.buffer["speed_sp"] = strconv.Itoa(speed)
}

func (device TachoMotor) SetPosition(position int) {
	device.buffer["position"] = strconv.Itoa(position)
}

func (device TachoMotor) GetPosition() int {
	position, _ := strconv.Atoi(strings.ReplaceAll(device.cache["position"], "\n", ""))
	return position
}

func (device TachoMotor) SetTargetPosition(position int) {
	device.buffer["position_sp"] = strconv.Itoa(position)
}

func (device TachoMotor) GetState() []string {
	state, _ := device.cache["state"]
	return strings.Split(state, " ")
}

func (device TachoMotor) GetSpeed() int {
	speed, _ := strconv.Atoi(strings.ReplaceAll(device.cache["speed"], "\n", ""))
	return speed
}

func (device *TachoMotor) SetStopAction(action string) {
	device.buffer["stop_action"] = action
}

func (device *TachoMotor) setBufferAttributes() {
	for attribute, value := range device.buffer {
		if device.buffer[attribute] != "" {
			err := device.device.SetAttribute(attribute, value)
			if err != nil {
				device.handleError("failed to set " + attribute)
			} else {
				device.buffer[attribute] = ""
			}
		}
	}

	if (device.command != "") {
		err := device.device.SetAttribute("command", device.command)
		if err != nil {
			device.handleError("failed to send command")
		} else {
			device.command = ""
		}
	}
}

func (device *TachoMotor) getCacheAttributes() {
	for attribute, _ := range device.cache {
		value, err := device.device.GetAttribute(attribute)
		if err != nil {
			device.handleError("failed to get " + attribute)
		} else {
			device.cache[attribute] = value
		}
	}
}


func (device TachoMotor) handleError(text string) {
	logger.Error("lego tacho motor: " + text + " (" + string(device.Port) + ")")
}