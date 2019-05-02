package helper

import (
	"github.com/liam-b/robocup-2019/bot"
)

const (
)

var drive = Helper{
	Setup: func() {
		bot.DriveMotorLeft.Coast()
		bot.DriveMotorRight.Coast()
	},

	Cleanup: func() {
		bot.DriveMotorLeft.Coast()
		bot.DriveMotorRight.Coast()
	},
}

func RunDrive(speed int) {
	bot.DriveMotorLeft.Run(speed)
	bot.DriveMotorRight.Run(speed)
}

func RunToPositionDrive(position int, speed int) {
	bot.DriveMotorLeft.RunToPositionAndBrake(position, speed)
	bot.DriveMotorRight.RunToPositionAndBrake(position, speed)
}

func StopDrive(speed int) {
	bot.DriveMotorLeft.Brake()
	bot.DriveMotorRight.Brake()
}

func TurnTankDrive(left int, right int) {
	bot.DriveMotorLeft.Run(left)
	bot.DriveMotorRight.Run(right)
}

func TurnRatioDrive(ratio float64, speed int) {
	bot.DriveMotorLeft.Run(int(ratio * float64(speed)))
	bot.DriveMotorRight.Run(int(1.0 / ratio * float64(speed)))
}