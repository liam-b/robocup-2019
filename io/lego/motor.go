package lego

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

func (motor Motor) Run(speed int) {
	motor.device.SetSpeed(speed)
	motor.device.SetCommand("run-forever")
}

func (motor Motor) RunToPositionAndBrake(position int, speed int) {
	motor.runToPosition(position, speed, "brake")
}

func (motor Motor) RunToPositionAndCoast(position int, speed int) {
	motor.runToPosition(position, speed, "coast")
}

func (motor Motor) RunToPositionAndHold(position int, speed int) {
	motor.runToPosition(position, speed, "hold")
}

func (motor Motor) Brake() {
	motor.device.SetStopAction("brake")
	motor.device.SetCommand("stop")
}

func (motor Motor) Coast() {
	motor.device.SetStopAction("coast")
	motor.device.SetCommand("stop")
}

func (motor Motor) Hold() {
	motor.device.SetStopAction("hold")
	motor.device.SetCommand("stop")
}

func (motor Motor) Position() int {
	return motor.device.GetPosition()
}

func (motor Motor) ResetPosition() {
	motor.device.SetPosition(0)
}

func (motor Motor) runToPosition(position int, speed int, stopAction string) {
	motor.device.SetTargetPosition(position)
	motor.device.SetSpeed(speed)
	motor.device.SetStopAction(stopAction)
	motor.device.SetCommand("run-to-abs-pos")
}