package behaviour

import (
	"github.com/liam-b/robocup-2019/bot"
	"github.com/liam-b/robocup-2019/helper"
	"github.com/liam-b/robocup-2019/logger"
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
	CHEMICAL_SPILL_SEARCH_ENABLE_POSITION = 200
	CHEMICAL_SPILL_SEARCH_DISTANCE_THRESHOLD = 4600
	CHEMICAL_SPILL_SEARCH_FOUND_COUNT_THRESHOLD = bot.Time(100)
	CHEMICAL_SPILL_SEARCH_FIRST_DIAGONAL_POSITION = 150
	CHEMICAL_SPILL_SEARCH_LAST_DIAGONAL_POSITION = 810

	CHEMICAL_SPILL_CHECK_SPEED = 150
	CHEMICAL_SPILL_CHECK_POSITION = 230
	CHEMICAL_SPILL_CHECK_DISTANCE_THRESHOLD = 4000
	CHEMICAL_SPILL_CHECK_TIME_LIMIT = bot.Time(200)
	CHEMICAL_SPILL_CHECK_FOUND_COUNT_THRESHOLD = bot.Time(50)

	CHEMICAL_SPILL_RESCUE_SPEED = 100
	CHEMICAL_SPILL_RESCUE_BLOCK_POSITION = 200
	CHEMICAL_SPILL_RESCUE_EXIT_OFFSET = 80
)

func ChemicalSpillVerify() {
	logger.Print("detected chemical spill")
	bot.DriveMotorLeft.RunToRelativePositionAndHold(CHEMICAL_SPILL_VERIFY_OVERSHOOT_POSITION, CHEMICAL_SPILL_VERIFY_SPEED)
	bot.DriveMotorRight.RunToRelativePositionAndHold(CHEMICAL_SPILL_VERIFY_OVERSHOOT_POSITION, CHEMICAL_SPILL_VERIFY_SPEED)
	for !helper.IsDriveStopped() { bot.CycleDelay() }

	chemicalSpillVerifyAlignAttemptCount = 0
	ChemicalSpillBackwardAlign()
}

func ChemicalSpillBackwardAlign() {
	logger.Print("backwards silver align")
	bot.DriveMotorLeft.Run(-CHEMICAL_SPILL_VERIFY_SPEED)
	bot.DriveMotorRight.Run(-CHEMICAL_SPILL_VERIFY_SPEED)
	for !helper.IsDriveStopped() {
		if bot.ColorSensorLeft.Intensity() > CHEMICAL_SPILL_VERIFY_SILVER_INTENSITY {
			bot.DriveMotorLeft.Hold()
		}

		if bot.ColorSensorRight.Intensity() > CHEMICAL_SPILL_VERIFY_SILVER_INTENSITY {
			bot.DriveMotorRight.Hold()
		}

		bot.CycleDelay()
	}

	ChemicalSpillForwardAlign()
}

func ChemicalSpillForwardAlign() {
	logger.Print("forwards silver align")
	bot.DriveMotorLeft.Run(CHEMICAL_SPILL_VERIFY_SPEED)
	bot.DriveMotorRight.Run(CHEMICAL_SPILL_VERIFY_SPEED)
	for !helper.IsDriveStopped() {
		if bot.ColorSensorLeft.Intensity() < CHEMICAL_SPILL_VERIFY_GREEN_INTENSITY {
			bot.DriveMotorLeft.Hold()
		}

		if bot.ColorSensorRight.Intensity() < CHEMICAL_SPILL_VERIFY_GREEN_INTENSITY {
			bot.DriveMotorRight.Hold()
		}

		bot.CycleDelay()
	}

	chemicalSpillVerifyAlignAttemptCount++
	if chemicalSpillVerifyAlignAttemptCount >= CHEMICAL_SPILL_VERIFY_ALIGN_ATTEMPTS {
		ChemicalSpillEnter()
	} else {
		ChemicalSpillBackwardAlign()
	}
}

func ChemicalSpillEnter() {
	logger.Print("entering spill")

	bot.DriveMotorLeft.RunToRelativePositionAndHold(CHEMICAL_SPILL_ENTER_POSITION, CHEMICAL_SPILL_ENTER_SPEED)
	bot.DriveMotorRight.RunToRelativePositionAndHold(CHEMICAL_SPILL_ENTER_POSITION, CHEMICAL_SPILL_ENTER_SPEED)
	for !helper.IsDriveStopped() { bot.CycleDelay() }

	if ChemicalSpillCanInGrab() {
		ChemicalSpillRescueCan()
		return
	}

	logger.Print("searching for can")
	ChemicalSpillSearch()
}

func ChemicalSpillSearch() {
	logger.Print("rotating to first diagonal")
	bot.DriveMotorLeft.ResetPosition()
	bot.DriveMotorLeft.Run(-CHEMICAL_SPILL_SEARCH_SPEED)
	bot.DriveMotorRight.Run(CHEMICAL_SPILL_SEARCH_SPEED)
	for !helper.IsDriveStopped() {
		if bot.DriveMotorLeft.Position() < -CHEMICAL_SPILL_SEARCH_FIRST_DIAGONAL_POSITION {
			bot.DriveMotorLeft.Hold()
			bot.DriveMotorRight.Hold()
		}

		bot.CycleDelay()
	}

	logger.Print("checking first diagonal")
	if ChemicalSpillCheckCurrentPosition() {
		return
	}

	logger.Print("searching for can until last diagonal")
	bot.DriveMotorLeft.Run(-CHEMICAL_SPILL_SEARCH_SPEED)
	bot.DriveMotorRight.Run(CHEMICAL_SPILL_SEARCH_SPEED)

	foundCount := 0
	for !helper.IsDriveStopped() {
		if bot.DriveMotorLeft.Position() < -CHEMICAL_SPILL_SEARCH_LAST_DIAGONAL_POSITION {
			bot.DriveMotorLeft.Hold()
			bot.DriveMotorRight.Hold()
		}

		if bot.DriveMotorLeft.Position() >= CHEMICAL_SPILL_SEARCH_ENABLE_POSITION {
			if bot.UltrasonicSensor.Distance() <= CHEMICAL_SPILL_SEARCH_DISTANCE_THRESHOLD {
				foundCount++
			} else {
				foundCount /= 2
			}

			if foundCount > CHEMICAL_SPILL_SEARCH_FOUND_COUNT_THRESHOLD {
				logger.Print("found can")
				if ChemicalSpillAlignWithCan() {
					return
				}
			}
		}

		bot.CycleDelay()
	}

	ChemicalSpillCheckCurrentPosition() // should return to block and try again after this
}

func ChemicalSpillAlignWithCan() bool {
	logger.Print("aligning with can")
	foundPosition := bot.DriveMotorLeft.Position()
	bot.DriveMotorLeft.Run(-CHEMICAL_SPILL_SEARCH_SPEED)
	bot.DriveMotorRight.Run(CHEMICAL_SPILL_SEARCH_SPEED)

	lostCount := 0
	for !helper.IsDriveStopped() {
		if bot.UltrasonicSensor.Distance() >= CHEMICAL_SPILL_SEARCH_DISTANCE_THRESHOLD {
			lostCount++
		} else {
			lostCount /= 2
		}

		if lostCount > CHEMICAL_SPILL_SEARCH_FOUND_COUNT_THRESHOLD {
			logger.Print("lost can")
			bot.DriveMotorLeft.Hold()
			bot.DriveMotorRight.Hold()
		}

		bot.CycleDelay()
	}

	logger.Print("turning back")
	alignPosition := int(float64(bot.DriveMotorRight.Position() - foundPosition) * (2.0 / 3.0))
	bot.DriveMotorLeft.Run(CHEMICAL_SPILL_SEARCH_SPEED)
	bot.DriveMotorRight.Run(-CHEMICAL_SPILL_SEARCH_SPEED)
	for !helper.IsDriveStopped() {
		if bot.DriveMotorLeft.Position() > alignPosition {
			bot.DriveMotorLeft.Hold()
			bot.DriveMotorRight.Hold()
		}

		bot.CycleDelay()
	}

	return ChemicalSpillCheckCurrentPosition()
}

func ChemicalSpillCheckCurrentPosition() bool {
	logger.Print("checking current position for can")
	bot.DriveMotorRight.ResetPosition()
	bot.DriveMotorLeft.RunToRelativePositionAndHold(CHEMICAL_SPILL_CHECK_POSITION, CHEMICAL_SPILL_CHECK_SPEED)
	bot.DriveMotorRight.RunToRelativePositionAndHold(CHEMICAL_SPILL_CHECK_POSITION, CHEMICAL_SPILL_CHECK_SPEED)
	for helper.IsDriveStopped() { bot.CycleDelay() }

	if ChemicalSpillCanInGrab() {
		ChemicalSpillRescueCan()
		return true
	}

	logger.Print("returning to middle")
	bot.DriveMotorLeft.RunToRelativePositionAndHold(-bot.DriveMotorRight.Position(), CHEMICAL_SPILL_CHECK_SPEED)
	bot.DriveMotorRight.RunToRelativePositionAndHold(-bot.DriveMotorRight.Position(), CHEMICAL_SPILL_CHECK_SPEED)
	for !helper.IsDriveStopped() { bot.CycleDelay() }

	return false
}

func ChemicalSpillCanInGrab() bool {
	logger.Print("test grab")
	helper.CloseClaw()
	for !helper.IsClawClosed() { bot.CycleDelay() }

	foundCount := 0
	for i := 0; i < CHEMICAL_SPILL_CHECK_TIME_LIMIT; i++ {
		if bot.UltrasonicSensor.Distance() < CHEMICAL_SPILL_CHECK_DISTANCE_THRESHOLD {
			foundCount++
			if foundCount > CHEMICAL_SPILL_CHECK_FOUND_COUNT_THRESHOLD {
				logger.Print("found can in grab")
				ChemicalSpillRescueCan()
				return true
			}
		} else {
			foundCount /= 2
		}
	}

	logger.Print("no can in grab")
	helper.OpenClaw()
	for !helper.IsClawOpen() { bot.CycleDelay() }

	return false
}

func ChemicalSpillRescueCan() {
	logger.Print("rescuing can")
	helper.CloseClaw()
	for !helper.IsClawClosed() { bot.CycleDelay() }

	logger.Print("returning to middle")
	bot.DriveMotorLeft.RunToRelativePositionAndHold(-bot.DriveMotorRight.Position(), CHEMICAL_SPILL_VERIFY_SPEED)
	bot.DriveMotorRight.RunToRelativePositionAndHold(-bot.DriveMotorRight.Position(), CHEMICAL_SPILL_VERIFY_SPEED)
	for !helper.IsDriveStopped() { bot.CycleDelay() }

	logger.Print("rotating back to block")
	bot.DriveMotorLeft.RunToRelativePositionAndHold(-bot.DriveMotorLeft.Position(), CHEMICAL_SPILL_SEARCH_SPEED * 2)
	bot.DriveMotorRight.RunToRelativePositionAndHold(bot.DriveMotorLeft.Position(), CHEMICAL_SPILL_SEARCH_SPEED * 2)
	for !helper.IsDriveStopped() { bot.CycleDelay() }

	ChemicalSpillPlaceCanOnBlock()
}

func ChemicalSpillPlaceCanOnBlock() {
	logger.Print("raising claw")
	helper.RaiseClaw()
	for !helper.IsClawRaised() { bot.CycleDelay() }

	bot.DriveMotorLeft.RunToRelativePositionAndHold(CHEMICAL_SPILL_RESCUE_BLOCK_POSITION, CHEMICAL_SPILL_RESCUE_SPEED)
	bot.DriveMotorRight.RunToRelativePositionAndHold(CHEMICAL_SPILL_RESCUE_BLOCK_POSITION, CHEMICAL_SPILL_RESCUE_SPEED)
	for !helper.IsDriveStopped() { bot.CycleDelay() }

	logger.Print("dropping can")
	helper.ReleaseClaw()
	for !helper.IsClawOpen() { bot.CycleDelay() }

	helper.OpenClaw()
	for !helper.IsClawOpen() { bot.CycleDelay() }

	logger.Print("escaping")
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