package behaviour

import (
	"github.com/liam-b/robocup-2019/bot"
	"github.com/liam-b/robocup-2019/helper"
	"github.com/liam-b/robocup-2019/state_machine"
	// "github.com/liam-b/robocup-2019/logger"
)

var followLine = Behaviour{
	Setup: func() {
		state_machine.Add(state_machine.State{
			Name: "follow_line",
			Transitions: []string{"pause", "green_turn.verify", "water_tower.verify"},
			Update: func() {
				left, right := helper.PID()
				bot.DriveMotorLeft.Run(left)
				bot.DriveMotorRight.Run(right)

				// logger.Debug(helper.LeftColor(), helper.RightColor(), helper.MiddleColor())

				if (helper.LeftColor() == helper.COLOR_GREEN || helper.RightColor() == helper.COLOR_GREEN) && helper.MiddleColor() == helper.COLOR_WHITE {
					state_machine.Transition("green_turn.verify")
				}
			},
		})
	},
}
