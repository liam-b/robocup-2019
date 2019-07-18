package behaviour

import (
	"github.com/liam-b/robocup-2019/bot"
	"github.com/liam-b/robocup-2019/helper"
	"github.com/liam-b/robocup-2019/state_machine"
	"github.com/liam-b/robocup-2019/logger"
)

const FOLLOW_LINE_GREEN_VERIFY_MIDDLE_INTENSITY = 0.4
const FOLLOW_LINE_LOST_RECAPTURE_SPEED = 200

var followLineFoundGreenCount = 0
var FOLLOW_LINE_FOUND_GREEN_THRESHOLD = bot.Time(100)

var followLineFoundChemicalSpillCount = 0
var FOLLOW_LINE_FOUND_CHEMICAL_SPILL_THRESHOLD = bot.Time(50)

var followLineLostCount = 0
var FOLLOW_LINE_LOST_LIMIT = bot.Time(1000)

var waterTowerCounter = 0
var WATER_TOWER_WAIT_LIMIT = bot.Time(200)
var WATER_TOWER_COUNT_LIMIT = 1
var waterTowerCheekyCount = 0

var followLine = Behaviour{ 
	Setup: func() {
		state_machine.Add(state_machine.State{
			Name: "follow_line.follow",
			Enter: func() {
				helper.ResetPID()
				followLineFoundGreenCount = 0
				followLineFoundChemicalSpillCount = 0
				followLineLostCount = 0
				waterTowerCounter = 0
				waterTowerCheekyCount = 0
			},
			Update: func() {
				left, right := helper.PID()
				bot.DriveMotorLeft.Run(left)
				bot.DriveMotorRight.Run(right)

				// if helper.MiddleError() < PAUSE_INTENSITY_THRESHOLD {
				// 	state_machine.Transition("pause.wait")
				// }
				
				if (bot.UltrasonicSensor.Distance() <= 2900) {
					waterTowerCounter += 1
				} else {
					waterTowerCounter /= 2
				}

				if (waterTowerCounter > WATER_TOWER_WAIT_LIMIT && waterTowerCheekyCount < WATER_TOWER_COUNT_LIMIT) {
					logger.Trace("move into water tower state")
					waterTowerCheekyCount += 1
					state_machine.Transition("water_tower.verify")
				}

				if (helper.LeftColor() == helper.COLOR_GREEN || helper.RightColor() == helper.COLOR_GREEN) && helper.MiddleError() < FOLLOW_LINE_GREEN_VERIFY_MIDDLE_INTENSITY && !(helper.LeftColor() == helper.COLOR_GREEN && helper.RightColor() == helper.COLOR_GREEN) {
					followLineFoundGreenCount += 1
				} else {
					followLineFoundGreenCount /= 2
				} 

				if followLineFoundGreenCount > FOLLOW_LINE_FOUND_GREEN_THRESHOLD {
					state_machine.Transition("green_turn.verify")
				}

				if helper.LeftColor() == helper.COLOR_GREEN && helper.RightColor() == helper.COLOR_GREEN {
					followLineFoundChemicalSpillCount += 1
				} else {
					followLineFoundChemicalSpillCount /= 2
				} 

				if followLineFoundChemicalSpillCount > FOLLOW_LINE_FOUND_CHEMICAL_SPILL_THRESHOLD {
					state_machine.Transition("chemical_spill.verify")
				}

				if helper.LeftColor() == helper.COLOR_WHITE && helper.RightColor() == helper.COLOR_WHITE && helper.MiddleColor() == helper.COLOR_WHITE {
					followLineLostCount += 1
				} else {
					followLineLostCount /= 2
				}

				// if followLineLostCount > FOLLOW_LINE_LOST_LIMIT {
				// 	state_machine.Transition("follow_line.recapture")
				// }
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
