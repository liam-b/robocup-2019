package behaviour

import (
	"github.com/liam-b/robocup-2019/bot"
	"github.com/liam-b/robocup-2019/helper"
)

var (
	CHEMICAL_SPILL_VERIFY_SPEED = 50
	CHEMICAL_SPILL_VERIFY_OVERSHOOT_POSITION = 60
	CHEMICAL_SPILL_VERIFY_SILVER_INTENSITY = 35
	CHEMICAL_SPILL_VERIFY_GREEN_INTENSITY = 30

	chemicalSpillVerifyAlignAttemptCount = 0
	CHEMICAL_SPILL_VERIFY_ALIGN_ATTEMPTS = 4

	CHEMICAL_SPILL_ENTER_SPEED = 250
	CHEMICAL_SPILL_ENTER_POSITION = 580

	CHEMICAL_SPILL_SEARCH_SPEED = 70
	CHEMICAL_SPILL_SEARCH_ENABLE_POSITION = 450
	CHEMICAL_SPILL_SEARCH_DISTANCE_THRESHOLD = 3400
	CHEMICAL_SPILL_SEARCH_FIRST_DIAGONAL_POSITION = 300
	CHEMICAL_SPILL_SEARCH_LAST_DIAGONAL_POSITION = 4000

	CHEMICAL_SPILL_SEARCH_CHECK_POSITION = 200
	CHEMICAL_SPILL_SEARCH_CHECK_DISTANCE_THRESHOLD = 2800

	CHEMICAL_SPILL_RESCUE_SPEED = 100
	CHEMICAL_SPILL_RESCUE_BLOCK_POSITION = 200
	CHEMICAL_SPILL_RESCUE_EXIT_OFFSET = 80
)

func ChemicalSpillVerify() {
	bot.DriveMotorLeft.RunToRelativePositionAndHold(CHEMICAL_SPILL_VERIFY_OVERSHOOT_POSITION, CHEMICAL_SPILL_VERIFY_SPEED)
	bot.DriveMotorRight.RunToRelativePositionAndHold(CHEMICAL_SPILL_VERIFY_OVERSHOOT_POSITION, CHEMICAL_SPILL_VERIFY_SPEED)
	for !helper.IsDriveStopped() { bot.CycleDelay() }

	chemicalSpillVerifyAlignAttemptCount = 0
	ChemicalSpillBackwardAlign()
}

func ChemicalSpillBackwardAlign() {
	bot.DriveMotorLeft.Run(-CHEMICAL_SPILL_VERIFY_SPEED)
	bot.DriveMotorRight.Run(-CHEMICAL_SPILL_VERIFY_SPEED)
	for !helper.IsDriveStopped() {
		if bot.ColorSensorLeft.Intensity() < CHEMICAL_SPILL_VERIFY_SILVER_INTENSITY {
			bot.DriveMotorLeft.Hold()
		}

		if bot.ColorSensorRight.Intensity() < CHEMICAL_SPILL_VERIFY_SILVER_INTENSITY {
			bot.DriveMotorRight.Hold()
		}

		bot.CycleDelay()
	}

	ChemicalSpillForwardAlign()
}

func ChemicalSpillForwardAlign() {
	bot.DriveMotorLeft.Run(CHEMICAL_SPILL_VERIFY_SPEED)
	bot.DriveMotorRight.Run(CHEMICAL_SPILL_VERIFY_SPEED)
	for !helper.IsDriveStopped() {
		if bot.ColorSensorLeft.Intensity() > CHEMICAL_SPILL_VERIFY_GREEN_INTENSITY {
			bot.DriveMotorLeft.Hold()
		}

		if bot.ColorSensorRight.Intensity() > CHEMICAL_SPILL_VERIFY_GREEN_INTENSITY {
			bot.DriveMotorRight.Hold()
		}

		bot.CycleDelay()
	}

	chemicalSpillVerifyAlignAttemptCount++
	if chemicalSpillVerifyAlignAttemptCount >= CHEMICAL_SPILL_VERIFY_ALIGN_ATTEMPTS {
		ChemicalSpillSearch()
	} else {
		ChemicalSpillBackwardAlign()
	}
}

func ChemicalSpillSearch() {
	bot.DriveMotorLeft.RunToRelativePositionAndHold(CHEMICAL_SPILL_ENTER_POSITION, CHEMICAL_SPILL_ENTER_SPEED)
	bot.DriveMotorRight.RunToRelativePositionAndHold(CHEMICAL_SPILL_ENTER_POSITION, CHEMICAL_SPILL_ENTER_SPEED)
	for !helper.IsDriveStopped() { bot.CycleDelay() }

	bot.DriveMotorLeft.ResetPosition()
	bot.DriveMotorLeft.RunToRelativePositionAndHold(CHEMICAL_SPILL_SEARCH_FIRST_DIAGONAL_POSITION, CHEMICAL_SPILL_SEARCH_SPEED)
	bot.DriveMotorRight.RunToRelativePositionAndHold(-CHEMICAL_SPILL_SEARCH_FIRST_DIAGONAL_POSITION, CHEMICAL_SPILL_SEARCH_SPEED)
	for !helper.IsDriveStopped() { bot.CycleDelay() }

	ChemicalSpillCheckCurrentPosition()

	bot.DriveMotorLeft.RunToRelativePositionAndHold(CHEMICAL_SPILL_SEARCH_LAST_DIAGONAL_POSITION, CHEMICAL_SPILL_SEARCH_SPEED)
	bot.DriveMotorRight.RunToRelativePositionAndHold(-CHEMICAL_SPILL_SEARCH_LAST_DIAGONAL_POSITION, CHEMICAL_SPILL_SEARCH_SPEED)
	for !helper.IsDriveStopped() {
		if bot.DriveMotorLeft.Position() >= CHEMICAL_SPILL_SEARCH_ENABLE_POSITION {
			if bot.UltrasonicSensor.Distance() <= CHEMICAL_SPILL_SEARCH_DISTANCE_THRESHOLD {
				foundCan := ChemicalSpillCheckCurrentPosition()
				if foundCan { return }
			}
		}

		bot.CycleDelay()
	}

	ChemicalSpillCheckCurrentPosition()
}

func ChemicalSpillCheckCurrentPosition() bool {
	bot.DriveMotorRight.ResetPosition()
	bot.DriveMotorLeft.RunToRelativePositionAndHold(CHEMICAL_SPILL_SEARCH_CHECK_POSITION, CHEMICAL_SPILL_VERIFY_SPEED)
	bot.DriveMotorRight.RunToRelativePositionAndHold(CHEMICAL_SPILL_SEARCH_CHECK_POSITION, CHEMICAL_SPILL_VERIFY_SPEED)
	for !helper.IsDriveStopped() {
		if bot.UltrasonicSensor.Distance() < CHEMICAL_SPILL_SEARCH_CHECK_DISTANCE_THRESHOLD {
			bot.DriveMotorLeft.Hold()
			bot.DriveMotorRight.Hold()

			ChemicalSpillRescueCan()
			return true
		}

		bot.CycleDelay()
	}

	helper.CloseClaw()
	for !helper.IsClawClosed() { bot.CycleDelay() }

	if bot.UltrasonicSensor.Distance() < CHEMICAL_SPILL_SEARCH_CHECK_DISTANCE_THRESHOLD {
		bot.DriveMotorLeft.Hold()
		bot.DriveMotorRight.Hold()

		ChemicalSpillRescueCan()
		return true
	}

	helper.OpenClaw()
	for !helper.IsClawOpen() { bot.CycleDelay() }

	bot.DriveMotorLeft.RunToRelativePositionAndHold(-bot.DriveMotorRight.Position(), CHEMICAL_SPILL_VERIFY_SPEED)
	bot.DriveMotorRight.RunToRelativePositionAndHold(-bot.DriveMotorRight.Position(), CHEMICAL_SPILL_VERIFY_SPEED)
	for !helper.IsDriveStopped() { bot.CycleDelay() }

	return false
}

func ChemicalSpillRescueCan() {
	helper.CloseClaw()
	for !helper.IsClawClosed() { bot.CycleDelay() }

	bot.DriveMotorLeft.RunToRelativePositionAndHold(-bot.DriveMotorRight.Position(), CHEMICAL_SPILL_VERIFY_SPEED)
	bot.DriveMotorRight.RunToRelativePositionAndHold(-bot.DriveMotorRight.Position(), CHEMICAL_SPILL_VERIFY_SPEED)
	for !helper.IsDriveStopped() { bot.CycleDelay() }

	bot.DriveMotorLeft.RunToRelativePositionAndHold(-bot.DriveMotorLeft.Position(), CHEMICAL_SPILL_SEARCH_SPEED * 2)
	bot.DriveMotorRight.RunToRelativePositionAndHold(bot.DriveMotorLeft.Position(), CHEMICAL_SPILL_SEARCH_SPEED * 2)
	for !helper.IsDriveStopped() { bot.CycleDelay() }

	helper.RaiseClaw()
	for !helper.IsClawRaised() { bot.CycleDelay() }

	bot.DriveMotorLeft.RunToRelativePositionAndHold(CHEMICAL_SPILL_RESCUE_BLOCK_POSITION, CHEMICAL_SPILL_RESCUE_SPEED)
	bot.DriveMotorRight.RunToRelativePositionAndHold(CHEMICAL_SPILL_RESCUE_BLOCK_POSITION, CHEMICAL_SPILL_RESCUE_SPEED)
	for !helper.IsDriveStopped() { bot.CycleDelay() }

	helper.ReleaseClaw()
	for !helper.IsClawOpen() { bot.CycleDelay() }

	helper.OpenClaw()
	for !helper.IsClawOpen() { bot.CycleDelay() }

	bot.DriveMotorLeft.RunToRelativePositionAndHold(-(CHEMICAL_SPILL_ENTER_POSITION + CHEMICAL_SPILL_RESCUE_BLOCK_POSITION + CHEMICAL_SPILL_RESCUE_EXIT_OFFSET), CHEMICAL_SPILL_RESCUE_SPEED)
	bot.DriveMotorRight.RunToRelativePositionAndHold(-(CHEMICAL_SPILL_ENTER_POSITION + CHEMICAL_SPILL_RESCUE_BLOCK_POSITION + CHEMICAL_SPILL_RESCUE_EXIT_OFFSET), CHEMICAL_SPILL_RESCUE_SPEED)
	for !helper.IsDriveStopped() {
		if bot.ColorSensorLeft.Intensity() > CHEMICAL_SPILL_VERIFY_SILVER_INTENSITY && bot.ColorSensorRight.Intensity() > CHEMICAL_SPILL_VERIFY_SILVER_INTENSITY {
			bot.DriveMotorLeft.Hold()
			bot.DriveMotorRight.Hold()
		}

		bot.CycleDelay()
	}
}