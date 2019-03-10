package helper

import (
	"github.com/liam-b/robocup-2019/bot"
)

const (
	CLAW_SPEED = 400
	CLAW_ELEVATOR_SPEED = 300

	CLAW_DEGREES = 315
	CLAW_RELEASE_DEGREES = 30
	CLAW_ELEVATOR_DEGREES = 520
	CLAW_FUDGE = 5
)

var claw = Helper{
	Setup: func() {
		bot.ClawMotor.ResetPosition()
		bot.ClawElevatorMotor.ResetPosition()
	},

	Cleanup: func() {
		OpenClaw()
		LowerClaw()
	},
}

func OpenClaw() {
	bot.ClawMotor.RunToPositionAndBrake(CLAW_FUDGE, CLAW_SPEED)
}

func CloseClaw() {
	bot.ClawMotor.RunToPositionAndHold(-CLAW_DEGREES, CLAW_SPEED)
}

func ReleaseClaw() {
	bot.ClawMotor.RunToPositionAndHold(-CLAW_DEGREES + CLAW_RELEASE_DEGREES, CLAW_SPEED)
}

func RaiseClaw() {
	bot.ClawElevatorMotor.RunToPositionAndHold(-CLAW_ELEVATOR_DEGREES, CLAW_ELEVATOR_SPEED)
}

func LowerClaw() {
	bot.ClawElevatorMotor.RunToPositionAndCoast(CLAW_FUDGE, CLAW_ELEVATOR_SPEED)
}