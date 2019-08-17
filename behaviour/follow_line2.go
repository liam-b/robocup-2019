package behaviour

import (
	"github.com/liam-b/robocup-2019/bot"
	"github.com/liam-b/robocup-2019/helper"
	// "github.com/liam-b/robocup-2019/logger"
)

var (
	FOLLOW_LINE_GREEN_TURN_THRESHOLD = 12
	FOLLOW_LINE_GREEN_TURN_TIME_LIMIT = bot.Time(50)
)

func FollowLine() {
	helper.ResetPID()

	leftGreenCount := 0
	rightGreenCount := 0
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

		bot.CycleDelay()
		// time.Sleep(time.Millisecond * 20)
	}
}
