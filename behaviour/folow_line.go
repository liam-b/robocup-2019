package behaviour

import (
	"github.com/liam-b/robocup-2019/bot"
	"github.com/liam-b/robocup-2019/helper"
	"github.com/liam-b/robocup-2019/state_machine"
)

const FOLLOW_LINE_GREEN_VERIFY_MIDDLE_INTENSITY = 0.2
const FOLLOW_LINE_LOST_RECAPTURE_SPEED = 200

var followLineLostCount = 0
var FOLLOW_LINE_LOST_LIMIT = bot.Time(1000)

var followLineFoundGreenCount = 0
var FOLLOW_LINE_FOUND_GREEN_THRESHOLD = bot.Time(100)

var followLine = Behaviour{ 
	Setup: func() {
		state_machine.Add(state_machine.State{
			Name: "follow_line.follow",
			Enter: func() {
				followLineFoundGreenCount = 0
				followLineLostCount = 0
			},
			Update: func() {
				left, right := helper.PID()
				bot.DriveMotorLeft.Run(left)
				bot.DriveMotorRight.Run(right)

				// if (helper.LeftColor() == helper.COLOR_GREEN || helper.RightColor() == helper.COLOR_GREEN) && helper.MiddleError() < FOLLOW_LINE_GREEN_VERIFY_MIDDLE_INTENSITY {
				// 	state_machine.Transition("green_turn.verify")
				// }

				if helper.LeftColor() == helper.COLOR_GREEN || helper.RightColor() == helper.COLOR_GREEN {
					followLineFoundGreenCount += 1
				}

				if followLineFoundGreenCount > FOLLOW_LINE_FOUND_GREEN_THRESHOLD {
					state_machine.Transition("green_turn.verify")
				}

				if helper.LeftColor() == helper.COLOR_WHITE && helper.RightColor() == helper.COLOR_WHITE && helper.MiddleColor() == helper.COLOR_WHITE {
					followLineLostCount += 1
				} else {
					followLineLostCount /= 2
				}

				if followLineLostCount > FOLLOW_LINE_LOST_LIMIT {
					state_machine.Transition("follow_line.recapture")
				}
			},
		})

		state_machine.Add(state_machine.State{
			Name: "follow_line.recapture",
			Update: func() {
				bot.DriveMotorLeft.Run(-FOLLOW_LINE_LOST_RECAPTURE_SPEED)
				bot.DriveMotorRight.Run(-FOLLOW_LINE_LOST_RECAPTURE_SPEED)

				if helper.MiddleColor() == helper.COLOR_BLACK {
					state_machine.Transition("follow_line.follow")
				}
			},
		})
	},
}
