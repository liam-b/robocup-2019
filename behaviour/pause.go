package behaviour

import (
	"github.com/liam-b/robocup-2019/bot"
	"github.com/liam-b/robocup-2019/helper"
	"github.com/liam-b/robocup-2019/state_machine"
)

const PAUSE_INTENSITY_THRESHOLD = 0.3
var PAUSE_WAIT_LIMIT = bot.Time(100)
var pauseWaitValue = 0

var pause = Behaviour{ 
	Setup: func() {
		state_machine.Add(state_machine.State{
			Name: "pause.wait",
			Enter: func() {
				bot.DriveMotorLeft.Coast()
				bot.DriveMotorRight.Coast()
			},
			Update: func() {
				if helper.MiddleError() > PAUSE_INTENSITY_THRESHOLD {
					state_machine.Transition("pause.delay")
				}
			},
		})

		state_machine.Add(state_machine.State{
			Name: "pause.delay",
			Enter: func() {
				pauseWaitValue = 0
			},
			Update: func() {
				pauseWaitValue += 1
				if pauseWaitValue > PAUSE_WAIT_LIMIT {
					state_machine.Transition("follow_line.follow")
				}
			},
		})
	},
}
