package lego

import (
	"github.com/liam-b/robocup-2019/logger"
	"time"
)

type MotorDevice struct {
	Port PortAddress

	device Device
}

func (device MotorDevice) New() MotorDevice {
	device.device = Device{Port: device.Port, Type: MotorDeviceType}.New()
	return device
}

func (device *MotorDevice) Setup() {
	err := device.device.Setup()
	if err != nil {
		device.handleError("failed to setup device")
	} else {
		device.SetCommand("reset")
		time.Sleep(time.Millisecond * 100)
	}
}

func (device MotorDevice) SetCommand(command string) {
	err := device.device.SetStringAttribute("command", command)
	if err != nil {
		device.handleError("failed to send command")
	}
}

func (device MotorDevice) SetSpeed(speed int) {
	err := device.device.SetIntAttribute("speed_sp", speed)
	if err != nil {
		device.handleError("failed to set speed")
	}
}

func (device MotorDevice) SetPosition(position int) {
	err := device.device.SetIntAttribute("position", position)
	if err != nil {
		device.handleError("failed to set position")
	}
}

func (device MotorDevice) GetPosition() int {
	position, err := device.device.GetIntAttribute("position")
	if err != nil {
		device.handleError("failed to get position")
	}

	return position
}

func (device MotorDevice) SetTargetPosition(position int) {
	err := device.device.SetIntAttribute("position_sp", position)
	if err != nil {
		device.handleError("failed to set target position")
	}
}

func (device MotorDevice) GetState() []string {
	state, err := device.device.GetStringArrayAttribute("state")
	if err != nil {
		device.handleError("failed to get state")
	}

	return state
}

func (device MotorDevice) SetStopAction(action string) {
	err := device.device.SetStringAttribute("stop_action", action)
	if err != nil {
		device.handleError("failed to set stop action")
	}
}

func (device MotorDevice) handleError(text string) {
	logger.Error("lego motor: " + text + " (" + string(device.Port) + ")")
}