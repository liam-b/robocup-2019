package lego

import "strings"

type Motor struct {
	Port PortAddress
	device TachoMotor
}

func (motor Motor) New() Motor {
	motor.device = TachoMotor{Port: motor.Port}.New()
	return motor
}

func (motor *Motor) Setup() {
	motor.device.Setup()
}

func (motor *Motor) Update() {
	motor.device.Update()
}

func (motor *Motor) Cleanup() {
	motor.device.SetStopAction("coast")
	motor.device.Update()
}

func (motor *Motor) Run(speed int) {
	motor.device.SetSpeed(speed)
	motor.device.SetCommand("run-forever")
}

func (motor *Motor) RunToAbsolutePositionAndBrake(position int, speed int) {
	motor.runToPosition("run-to-abs-pos", position, speed, "brake")
}

func (motor *Motor) RunToAbsolutePositionAndCoast(position int, speed int) {
	motor.runToPosition("run-to-abs-pos", position, speed, "coast")
}

func (motor *Motor) RunToAbsolutePositionAndHold(position int, speed int) {
	motor.runToPosition("run-to-abs-pos", position, speed, "hold")
}

func (motor *Motor) RunToRelativePositionAndBrake(position int, speed int) {
	motor.runToPosition("run-to-rel-pos", position, speed, "brake")
}

func (motor *Motor) RunToRelativePositionAndCoast(position int, speed int) {
	motor.runToPosition("run-to-rel-pos", position, speed, "coast")
}

func (motor *Motor) RunToRelativePositionAndHold(position int, speed int) {
	motor.runToPosition("run-to-rel-pos", position, speed, "hold")
}

func (motor *Motor) Brake() {
	motor.device.SetStopAction("brake")
	motor.device.SetCommand("stop")
}

func (motor *Motor) Coast() {
	motor.device.SetStopAction("coast")
	motor.device.SetCommand("stop")
}

func (motor *Motor) Hold() {
	motor.device.SetStopAction("hold")
	motor.device.SetCommand("stop")
}

func (motor *Motor) Position() int {
	return motor.device.GetPosition()
}

func (motor *Motor) ResetPosition() {
	motor.device.SetPosition(0)
}

func (motor *Motor) State() []string {
	return motor.device.GetState()
}

func (motor *Motor) StateContains(search string) bool {
	for _, state := range motor.device.GetState() {
		state = strings.ReplaceAll(state, "\n", "")
		if state == search {
			return true
		}
	}
	
	return false
}

func (motor *Motor) runToPosition(command string, position int, speed int, stopAction string) {
	motor.device.SetTargetPosition(position)
	motor.device.SetSpeed(speed)
	motor.device.SetStopAction(stopAction)
	motor.device.SetCommand(command)
}