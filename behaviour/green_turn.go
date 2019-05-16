package behaviour

import (
	"github.com/liam-b/robocup-2019/bot"
	"github.com/liam-b/robocup-2019/helper"
	"github.com/liam-b/robocup-2019/state_machine"
)

var greenTurn = Behaviour{
	Setup: func() {
		state_machine.Add(state_machine.State{
			Name: "green_turn.verify",
			Transitions: []string{"green_turn.left", "green_turn.right"},
			Update: func() {
				if (helper.LeftColor() == helper.COLOR_GREEN) {
					state_machine.Transition("green_turn.left")
				}

				if (helper.RightColor() == helper.COLOR_GREEN) {
					state_machine.Transition("green_turn.right")
				}
			},
		})

		state_machine.Add(state_machine.State{
			Name: "green_turn.left",
			Transitions: []string{"green_turn.cooldown"},
			Update: func() {
				bot.DriveMotorLeft.Run(-100)
				bot.DriveMotorRight.Run(100)
			},
		})

		state_machine.Add(state_machine.State{
			Name: "green_turn.right",
			Transitions: []string{"green_turn.cooldown"},
			Update: func() {
				bot.DriveMotorLeft.Run(100)
				bot.DriveMotorRight.Run(-100)
			},
		})
	},
}
