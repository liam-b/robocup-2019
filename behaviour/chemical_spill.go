package behaviour

import (
	"github.com/liam-b/robocup-2019/bot"
	"github.com/liam-b/robocup-2019/helper"
	"github.com/liam-b/robocup-2019/logger"
	"github.com/liam-b/robocup-2019/state_machine"
)

const CHEMICAL_SPILL_VERIFY_SPEED = 50
const CHEMICAL_SPILL_VERIFY_OVERSHOOT_POSITION = 100
const CHEMICAL_SPILL_VERIFY_SILVER_COLOR = 35

const CHEMICAL_SPILL_ENTER_SPEED = 200
const CHEMICAL_SPILL_ENTER_POSITION = 600

const CHEMICAL_SPILL_SEARCH_SPEED = 70
const CHEMICAL_SPILL_SEARCH_DISTANCE_THRESHOLD = 4500
var chemicalSpillSearchFoundCount = 0
var CHEMICAL_SPILL_SEARCH_FOUND_LIMIT = bot.Time(300)

const CHEMICAL_SPILL_CAPTURE_SPEED = 200
const CHEMICAL_SPILL_APPROACH_DISTANCE_THRESHOLD = 1000
var chemicalSpillApproachFoundCount = 0
var CHEMICAL_SPILL_APPROACH_FOUND_LIMIT = bot.Time(100)

const CHEMICAL_SPILL_LIFT_SPEED = 100
const CHEMICAL_SPILL_LIFT_POSITION = 400

var chemicalSpillDropCount = 0
var CHEMICAL_SPILL_DROP_OPEN_LIMIT = bot.Time(100)
var CHEMICAL_SPILL_RETREAT_LIMIT = bot.Time(300)

var chemicalSpill = Behaviour{ 
	Setup: func() {
		state_machine.Add(state_machine.State{
			Name: "chemical_spill.overshoot",
			Enter: func() {
				bot.DriveMotorLeft.RunToRelativePositionAndHold(CHEMICAL_SPILL_VERIFY_OVERSHOOT_POSITION, CHEMICAL_SPILL_VERIFY_SPEED)
				bot.DriveMotorRight.RunToRelativePositionAndHold(CHEMICAL_SPILL_VERIFY_OVERSHOOT_POSITION, CHEMICAL_SPILL_VERIFY_SPEED)
			},
			Update: func() {
				logger.Debug(bot.DriveMotorLeft.State(), bot.DriveMotorLeft.State())
				if helper.IsDriveHolding() {
					state_machine.Transition("chemical_spill.align")
				}
			},
		})

		state_machine.Add(state_machine.State{
			Name: "chemical_spill.align",
			Enter: func() {
				bot.DriveMotorLeft.Run(-CHEMICAL_SPILL_VERIFY_SPEED)
				bot.DriveMotorRight.Run(-CHEMICAL_SPILL_VERIFY_SPEED)
			},
			Update: func() {
				if bot.ColorSensorLeft.Intensity() > CHEMICAL_SPILL_VERIFY_SILVER_COLOR {
					bot.DriveMotorLeft.Hold()
				}

				if bot.ColorSensorRight.Intensity() > CHEMICAL_SPILL_VERIFY_SILVER_COLOR {
					bot.DriveMotorRight.Hold()
				}

				if helper.IsDriveHolding() {
					state_machine.Transition("chemical_spill.enter")
				}
			},
		})

		state_machine.Add(state_machine.State{
			Name: "chemical_spill.enter",
			Enter: func() {
				bot.DriveMotorLeft.RunToRelativePositionAndHold(CHEMICAL_SPILL_ENTER_POSITION, CHEMICAL_SPILL_ENTER_SPEED)
				bot.DriveMotorRight.RunToRelativePositionAndHold(CHEMICAL_SPILL_ENTER_POSITION, CHEMICAL_SPILL_ENTER_SPEED)
			},
			Update: func() {
				if helper.IsDriveHolding() {
					state_machine.Transition("chemical_spill.capture.search")
				}
			},
		})

		state_machine.Add(state_machine.State{
			Name: "chemical_spill.capture.search",
			Enter: func() {
				bot.DriveMotorLeft.Run(-CHEMICAL_SPILL_SEARCH_SPEED)
				bot.DriveMotorRight.Run(CHEMICAL_SPILL_SEARCH_SPEED)
				bot.GyroSensor.Reset()
				chemicalSpillSearchFoundCount = 0
			},
			Update: func() {
				logger.Debug(bot.UltrasonicSensor.Distance())
				if bot.UltrasonicSensor.Distance() <= CHEMICAL_SPILL_SEARCH_DISTANCE_THRESHOLD {
					chemicalSpillSearchFoundCount += 1
				} else {
					chemicalSpillSearchFoundCount /= 2
				}

				if chemicalSpillSearchFoundCount > CHEMICAL_SPILL_SEARCH_FOUND_LIMIT {
					state_machine.Transition("chemical_spill.capture.approach")
				}
			},
		})

		state_machine.Add(state_machine.State{
			Name: "chemical_spill.capture.approach",
			Enter: func() {
				bot.DriveMotorLeft.ResetPosition()
				bot.DriveMotorRight.ResetPosition()

				bot.DriveMotorLeft.Run(CHEMICAL_SPILL_CAPTURE_SPEED)
				bot.DriveMotorRight.Run(CHEMICAL_SPILL_CAPTURE_SPEED)
				chemicalSpillApproachFoundCount = 0
			},
			Update: func() {
				logger.Debug(bot.UltrasonicSensor.Distance())
				if bot.UltrasonicSensor.Distance() <= CHEMICAL_SPILL_APPROACH_DISTANCE_THRESHOLD {
					chemicalSpillApproachFoundCount += 1
				} else {
					chemicalSpillApproachFoundCount /= 2
				}

				if chemicalSpillApproachFoundCount > CHEMICAL_SPILL_APPROACH_FOUND_LIMIT {
					state_machine.Transition("chemical_spill.capture.grab")
				}
			},
		})

		state_machine.Add(state_machine.State{
			Name: "chemical_spill.capture.grab",
			Enter: func() {
				bot.DriveMotorLeft.Hold()
				bot.DriveMotorRight.Hold()
				helper.CloseClaw()
			},
			Update: func() {
				if helper.IsClawClosed() {
					state_machine.Transition("chemical_spill.capture.return")
				}
			},
		})

		state_machine.Add(state_machine.State{
			Name: "chemical_spill.capture.return",
			Enter: func() {
				bot.DriveMotorLeft.RunToAbsolutePositionAndHold(0, -CHEMICAL_SPILL_CAPTURE_SPEED)
				bot.DriveMotorRight.RunToAbsolutePositionAndHold(0, -CHEMICAL_SPILL_CAPTURE_SPEED)
			},
			Update: func() {
				if helper.IsDriveHolding() {
					state_machine.Transition("chemical_spill.save.rotate")
				}
			},
		})

		state_machine.Add(state_machine.State{
			Name: "chemical_spill.save.rotate",
			Enter: func() {
				bot.DriveMotorLeft.Run(CHEMICAL_SPILL_SEARCH_SPEED)
				bot.DriveMotorRight.Run(-CHEMICAL_SPILL_SEARCH_SPEED)
			},
			Update: func() {
				if bot.GyroSensor.Rotation() < 0 {
					state_machine.Transition("chemical_spill.save.lift")
				}
			},
		})

		state_machine.Add(state_machine.State{
			Name: "chemical_spill.save.lift",
			Enter: func() {
				bot.DriveMotorLeft.RunToAbsolutePositionAndHold(CHEMICAL_SPILL_LIFT_POSITION, CHEMICAL_SPILL_LIFT_SPEED)
				bot.DriveMotorRight.RunToAbsolutePositionAndHold(CHEMICAL_SPILL_LIFT_POSITION, CHEMICAL_SPILL_LIFT_SPEED)
				helper.RaiseClaw()
			},
			Update: func() {
				if helper.IsDriveHolding() {
					state_machine.Transition("chemical_spill.save.drop")
				}
			},
		})

		state_machine.Add(state_machine.State{
			Name: "chemical_spill.save.drop",
			Enter: func() {
				helper.ReleaseClaw()
				chemicalSpillDropCount = 0
			},
			Update: func() {
				chemicalSpillDropCount += 1

				if chemicalSpillDropCount > CHEMICAL_SPILL_DROP_OPEN_LIMIT {
					helper.OpenClaw()
				}

				if chemicalSpillDropCount > CHEMICAL_SPILL_RETREAT_LIMIT {
					state_machine.Transition("chemical_spill.save.retreat")
				}
			},
		})
	},
}
