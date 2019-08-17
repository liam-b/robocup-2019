package behaviour

import (
	"github.com/liam-b/robocup-2019/bot"
	"github.com/liam-b/robocup-2019/helper"
	"github.com/liam-b/robocup-2019/logger"
)

var (
	GREEN_TURN_JUNCTION_SPEED = 200
	GREEN_TURN_JUNCTION_INTENSNITY = 0.3
	GREEN_TURN_DOUBLE_GREEN_INTENSITY = 11
	GREEN_TURN_DOUBLE_GREEN_TIME_LIMIT = bot.Time(200)

	GREEN_TURN_INNER_SPEED = -150
	GREEN_TURN_OUTER_SPEED = 100
	GREEN_TURN_MIN_TURN_DEGREES = 100 // TEST

	GREEN_TURN_COOLDOWN = bot.Time(800)
)

func GreenTurnLeft() {
	bot.DriveMotorLeft.Run(GREEN_TURN_JUNCTION_SPEED)
	bot.DriveMotorRight.Run(GREEN_TURN_JUNCTION_SPEED)

	logger.Debug("gren hello from left turn")

	doubleGreenCount := 0
	for helper.LeftError() > GREEN_TURN_JUNCTION_INTENSNITY {
		if helper.LeftGreen() > GREEN_TURN_DOUBLE_GREEN_INTENSITY && helper.RightGreen() > GREEN_TURN_DOUBLE_GREEN_INTENSITY {
			// doubleGreenCount += 1
		}

		if doubleGreenCount > GREEN_TURN_DOUBLE_GREEN_TIME_LIMIT {
			ChemicalSpillVerify() // FIX ME!!!
			return
		}

		bot.CycleDelay() 
	}

	bot.DriveMotorLeft.ResetPosition()
	bot.DriveMotorLeft.Run(GREEN_TURN_INNER_SPEED)
	bot.DriveMotorRight.Run(GREEN_TURN_OUTER_SPEED)
	for helper.MiddleError() > GREEN_TURN_JUNCTION_INTENSNITY && bot.DriveMotorLeft.Position() < GREEN_TURN_MIN_TURN_DEGREES { bot.CycleDelay() }

	GreenTurnCooldown()
}

func GreenTurnRight() {
	bot.DriveMotorLeft.Run(GREEN_TURN_JUNCTION_SPEED)
	bot.DriveMotorRight.Run(GREEN_TURN_JUNCTION_SPEED)

	doubleGreenCount := 0
	for helper.RightError() > GREEN_TURN_JUNCTION_INTENSNITY {
		if helper.LeftGreen() > GREEN_TURN_DOUBLE_GREEN_INTENSITY && helper.RightGreen() > GREEN_TURN_DOUBLE_GREEN_INTENSITY {
			// doubleGreenCount += 1
		}

		if doubleGreenCount > bot.Time(GREEN_TURN_DOUBLE_GREEN_TIME_LIMIT) {
			ChemicalSpillVerify() // FIX ME!!!
			return
		}

		bot.CycleDelay() 
	}

	bot.DriveMotorRight.ResetPosition()
	bot.DriveMotorLeft.Run(GREEN_TURN_OUTER_SPEED)
	bot.DriveMotorRight.Run(GREEN_TURN_INNER_SPEED)
	for helper.MiddleError() > GREEN_TURN_JUNCTION_INTENSNITY && bot.DriveMotorRight.Position() < GREEN_TURN_MIN_TURN_DEGREES { bot.CycleDelay() }

	GreenTurnCooldown()
}

func GreenTurnCooldown() {
	helper.ResetPID()

	for i := 0; i < GREEN_TURN_COOLDOWN; i++ {
		left, right := helper.PID()
		bot.DriveMotorLeft.Run(left)
		bot.DriveMotorRight.Run(right)

		bot.CycleDelay()
	}
}
