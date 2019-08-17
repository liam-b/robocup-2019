package behaviour

// import (
// 	"github.com/liam-b/robocup-2019/bot"
// 	"github.com/liam-b/robocup-2019/helper"
// 	"github.com/liam-b/robocup-2019/state_machine"
// )

// const GREEN_TURN_JUNCTION_INTENSNITY = 0.3
// const GREEN_TURN_JUNCTION_SPEED = 200
// const GREEN_TURN_INNER_SPEED = -150
// const GREEN_TURN_OUTER_SPEED = 100
// const GREEN_TURN_EXIT_INTENSNITY = 0.6

// var greenTurnEndWait = 0
// var GREEN_TURN_END_WAIT_LIMIT = bot.Time(400)
// const GREEN_TURN_MIDDLE_JUNCTION_INTENSNITY = 0.3

// var greenTurnEndCooldown = 0
// var GREEN_TURN_END_COOLDOWN_LIMIT = bot.Time(1000)

// var greenTurnFoundChemicalSpillCount = 0
// var GREEN_TURN_FOUND_CHEMICAL_SPILL_THRESHOLD = bot.Time(50)

// var greenTurn = Behaviour{
// 	Setup: func() {
// 		state_machine.Add(state_machine.State{
// 			Name: "green_turn.verify",
// 			Enter: func() {
// 				greenTurnFoundChemicalSpillCount = 0
// 			},
// 			Update: func() {
// 				if helper.LeftColor() == helper.COLOR_GREEN {
// 					state_machine.Transition("green_turn.left")
// 				}

// 				if helper.RightColor() == helper.COLOR_GREEN {
// 					state_machine.Transition("green_turn.right")
// 				}

// 				if helper.LeftColor() == helper.COLOR_GREEN && helper.RightColor() == helper.COLOR_GREEN {
// 					greenTurnFoundChemicalSpillCount += 1
// 				} else {
// 					greenTurnFoundChemicalSpillCount /= 2
// 				}
// 			},
// 		})

// 		state_machine.Add(state_machine.State{
// 			Name: "green_turn.left",
// 			Update: func() {
// 				bot.DriveMotorLeft.Run(GREEN_TURN_JUNCTION_SPEED)
// 				bot.DriveMotorRight.Run(GREEN_TURN_JUNCTION_SPEED)

// 				if helper.LeftError() < GREEN_TURN_JUNCTION_INTENSNITY {
// 					state_machine.Transition("green_turn.left_turn")
// 				}

// 				if helper.LeftColor() == helper.COLOR_GREEN && helper.RightColor() == helper.COLOR_GREEN {
// 					greenTurnFoundChemicalSpillCount += 1
// 				} else {
// 					greenTurnFoundChemicalSpillCount /= 2
// 				}
// 				if greenTurnFoundChemicalSpillCount > GREEN_TURN_FOUND_CHEMICAL_SPILL_THRESHOLD {
// 					state_machine.Transition("chemical_spill.verify")
// 				}
// 			},
// 		})

// 		state_machine.Add(state_machine.State{
// 			Name: "green_turn.left_turn",
// 			Enter: func() {
// 				greenTurnEndWait = 0
// 			},
// 			Update: func() {
// 				bot.DriveMotorLeft.Run(GREEN_TURN_INNER_SPEED)
// 				bot.DriveMotorRight.Run(GREEN_TURN_OUTER_SPEED)

// 				greenTurnEndWait += 1
// 				if greenTurnEndWait > GREEN_TURN_END_WAIT_LIMIT && helper.MiddleError() < GREEN_TURN_MIDDLE_JUNCTION_INTENSNITY {
// 					state_machine.Transition("green_turn.cooldown")
// 				}
// 			},
// 		})

// 		state_machine.Add(state_machine.State{
// 			Name: "green_turn.right",
// 			Update: func() {
// 				bot.DriveMotorLeft.Run(GREEN_TURN_JUNCTION_SPEED)
// 				bot.DriveMotorRight.Run(GREEN_TURN_JUNCTION_SPEED)

// 				if helper.RightError() < GREEN_TURN_JUNCTION_INTENSNITY {
// 					state_machine.Transition("green_turn.right_turn")
// 				}

// 				if helper.LeftColor() == helper.COLOR_GREEN && helper.RightColor() == helper.COLOR_GREEN {
// 					greenTurnFoundChemicalSpillCount += 1
// 				} else {
// 					greenTurnFoundChemicalSpillCount /= 2
// 				}
// 				if greenTurnFoundChemicalSpillCount > GREEN_TURN_FOUND_CHEMICAL_SPILL_THRESHOLD {
// 					state_machine.Transition("chemical_spill.verify")
// 				}
// 			},
// 		})
		
// 		state_machine.Add(state_machine.State{
// 			Name: "green_turn.right_turn",
// 			Enter: func() {
// 				greenTurnEndWait = 0
// 			},
// 			Update: func() {
// 				bot.DriveMotorLeft.Run(GREEN_TURN_OUTER_SPEED)
// 				bot.DriveMotorRight.Run(GREEN_TURN_INNER_SPEED)

// 				greenTurnEndWait += 1
// 				if greenTurnEndWait > GREEN_TURN_END_WAIT_LIMIT && helper.MiddleError() < GREEN_TURN_MIDDLE_JUNCTION_INTENSNITY {
// 					state_machine.Transition("green_turn.cooldown")
// 				}
// 			},
// 		})

// 		state_machine.Add(state_machine.State{
// 			Name: "green_turn.cooldown",
// 			Enter: func() {
// 				helper.ResetPID()
// 				greenTurnEndCooldown = 0
// 			},
// 			Update: func() {
// 				left, right := helper.PID()
// 				bot.DriveMotorLeft.Run(left)
// 				bot.DriveMotorRight.Run(right)

// 				greenTurnEndCooldown += 1
// 				if greenTurnEndCooldown > GREEN_TURN_END_COOLDOWN_LIMIT {
// 					state_machine.Transition("follow_line.follow")
// 				}
// 			},
// 		})
// 	},
// }
