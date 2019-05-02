package lego

import (
	"github.com/liam-b/robocup-2019/logger"

	"strings"
	"strconv"
	"time"
	"sync"
)

type TachoMotor struct {
	Port PortAddress
	device Device

	buffer map[string]string
	cache map[string]string
	mutex *sync.Mutex
}

func (device TachoMotor) New() TachoMotor {
	device.device = Device{Port: device.Port, Type: MotorDeviceType}.New()
	device.buffer = map[string]string{"command": "", "speed_sp": "", "position": "", "position_sp": "", "stop_action": ""}
	device.cache = map[string]string{"position": "", "state": ""}
	device.mutex = &sync.Mutex{}

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

func (device TachoMotor) SetCommand(command string) {
	device.mutex.Lock()
	device.buffer["command"] = command
	device.mutex.Unlock()
}

func (device TachoMotor) SetSpeed(speed int) {
	device.mutex.Lock()
	device.buffer["speed_sp"] = strconv.Itoa(speed)
	device.mutex.Unlock()
}

func (device TachoMotor) SetPosition(position int) {
	device.mutex.Lock()
	device.buffer["position"] = strconv.Itoa(position)
	device.mutex.Unlock()
}

func (device TachoMotor) GetPosition() int {
	position, _ := strconv.Atoi(device.cache["position"])
	return position
}

func (device TachoMotor) SetTargetPosition(position int) {
	device.mutex.Lock()
	device.buffer["position_sp"] = strconv.Itoa(position)
	device.mutex.Unlock()
}

func (device TachoMotor) GetState() []string {
	state, _ := device.cache["state"]
	return strings.Split(state, " ")
}

func (device *TachoMotor) SetStopAction(action string) {
	device.mutex.Lock()
	device.buffer["stop_action"] = action
	device.mutex.Unlock()
}

func (device *TachoMotor) setBufferAttributes() {
	device.mutex.Lock()
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
	device.mutex.Unlock()
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