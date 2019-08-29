package behaviour

import (
	"github.com/liam-b/robocup-2019/bot"
	"github.com/liam-b/robocup-2019/helper"
)

var (
	WATER_TOWER_SPEED = 100
	WATER_TOWER_VERIFY_DISTANCE = 3870
	WATER_TOWER_VERIFY_COUNT = bot.Time(300)
	WATER_TOWER_VERIFY_LOST_COUNT = bot.Time(300)
	WATER_TOWER_VERIFY_POSITION_LIMIT = 500

	WATER_TOWER_INNER_SPEED = 205
	WATER_TOWER_OUTER_SPEED = 435
	WATER_TOWER_SKIRT_TURN_DEGREES = 200
	WATER_TOWER_SKIRT_LINE_INTENSITY = 20
	WATER_TOWER_SKIRT_FOUND_LINE_COUNT = bot.Time(50)
	WATER_TOWER_SKIRT_DEGREES = 1000

	WATER_TOWER_RECAPTURE_INNER_POSITION = 280
	WATER_TOWER_RECAPTURE_OUTER_POSITION = -15

	WATER_TOWER_RECAPTURE_LINE_INTENSITY = 15
	WATER_TOWER_RECAPTURE_90_TURN_POSITION = 400
)

func WaterTowerVerify() {
	bot.DriveMotorLeft.RunToRelativePositionAndHold(-WATER_TOWER_VERIFY_POSITION_LIMIT, WATER_TOWER_SPEED)
	bot.DriveMotorRight.RunToRelativePositionAndHold(-WATER_TOWER_VERIFY_POSITION_LIMIT, WATER_TOWER_SPEED)

	lostCount := 0
	for !helper.IsDriveStopped() {
		if bot.UltrasonicSensor.Distance() > WATER_TOWER_VERIFY_DISTANCE {
			lostCount += 1
		}

		if lostCount >= WATER_TOWER_VERIFY_LOST_COUNT {
			bot.DriveMotorLeft.Hold()
			bot.DriveMotorRight.Hold()
		}

		bot.CycleDelay()
	}

	WaterTowerSkirt()
}

func WaterTowerSkirt() {
	bot.DriveMotorLeft.ResetPosition()
	bot.DriveMotorLeft.RunToRelativePositionAndHold(WATER_TOWER_SKIRT_TURN_DEGREES, WATER_TOWER_SPEED)
	bot.DriveMotorRight.RunToRelativePositionAndHold(-WATER_TOWER_SKIRT_TURN_DEGREES, WATER_TOWER_SPEED)
	for !helper.IsDriveStopped() { bot.CycleDelay() }

	bot.DriveMotorLeft.ResetPosition()
	bot.DriveMotorLeft.Run(WATER_TOWER_INNER_SPEED)
	bot.DriveMotorRight.Run(WATER_TOWER_OUTER_SPEED)

	foundLineCount := 0
	for !helper.IsDriveStopped() { 
		if bot.DriveMotorLeft.Position() > WATER_TOWER_SKIRT_DEGREES {
			WaterTowerOvershootRecapture()
			// WaterTowerRecapture()
			return
		}

		if bot.ColorSensorMiddle.Intensity() < WATER_TOWER_SKIRT_LINE_INTENSITY {
			foundLineCount++
			if foundLineCount > WATER_TOWER_SKIRT_FOUND_LINE_COUNT {
				bot.DriveMotorLeft.Hold()
				bot.DriveMotorRight.Hold()
			}
		} else {
			foundLineCount /= 2
		}

		bot.CycleDelay()
	}

	WaterTowerRecapture()
}

func WaterTowerRecapture() {
	bot.DriveMotorLeft.RunToRelativePositionAndHold(WATER_TOWER_RECAPTURE_INNER_POSITION, WATER_TOWER_SPEED)
	bot.DriveMotorRight.RunToRelativePositionAndHold(WATER_TOWER_RECAPTURE_OUTER_POSITION, WATER_TOWER_SPEED)
	for !helper.IsDriveStopped() { bot.CycleDelay() }
}

func WaterTowerOvershootRecapture() {
	bot.DriveMotorLeft.Run(-WATER_TOWER_SPEED)
	bot.DriveMotorRight.Run(-WATER_TOWER_SPEED)
	for !helper.IsDriveStopped() {
		if bot.ColorSensorLeft.Intensity() < WATER_TOWER_RECAPTURE_LINE_INTENSITY {
			bot.DriveMotorLeft.Hold()
		}

		if bot.ColorSensorRight.Intensity() < WATER_TOWER_RECAPTURE_LINE_INTENSITY {
			bot.DriveMotorRight.Hold()
		}

		bot.CycleDelay()
	}

	bot.DriveMotorLeft.RunToRelativePositionAndHold(WATER_TOWER_RECAPTURE_90_TURN_POSITION, WATER_TOWER_SPEED)
	bot.DriveMotorRight.RunToRelativePositionAndHold(-WATER_TOWER_RECAPTURE_90_TURN_POSITION, WATER_TOWER_SPEED)
	for !helper.IsDriveStopped() { bot.CycleDelay() }
}