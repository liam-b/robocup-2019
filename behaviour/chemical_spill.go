package behaviour

import (
	"time"
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
	CHEMICAL_SPILL_VERIFY_ALIGN_ATTEMPTS = 1 // 4

	CHEMICAL_SPILL_ENTER_SPEED = 250
	CHEMICAL_SPILL_ENTER_POSITION = 565

	CHEMICAL_SPILL_SEARCH_SPEED = 65
	CHEMICAL_SPILL_SEARCH_ENABLE_POSITION = 200
	CHEMICAL_SPILL_SEARCH_DISTANCE_FOUND_THRESHOLD = 4800
	CHEMICAL_SPILL_SEARCH_DISTANCE_LOST_THRESHOLD = 5400
	CHEMICAL_SPILL_SEARCH_FOUND_COUNT_THRESHOLD = bot.Time(100)
	CHEMICAL_SPILL_SEARCH_LOST_COUNT_THRESHOLD = bot.Time(200)
	CHEMICAL_SPILL_SEARCH_FIRST_DIAGONAL_POSITION = 150
	CHEMICAL_SPILL_SEARCH_LAST_DIAGONAL_POSITION = 810

	CHEMICAL_SPILL_CHECK_SPEED = 150
	CHEMICAL_SPILL_CHECK_POSITION = 260
	CHEMICAL_SPILL_CHECK_DISTANCE_THRESHOLD = 2400
	CHEMICAL_SPILL_CHECK_TIME_LIMIT = bot.Time(3000)
	CHEMICAL_SPILL_CHECK_FOUND_COUNT_THRESHOLD = bot.Time(200)

	CHEMICAL_SPILL_RESCUE_SPEED = 100
	CHEMICAL_SPILL_RESCUE_BLOCK_POSITION = 295
	CHEMICAL_SPILL_RESCUE_EXIT_OFFSET = 80
	CHEMICAL_SPILL_RESCUE_EXIT_SILVER_INTENSITY = 37
	CHEMICAL_SPILL_RESCUE_EXIT_SPIN_POSITION = 500
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

	if ChemicalSpillCanInGrab(false) {
		// ChemicalSpillRescueCan()
		logger.Print("closing claw")
		helper.CloseClaw()
		for !helper.IsClawClosed() { bot.CycleDelay() }

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

		ChemicalSpillEscape()
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

	doLoop := true
	foundCount := 0
	for !helper.IsDriveStopped() || doLoop {
		logger.Print(bot.UltrasonicSensor.Distance())

		if bot.DriveMotorLeft.Position() < -CHEMICAL_SPILL_SEARCH_LAST_DIAGONAL_POSITION {
			bot.DriveMotorLeft.Hold()
			bot.DriveMotorRight.Hold()
			doLoop = false
		}

		if bot.DriveMotorLeft.Position() <= -CHEMICAL_SPILL_SEARCH_ENABLE_POSITION {
			if bot.UltrasonicSensor.Distance() <= CHEMICAL_SPILL_SEARCH_DISTANCE_FOUND_THRESHOLD {
				logger.Print("inc")
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

	lostCount := 0
	logger.Print("overshooting")
	for !helper.IsDriveStopped() {
		// logger.Print(bot.UltrasonicSensor.Distance())
		if bot.UltrasonicSensor.Distance() >= CHEMICAL_SPILL_SEARCH_DISTANCE_LOST_THRESHOLD {
			lostCount++
		} else {
			lostCount /= 2
		}

		if lostCount > CHEMICAL_SPILL_SEARCH_LOST_COUNT_THRESHOLD || bot.DriveMotorLeft.Position() < -CHEMICAL_SPILL_SEARCH_LAST_DIAGONAL_POSITION {
			logger.Print("lost can")
			bot.DriveMotorLeft.Hold()
			bot.DriveMotorRight.Hold()
		}

		bot.CycleDelay()
	}
	time.Sleep(time.Second)

	logger.Print("turning back")
	alignPosition := foundPosition + int(float64(bot.DriveMotorLeft.Position() - foundPosition) * (1.0 / 3.0))
	bot.DriveMotorLeft.Run(CHEMICAL_SPILL_SEARCH_SPEED)
	bot.DriveMotorRight.Run(-CHEMICAL_SPILL_SEARCH_SPEED)
	for !helper.IsDriveStopped() {
		logger.Print(bot.DriveMotorLeft.Position(), alignPosition)
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
	for !helper.IsDriveStopped() { bot.CycleDelay() }

	if ChemicalSpillCanInGrab(false) {
		ChemicalSpillRescueCan()
		return true
	}

	logger.Print("returning to middle")
	bot.DriveMotorLeft.RunToRelativePositionAndHold(-bot.DriveMotorRight.Position(), CHEMICAL_SPILL_CHECK_SPEED)
	bot.DriveMotorRight.RunToRelativePositionAndHold(-bot.DriveMotorRight.Position(), CHEMICAL_SPILL_CHECK_SPEED)
	for !helper.IsDriveStopped() { bot.CycleDelay() }

	return false
}

func ChemicalSpillCanInGrab(doRescue bool) bool {
	logger.Print("test grab")
	helper.CloseClaw()
	for !helper.IsClawClosed() { bot.CycleDelay() }

	foundCount := 0
	for i := 0; i < CHEMICAL_SPILL_CHECK_TIME_LIMIT; i++ {
		logger.Print(bot.UltrasonicSensor.Distance())
		if bot.UltrasonicSensor.Distance() < CHEMICAL_SPILL_CHECK_DISTANCE_THRESHOLD {
			foundCount++
			if foundCount > CHEMICAL_SPILL_CHECK_FOUND_COUNT_THRESHOLD {
				logger.Print("found can in grab")
				if doRescue { ChemicalSpillRescueCan() }
				return true
			}
		} else {
			foundCount /= 2
		}

		bot.CycleDelay()
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
	bot.DriveMotorLeft.RunToRelativePositionAndHold(-bot.DriveMotorRight.Position(), CHEMICAL_SPILL_VERIFY_SPEED * 2)
	bot.DriveMotorRight.RunToRelativePositionAndHold(-bot.DriveMotorRight.Position(), CHEMICAL_SPILL_VERIFY_SPEED * 2)
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

	ChemicalSpillEscape()
}

func ChemicalSpillEscape() {
	logger.Print("escaping")
	helper.LowerClaw()
	bot.DriveMotorLeft.RunToRelativePositionAndHold(-(CHEMICAL_SPILL_ENTER_POSITION + CHEMICAL_SPILL_RESCUE_BLOCK_POSITION + CHEMICAL_SPILL_RESCUE_EXIT_OFFSET), CHEMICAL_SPILL_RESCUE_SPEED * 2)
	bot.DriveMotorRight.RunToRelativePositionAndHold(-(CHEMICAL_SPILL_ENTER_POSITION + CHEMICAL_SPILL_RESCUE_BLOCK_POSITION + CHEMICAL_SPILL_RESCUE_EXIT_OFFSET), CHEMICAL_SPILL_RESCUE_SPEED * 2)
	for !helper.IsDriveStopped() && (bot.ColorSensorLeft.Intensity() <= CHEMICAL_SPILL_RESCUE_EXIT_SILVER_INTENSITY && bot.ColorSensorRight.Intensity() <= CHEMICAL_SPILL_RESCUE_EXIT_SILVER_INTENSITY) {
		if bot.ColorSensorLeft.Intensity() > CHEMICAL_SPILL_VERIFY_SILVER_INTENSITY && bot.ColorSensorRight.Intensity() > CHEMICAL_SPILL_VERIFY_SILVER_INTENSITY {
			bot.DriveMotorLeft.Hold()
			bot.DriveMotorRight.Hold()
		}

		bot.CycleDelay()
	}

	bot.DriveMotorLeft.RunToRelativePositionAndHold(CHEMICAL_SPILL_RESCUE_EXIT_SPIN_POSITION, CHEMICAL_SPILL_RESCUE_SPEED * 2)
	bot.DriveMotorRight.RunToRelativePositionAndHold(-CHEMICAL_SPILL_RESCUE_EXIT_SPIN_POSITION, CHEMICAL_SPILL_RESCUE_SPEED * 2)
	for !helper.IsDriveStopped() { bot.CycleDelay() }
}