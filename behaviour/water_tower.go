package behaviour

import (
	"github.com/liam-b/robocup-2019/bot"
	"github.com/liam-b/robocup-2019/helper"
)

var (
	WATER_TOWER_SPEED = 100
	WATER_TOWER_VERIFY_DISTANCE = 3350
	WATER_TOWER_VERIFY_POSITION_LIMIT = 500

	WATER_TOWER_INNER_SPEED = 200
	WATER_TOWER_OUTER_SPEED = 435
	WATER_TOWER_SKIRT_TURN_DEGREES = 200
	WATER_TOWER_SKIRT_LINE_INTENSITY = 20
	WATER_TOWER_SKIRT_DEGREES = 1000

	WATER_TOWER_RECAPTURE_INNER_POSITION = 280
	WATER_TOWER_RECAPTURE_OUTER_POSITION = -15

	WATER_TOWER_RECAPTURE_LINE_INTENSITY = 15
	WATER_TOWER_RECAPTURE_90_TURN_POSITION = 400
)

func WaterTowerVerify() {
	bot.DriveMotorLeft.RunToRelativePositionAndHold(-WATER_TOWER_VERIFY_POSITION_LIMIT, WATER_TOWER_SPEED)
	bot.DriveMotorRight.RunToRelativePositionAndHold(-WATER_TOWER_VERIFY_POSITION_LIMIT, WATER_TOWER_SPEED)
	for bot.UltrasonicSensor.Distance() < WATER_TOWER_VERIFY_DISTANCE {
		if helper.IsDriveStopped() {
			return
		}

		bot.CycleDelay()
	}

	WaterTowerSkirt()
}

func WaterTowerSkirt() {
	bot.DriveMotorLeft.ResetPosition()
	bot.DriveMotorLeft.RunToRelativePositionAndHold(WATER_TOWER_SKIRT_TURN_DEGREES, WATER_TOWER_SPEED)
	bot.DriveMotorRight.RunToRelativePositionAndHold(-WATER_TOWER_SKIRT_TURN_DEGREES, WATER_TOWER_SPEED)
	for helper.IsDriveStopped() { bot.CycleDelay() }

	bot.DriveMotorLeft.ResetPosition()
	bot.DriveMotorLeft.Run(WATER_TOWER_INNER_SPEED)
	bot.DriveMotorRight.Run(WATER_TOWER_OUTER_SPEED)
	for bot.ColorSensorMiddle.Intensity() > WATER_TOWER_SKIRT_LINE_INTENSITY { 
		if bot.DriveMotorLeft.Position() > WATER_TOWER_SKIRT_DEGREES {
			WaterTowerOvershootRecapture()
			return
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