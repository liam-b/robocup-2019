package helper

import (
	"github.com/liam-b/robocup-2019/bot"
)

const (
	CLAW_SPEED = 400
	CLAW_ELEVATOR_SPEED = 300

	CLAW_DEGREES = 330
	CLAW_RELEASE_DEGREES = 60
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
	bot.ClawMotor.RunToAbsolutePositionAndCoast(CLAW_FUDGE, CLAW_SPEED)
}

func CloseClaw() {
	bot.ClawMotor.RunToAbsolutePositionAndHold(-CLAW_DEGREES, CLAW_SPEED)
}

func ReleaseClaw() {
	bot.ClawMotor.RunToAbsolutePositionAndHold(-CLAW_DEGREES + CLAW_RELEASE_DEGREES, CLAW_SPEED)
}

func RaiseClaw() {
	bot.ClawElevatorMotor.RunToAbsolutePositionAndHold(-CLAW_ELEVATOR_DEGREES, CLAW_ELEVATOR_SPEED)
}

func LowerClaw() {
	bot.ClawElevatorMotor.RunToAbsolutePositionAndCoast(CLAW_FUDGE, CLAW_ELEVATOR_SPEED)
}

func IsClawClosed() bool {
	return bot.ClawMotor.IsStopped()
}

func IsClawRaised() bool {
	return bot.ClawElevatorMotor.IsStopped()
}