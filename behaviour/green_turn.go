package behaviour

import (
	"github.com/liam-b/robocup-2019/bot"
	"github.com/liam-b/robocup-2019/helper"
	"github.com/liam-b/robocup-2019/state_machine"
)

const GREEN_TURN_EXIT_INTENSNITY = 0.2
const GREEN_TURN_INNER_SPEED = -50
const GREEN_TURN_OUTER_SPEED = 200

var greenTurnEndCooldown = 0
var GREEN_TURN_END_COOLDOWN_LIMIT = bot.Time(500)

var greenTurn = Behaviour{
	Setup: func() {
		state_machine.Add(state_machine.State{
			Name: "green_turn.verify",
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
			Enter: func() {
				greenTurnEndCooldown = 0
				bot.DriveMotorLeft.Run(GREEN_TURN_INNER_SPEED)
				bot.DriveMotorRight.Run(GREEN_TURN_OUTER_SPEED)
			},
			Update: func() {
				greenTurnEndCooldown += 1
				if greenTurnEndCooldown > GREEN_TURN_END_COOLDOWN_LIMIT {
					if (helper.RightColor() == helper.COLOR_BLACK) {
						state_machine.Transition("follow_line.follow")
					}
				}
			},
			Exit: func() {
				bot.DriveMotorLeft.Brake()
				bot.DriveMotorRight.Brake()
			},
		})

		state_machine.Add(state_machine.State{
			Name: "green_turn.right",
			Enter: func() {
				greenTurnEndCooldown = 0
				bot.DriveMotorLeft.Run(GREEN_TURN_OUTER_SPEED)
				bot.DriveMotorRight.Run(GREEN_TURN_INNER_SPEED)
			},
			Update: func() {
				greenTurnEndCooldown += 1
				if greenTurnEndCooldown > GREEN_TURN_END_COOLDOWN_LIMIT {
					if (helper.LeftColor() == helper.COLOR_BLACK) {
						state_machine.Transition("follow_line.follow")
					}
				}
			},
			Exit: func() {
				bot.DriveMotorLeft.Brake()
				bot.DriveMotorRight.Brake()
			},
		})
	},
}
