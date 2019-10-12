package behaviour

import (
	"github.com/liam-b/robocup-2019/bot"
	"github.com/liam-b/robocup-2019/helper"
	"github.com/liam-b/robocup-2019/logger"
	"time"
)

var (
	FOLLOW_LINE_GREEN_TURN_THRESHOLD = 10
	FOLLOW_LINE_GREEN_TURN_TIME_LIMIT = bot.Time(50)

	FOLLOW_LINE_RECOVER_LOST_THRESHOLD = 39
	FOLLOW_LINE_RECOVER_LOST_MIDDLE_THRESHOLD = 24
	FOLLOW_LINE_RECOVER_FOUND_THRESHOLD = 15
	FOLLOW_LINE_RECOVER_LOST_TIME_LIMIT = bot.Time(1000) // 700
	FOLLOW_LINE_RECOVER_REVERSE_SPEED = 150
	FOLLOW_LINE_RECOVER_REVERSE_POSITION_LIMIT = 700

	FOLLOW_LINE_GREEN_COLOUR_TIME_LIMIT = bot.Time(50)
	FOLLOW_LINE_GREEN_COLOUR_THRESHOLD = 55
)

func FollowLine() {
	logger.Print("started following line")
	helper.ResetPID()

	leftGreenCount := 0
	rightGreenCount := 0
	lineLostCount := 0
	waterTowerCount := 0
	// greenColourNotOverflowCount := 0
	for {
		logger.Print(bot.ColorSensorLeft.Intensity())
		left, right := helper.PID()
		bot.DriveMotorLeft.Run(left)
		bot.DriveMotorRight.Run(right)

		// write yellow check here if double is took an L then remove the furthermost wrapper
		// if helper.RightColor() != 4 && helper.LeftColor() != 4 {
		// 	if bot.ColorSensorLeft.Intensity() <= FOLLOW_LINE_GREEN_COLOUR_THRESHOLD && bot.ColorSensorRight.Intensity() <= FOLLOW_LINE_GREEN_COLOUR_THRESHOLD {
		// 		greenColourNotOverflowCount ++
		// 		if helper.LeftGreen() > FOLLOW_LINE_GREEN_TURN_THRESHOLD {
		// 			leftGreenCount++
		// 			if leftGreenCount > FOLLOW_LINE_GREEN_TURN_TIME_LIMIT {
		// 				logger.Print("now we turn left")
		// 				logger.Print(bot.ColorSensorLeft.Intensity())
		// 				GreenTurnLeft()
		// 				leftGreenCount = 0
		// 			}
		// 		} else {
		// 			leftGreenCount /= 2
		// 		}
		
		// 		if helper.RightGreen() > FOLLOW_LINE_GREEN_TURN_THRESHOLD {
		// 			rightGreenCount++
		// 			if rightGreenCount > FOLLOW_LINE_GREEN_TURN_TIME_LIMIT {
		// 				logger.Print("now we turn right")
		// 				logger.Print(bot.ColorSensorRight.Intensity())
		// 				GreenTurnRight()
		// 				rightGreenCount = 0
		// 			}
		// 		} else {
		// 			rightGreenCount /= 2
		// 		}
		// 	} else {
		// 		greenColourNotOverflowCount /= 2
		// 	}
		// }
		

		if helper.LeftGreen() > FOLLOW_LINE_GREEN_TURN_THRESHOLD {
			leftGreenCount++
			if leftGreenCount > FOLLOW_LINE_GREEN_TURN_TIME_LIMIT {
				logger.Print("now we turn left")
				logger.Print(bot.ColorSensorLeft.Intensity())
				GreenTurnLeft()
				leftGreenCount = 0
			}
		} else {
			leftGreenCount /= 2
		}

		if helper.RightGreen() > FOLLOW_LINE_GREEN_TURN_THRESHOLD {
			rightGreenCount++
			if rightGreenCount > FOLLOW_LINE_GREEN_TURN_TIME_LIMIT {
				logger.Print("now we turn right")
				logger.Print(bot.ColorSensorRight.Intensity())
				GreenTurnRight()
				rightGreenCount = 0
			}
		} else {
			rightGreenCount /= 2
		}

		if bot.ColorSensorLeft.Intensity() > FOLLOW_LINE_RECOVER_LOST_THRESHOLD && bot.ColorSensorRight.Intensity() > FOLLOW_LINE_RECOVER_LOST_THRESHOLD && bot.ColorSensorMiddle.Intensity() > FOLLOW_LINE_RECOVER_LOST_MIDDLE_THRESHOLD {
			lineLostCount++
			if lineLostCount > FOLLOW_LINE_RECOVER_LOST_TIME_LIMIT {
				FollowLineRecoverLine()
				lineLostCount = 0
			}
		} else {
			lineLostCount /= 2
		}

		if bot.UltrasonicSensor.Distance() < WATER_TOWER_VERIFY_DISTANCE {
			waterTowerCount++
			if waterTowerCount > WATER_TOWER_VERIFY_COUNT {
				WaterTowerVerify()
				waterTowerCount = 0
			}
		} else {
			waterTowerCount /= 2
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
	for bot.DriveMotorLeft.Position() > -FOLLOW_LINE_RECOVER_REVERSE_POSITION_LIMIT && (leftFoundPosition == 0 || rightFoundPosition == 0) {
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

	bot.DriveMotorLeft.Hold()
	bot.DriveMotorRight.Hold()

	time.Sleep(time.Millisecond * 400)

	logger.Print("back to line following")
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}