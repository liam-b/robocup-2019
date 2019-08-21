package behaviour

import (
	"github.com/liam-b/robocup-2019/bot"
	"github.com/liam-b/robocup-2019/helper"
	"github.com/liam-b/robocup-2019/logger"
)

var (
	FOLLOW_LINE_GREEN_TURN_THRESHOLD = 12
	FOLLOW_LINE_GREEN_TURN_TIME_LIMIT = bot.Time(50)

	FOLLOW_LINE_RECOVER_LOST_THRESHOLD = 35
	FOLLOW_LINE_RECOVER_FOUND_THRESHOLD = 15
	FOLLOW_LINE_RECOVER_LOST_TIME_LIMIT = bot.Time(300)
	FOLLOW_LINE_RECOVER_REVERSE_SPEED = 150
	FOLLOW_LINE_RECOVER_REVERSE_POSITION_LIMIT = 180
)

func FollowLine() {
	logger.Print("started following line")
	helper.ResetPID()

	leftGreenCount := 0
	rightGreenCount := 0
	lineLostCount := 0
	for {
		left, right := helper.PID()
		bot.DriveMotorLeft.Run(left)
		bot.DriveMotorRight.Run(right)

		if helper.LeftGreen() > FOLLOW_LINE_GREEN_TURN_THRESHOLD {
			leftGreenCount++
			if leftGreenCount > FOLLOW_LINE_GREEN_TURN_TIME_LIMIT {
				GreenTurnLeft()
				leftGreenCount = 0
			}
		} else {
			leftGreenCount /= 2
		}

		if helper.RightGreen() > FOLLOW_LINE_GREEN_TURN_THRESHOLD {
			rightGreenCount++
			if rightGreenCount > FOLLOW_LINE_GREEN_TURN_TIME_LIMIT {
				GreenTurnRight()
				rightGreenCount = 0
			}
		} else {
			rightGreenCount /= 2
		}

		if bot.ColorSensorLeft.Intensity() > FOLLOW_LINE_RECOVER_LOST_THRESHOLD && bot.ColorSensorRight.Intensity() > FOLLOW_LINE_RECOVER_LOST_THRESHOLD {
			lineLostCount++
			if lineLostCount > FOLLOW_LINE_RECOVER_LOST_TIME_LIMIT {
				FollowLineRecoverLine()
				lineLostCount = 0
			}
		} else {
			lineLostCount /= 2
		}

		bot.CycleDelay()
	}
}

func FollowLineRecoverLine() {
	logger.Print("lost line during line following, now reversing")

	bot.DriveMotorLeft.ResetPosition()
	bot.DriveMotorLeft.Run(-FOLLOW_LINE_RECOVER_REVERSE_SPEED)
	bot.DriveMotorRight.Run(-FOLLOW_LINE_RECOVER_REVERSE_SPEED)

	leftFoundPosition := 0
	rightFoundPosition := 0
	for bot.DriveMotorLeft.Position() > -FOLLOW_LINE_RECOVER_REVERSE_POSITION_LIMIT {
		if bot.ColorSensorLeft.Intensity() < FOLLOW_LINE_RECOVER_FOUND_THRESHOLD {
			leftFoundPosition = bot.DriveMotorLeft.Position()
		}

		if bot.ColorSensorLeft.Intensity() < FOLLOW_LINE_RECOVER_FOUND_THRESHOLD {
			rightFoundPosition = bot.DriveMotorLeft.Position()
		}

		bot.CycleDelay()
	}

	targetPosition := 0
	if leftFoundPosition == 0 && rightFoundPosition == 0 {
		logger.Print("neither sensor found the line (ur screwed)")
		// FollowLineRecoverLine() // don't know what should happen here
		return
	} else if leftFoundPosition == 0 || rightFoundPosition == 0 {
		logger.Print("one sensor found the line")
		targetPosition = max(leftFoundPosition, rightFoundPosition)
	} else {
		logger.Print("both sensors found the line")
		targetPosition = (leftFoundPosition + rightFoundPosition) / 2
	}

	logger.Print("driving back to line")
	bot.DriveMotorLeft.Run(FOLLOW_LINE_RECOVER_REVERSE_SPEED)
	bot.DriveMotorRight.Run(FOLLOW_LINE_RECOVER_REVERSE_SPEED)
	for bot.DriveMotorLeft.Position() < targetPosition { bot.CycleDelay() }

	logger.Print("back to line following")
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}