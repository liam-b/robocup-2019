package behaviour

import (
	"github.com/liam-b/robocup-2019/bot"
	"github.com/liam-b/robocup-2019/helper"
	"github.com/liam-b/robocup-2019/state_machine"
)

const CHEMICAL_SPILL_VERIFY_SPEED = 50
const CHEMICAL_SPILL_VERIFY_OVERSHOOT_POSITION = 20
const CHEMICAL_SPILL_VERIFY_SILVER_INTENSITY = 35
const CHEMICAL_SPILL_VERIFY_GREEN_INTENSITY = 30
var chemicalSpillAlignAttemptTimer = 0
var CHEMICAL_SPILL_ALIGNMENT_ATTEMPT_TIME_LIMIT = bot.Time(7000)

const CHEMICAL_SPILL_ENTER_SPEED = 200
const CHEMICAL_SPILL_ENTER_POSITION = 580

const CHEMICAL_SPILL_SEARCH_SPEED = 70
const CHEMICAL_SPILL_SEARCH_DISTANCE_THRESHOLD = 6100
const CHEMICAL_SPILL_SEARCH_ENABLE_POSITION = 280
const CHEMICAL_SPILL_SEARCH_POSITION_LIMIT = 1500
var chemicalSpillSearchFoundCount = 0
var chemicalSpillSearchLostCount = 0
var CHEMICAL_SPILL_SEARCH_FOUND_LIMIT = bot.Time(150)
var chemicalSpillSearchFoundRotation = 0
var chemicalSpillSearchLostRotation = 0
var chemicalSpillSearchAlignRotation = 0
var chemicalSpillSearchPosition = 0

const CHEMICAL_SPILL_CAPTURE_SPEED = 150
const CHEMICAL_SPILL_CAPTURE_POSITION_LIMIT = 300
const CHEMICAL_SPILL_APPROACH_DISTANCE_THRESHOLD = 1700
var chemicalSpillCaptureApproachFoundCount = 0
var CHEMICAL_SPILL_CAPTURE_APPROACH_FOUND_LIMIT = bot.Time(100)

const CHEMICAL_SPILL_SAVE_APPROACH_SPEED = 150
const CHEMICAL_SPILL_SAVE_APPROACH_DISTANCE_THRESHOLD = 2000
const CHEMICAL_SPILL_SAVE_APPROACH_POSITION_LIMIT = 290
const CHEMICAL_SPILL_SAVE_APPROACH_SMALL_POSITION_LIMIT = 30
var chemicalSpillSaveApproachFoundCount = 0
var CHEMICAL_SPILL_SAVE_APPROACH_FOUND_LIMIT = bot.Time(100)

var chemicalSpillDropCount = 0
var CHEMICAL_SPILL_DROP_RELEASE_LIMIT = bot.Time(500)
var CHEMICAL_SPILL_DROP_OPEN_LIMIT = bot.Time(1000)
var CHEMICAL_SPILL_DROP_RETREAT_LIMIT = bot.Time(1000)

const CHEMICAL_SPILL_EXIT_SPIN_POSITION = 500

var chemicalSpill = Behaviour{ 
	Setup: func() {
		state_machine.Add(state_machine.State{
			Name: "chemical_spill.verify",
			Enter: func() {
				chemicalSpillAlignAttemptTimer = 0
				bot.DriveMotorLeft.RunToRelativePositionAndHold(CHEMICAL_SPILL_VERIFY_OVERSHOOT_POSITION, CHEMICAL_SPILL_VERIFY_SPEED)
				bot.DriveMotorRight.RunToRelativePositionAndHold(CHEMICAL_SPILL_VERIFY_OVERSHOOT_POSITION, CHEMICAL_SPILL_VERIFY_SPEED)
			},
			Update: func() {
				if helper.IsDriveStopped() {
					state_machine.Transition("chemical_spill.align.backward")
				}
			},
		})

		state_machine.Add(state_machine.State{
			Name: "chemical_spill.align.backward",
			Enter: func() {
				bot.DriveMotorLeft.Run(-CHEMICAL_SPILL_VERIFY_SPEED)
				bot.DriveMotorRight.Run(-CHEMICAL_SPILL_VERIFY_SPEED)
			},
			Update: func() {
				chemicalSpillAlignAttemptTimer += 1
				if bot.ColorSensorLeft.Intensity() > CHEMICAL_SPILL_VERIFY_SILVER_INTENSITY {
					bot.DriveMotorLeft.Hold()
				}

				if bot.ColorSensorRight.Intensity() > CHEMICAL_SPILL_VERIFY_SILVER_INTENSITY {
					bot.DriveMotorRight.Hold()
				}

				if helper.IsDriveStopped() {
					if chemicalSpillAlignAttemptTimer >= CHEMICAL_SPILL_ALIGNMENT_ATTEMPT_TIME_LIMIT {
						state_machine.Transition("chemical_spill.enter")
					} else {
						state_machine.Transition("chemical_spill.align.forward")
					}
				}
			},
		})

		state_machine.Add(state_machine.State{
			Name: "chemical_spill.align.forward",
			Enter: func() {
				bot.DriveMotorLeft.Run(CHEMICAL_SPILL_VERIFY_SPEED)
				bot.DriveMotorRight.Run(CHEMICAL_SPILL_VERIFY_SPEED)
			},
			Update: func() {
				chemicalSpillAlignAttemptTimer += 1
				if bot.ColorSensorLeft.Intensity() < CHEMICAL_SPILL_VERIFY_GREEN_INTENSITY {
					bot.DriveMotorLeft.Hold()
				}

				if bot.ColorSensorRight.Intensity() < CHEMICAL_SPILL_VERIFY_GREEN_INTENSITY {
					bot.DriveMotorRight.Hold()
				}

				if helper.IsDriveStopped() {
					state_machine.Transition("chemical_spill.align.backward")
				}
			},
		})

		state_machine.Add(state_machine.State{
			Name: "chemical_spill.enter",
			Enter: func() {
				bot.DriveMotorLeft.RunToRelativePositionAndHold(CHEMICAL_SPILL_ENTER_POSITION, CHEMICAL_SPILL_ENTER_SPEED)
				bot.DriveMotorRight.RunToRelativePositionAndHold(CHEMICAL_SPILL_ENTER_POSITION, CHEMICAL_SPILL_ENTER_SPEED)

				helper.CloseClaw()
			},
			Update: func() {
				if helper.IsDriveStopped() {
					state_machine.Transition("chemical_spill.capture.search")
				}
			},
		})

		state_machine.Add(state_machine.State{
			Name: "chemical_spill.capture.search",
			Enter: func() {
				helper.OpenClaw()
				bot.GyroSensor.Reset()
				bot.DriveMotorLeft.ResetPosition()
				bot.DriveMotorLeft.Run(-CHEMICAL_SPILL_SEARCH_SPEED)
				bot.DriveMotorRight.Run(CHEMICAL_SPILL_SEARCH_SPEED)
				chemicalSpillSearchFoundCount = 0
			},
			Update: func() {
				if bot.DriveMotorLeft.Position() < -CHEMICAL_SPILL_SEARCH_ENABLE_POSITION && bot.UltrasonicSensor.Distance() <= CHEMICAL_SPILL_SEARCH_DISTANCE_THRESHOLD {
					chemicalSpillSearchFoundCount += 1
				} else {
					chemicalSpillSearchFoundCount /= 2
				}

				// if bot.DriveMotorLeft.Position() < -CHEMICAL_SPILL_SEARCH_ENABLE_POSITION {
				// 	logger.Debug("hi")
				// 	bot.DriveMotorLeft.Coast()
				// 	bot.DriveMotorRight.Coast()
				// }

				if chemicalSpillSearchFoundCount > CHEMICAL_SPILL_SEARCH_FOUND_LIMIT {
					state_machine.Transition("chemical_spill.capture.overshoot")
				}
			},
			Exit: func() {
				chemicalSpillSearchFoundRotation = bot.GyroSensor.Rotation()
			},
		})

		state_machine.Add(state_machine.State{
			Name: "chemical_spill.capture.overshoot",
			Enter: func() {
				chemicalSpillSearchLostCount = 0
			},
			Update: func() {
				if bot.UltrasonicSensor.Distance() >= CHEMICAL_SPILL_SEARCH_DISTANCE_THRESHOLD {
					chemicalSpillSearchLostCount += 1
				} else {
					chemicalSpillSearchLostCount /= 2
				}

				if chemicalSpillSearchLostCount > CHEMICAL_SPILL_SEARCH_FOUND_LIMIT {
					state_machine.Transition("chemical_spill.capture.align")
				}
			},
			Exit: func() {
				chemicalSpillSearchLostRotation = bot.GyroSensor.Rotation()
			},
		})

		state_machine.Add(state_machine.State{
			Name: "chemical_spill.capture.align",
			Enter: func() {
				bot.DriveMotorLeft.Run(CHEMICAL_SPILL_SEARCH_SPEED)
				bot.DriveMotorRight.Run(-CHEMICAL_SPILL_SEARCH_SPEED)
				chemicalSpillSearchAlignRotation = (chemicalSpillSearchFoundRotation * 2 + chemicalSpillSearchLostRotation) / 3
			},
			Update: func() {
				if bot.GyroSensor.Rotation() < chemicalSpillSearchAlignRotation {
					chemicalSpillSearchPosition = bot.DriveMotorLeft.Position()
					state_machine.Transition("chemical_spill.capture.approach")
				}
			},
		})

		state_machine.Add(state_machine.State{
			Name: "chemical_spill.capture.approach",
			Enter: func() {
				bot.DriveMotorLeft.ResetPosition()
				bot.DriveMotorRight.ResetPosition()
				bot.DriveMotorLeft.RunToRelativePositionAndHold(CHEMICAL_SPILL_CAPTURE_POSITION_LIMIT, CHEMICAL_SPILL_CAPTURE_SPEED)
				bot.DriveMotorRight.RunToRelativePositionAndHold(CHEMICAL_SPILL_CAPTURE_POSITION_LIMIT, CHEMICAL_SPILL_CAPTURE_SPEED)
				chemicalSpillCaptureApproachFoundCount = 0
			},
			Update: func() {
				if bot.UltrasonicSensor.Distance() <= CHEMICAL_SPILL_APPROACH_DISTANCE_THRESHOLD {
					chemicalSpillCaptureApproachFoundCount += 1
				} else {
					chemicalSpillCaptureApproachFoundCount /= 2
				}

				if chemicalSpillCaptureApproachFoundCount > CHEMICAL_SPILL_SAVE_APPROACH_FOUND_LIMIT || helper.IsDriveStopped() {
					state_machine.Transition("chemical_spill.capture.grab")
				}
			},
		})

		state_machine.Add(state_machine.State{
			Name: "chemical_spill.capture.grab",
			Enter: func() {
				bot.DriveMotorLeft.Run(CHEMICAL_SPILL_VERIFY_SPEED)
				bot.DriveMotorRight.Run(CHEMICAL_SPILL_VERIFY_SPEED)
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
				bot.DriveMotorLeft.RunToAbsolutePositionAndHold(0, CHEMICAL_SPILL_CAPTURE_SPEED)
				bot.DriveMotorRight.RunToAbsolutePositionAndHold(0, CHEMICAL_SPILL_CAPTURE_SPEED)
			},
			Update: func() {
				if helper.IsDriveStopped() {
					state_machine.Transition("chemical_spill.save.rotate")
				}
			},
		})

		state_machine.Add(state_machine.State{
			Name: "chemical_spill.save.rotate",
			Enter: func() {
				bot.DriveMotorLeft.ResetPosition()
				bot.DriveMotorLeft.Run(CHEMICAL_SPILL_SEARCH_SPEED)
				bot.DriveMotorRight.Run(-CHEMICAL_SPILL_SEARCH_SPEED)
			},
			Update: func() {
				if bot.DriveMotorLeft.Position() >= -chemicalSpillSearchPosition {
					bot.DriveMotorLeft.Hold()
					bot.DriveMotorRight.Hold()
				}

				if helper.IsDriveStopped() {
					state_machine.Transition("chemical_spill.save.lift")
				}
			},
		})

		state_machine.Add(state_machine.State{
			Name: "chemical_spill.save.lift",
			Enter: func() {
				helper.RaiseClaw()
			},
			Update: func() {
				if helper.IsClawRaised() {
					state_machine.Transition("chemical_spill.save.approach")
				}
			},
		})

		state_machine.Add(state_machine.State{
			Name: "chemical_spill.save.approach",
			Enter: func() {
				bot.DriveMotorLeft.RunToRelativePositionAndHold(CHEMICAL_SPILL_SAVE_APPROACH_POSITION_LIMIT, CHEMICAL_SPILL_SAVE_APPROACH_SPEED)
				bot.DriveMotorRight.RunToRelativePositionAndHold(CHEMICAL_SPILL_SAVE_APPROACH_POSITION_LIMIT, CHEMICAL_SPILL_SAVE_APPROACH_SPEED)
				chemicalSpillSaveApproachFoundCount = 0
			},
			Update: func() {
				if bot.UltrasonicSensor.Distance() <= CHEMICAL_SPILL_SAVE_APPROACH_DISTANCE_THRESHOLD && chemicalSpillSaveApproachFoundCount >= 0 {
					chemicalSpillSaveApproachFoundCount += 1
				} else {
					chemicalSpillSaveApproachFoundCount /= 2
				}

				// if chemicalSpillSaveApproachFoundCount > CHEMICAL_SPILL_SAVE_APPROACH_FOUND_LIMIT {
				// 	chemicalSpillSaveApproachFoundCount = -1
				// 	bot.DriveMotorLeft.Hold()
				// 	bot.DriveMotorRight.Hold()
				// 	// bot.DriveMotorLeft.RunToRelativePositionAndHold(CHEMICAL_SPILL_SAVE_APPROACH_SMALL_POSITION_LIMIT, CHEMICAL_SPILL_SAVE_APPROACH_SPEED)
				// 	// bot.DriveMotorRight.RunToRelativePositionAndHold(CHEMICAL_SPILL_SAVE_APPROACH_SMALL_POSITION_LIMIT, CHEMICAL_SPILL_SAVE_APPROACH_SPEED)
				// }

				if helper.IsDriveStopped() {
					state_machine.Transition("chemical_spill.save.drop")
				}
			},
		})

		state_machine.Add(state_machine.State{
			Name: "chemical_spill.save.drop",
			Enter: func() {
				chemicalSpillDropCount = 0
			},
			Update: func() {
				chemicalSpillDropCount += 1

				if chemicalSpillDropCount > CHEMICAL_SPILL_DROP_RELEASE_LIMIT {
					helper.ReleaseClaw()
				}

				if chemicalSpillDropCount > CHEMICAL_SPILL_DROP_RELEASE_LIMIT + CHEMICAL_SPILL_DROP_OPEN_LIMIT {
					helper.OpenClaw()
				}

				if chemicalSpillDropCount > CHEMICAL_SPILL_DROP_RELEASE_LIMIT + CHEMICAL_SPILL_DROP_OPEN_LIMIT + CHEMICAL_SPILL_DROP_RETREAT_LIMIT {
					state_machine.Transition("chemical_spill.exit")
				}
			},
		})

		state_machine.Add(state_machine.State{
			Name: "chemical_spill.exit",
			Enter: func() {
				bot.DriveMotorLeft.Run(-CHEMICAL_SPILL_ENTER_SPEED)
				bot.DriveMotorRight.Run(-CHEMICAL_SPILL_ENTER_SPEED)
				helper.LowerClaw()
			},
			Update: func() {
				if bot.ColorSensorLeft.Intensity() > CHEMICAL_SPILL_VERIFY_SILVER_INTENSITY {
					bot.DriveMotorLeft.Hold()
				}

				if bot.ColorSensorRight.Intensity() > CHEMICAL_SPILL_VERIFY_SILVER_INTENSITY {
					bot.DriveMotorRight.Hold()
				}

				if helper.IsDriveStopped() {
					state_machine.Transition("chemical_spill.exit.spin")
				}
			},
		})

		state_machine.Add(state_machine.State{
			Name: "chemical_spill.exit.spin",
			Enter: func() {
				bot.DriveMotorLeft.RunToRelativePositionAndHold(-CHEMICAL_SPILL_EXIT_SPIN_POSITION, CHEMICAL_SPILL_ENTER_SPEED)
				bot.DriveMotorRight.RunToRelativePositionAndHold(CHEMICAL_SPILL_EXIT_SPIN_POSITION, CHEMICAL_SPILL_ENTER_SPEED)
			},
			Update: func() {
				if helper.IsDriveStopped() {
					state_machine.Transition("follow_line.follow")
				}
			},
		})

		// state_machine.Add(state_machine.State{
		// 	Name: "chemical_spill.exit.backup",
		// 	Enter: func() {
		// 		bot.DriveMotorLeft.Run(-CHEMICAL_SPILL_ENTER_SPEED)
		// 		bot.DriveMotorRight.Run(-CHEMICAL_SPILL_ENTER_SPEED)
		// 	},
		// 	Update: func() {
		// 		if bot.ColorSensorLeft.Intensity() > CHEMICAL_SPILL_VERIFY_SILVER_INTENSITY {
		// 			bot.DriveMotorLeft.Hold()
		// 		}

		// 		if bot.ColorSensorRight.Intensity() > CHEMICAL_SPILL_VERIFY_SILVER_INTENSITY {
		// 			bot.DriveMotorRight.Hold()
		// 		}

		// 		if helper.IsDriveStopped() {
		// 			state_machine.Transition("chemical_spill.exit.spin")
		// 		}
		// 	},
		// })
	},
}
