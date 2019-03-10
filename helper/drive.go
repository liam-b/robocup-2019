package helper

import (
	"github.com/liam-b/robocup-2019/bot"
)

const (
)

var drive = Helper{
	Setup: func() {
		bot.LeftDriveMotor.Coast()
		bot.RightDriveMotor.Coast()
	},

	Cleanup: func() {
		bot.LeftDriveMotor.Coast()
		bot.RightDriveMotor.Coast()
	},
}

func RunDrive(speed int) {
	bot.LeftDriveMotor.Run(speed)
	bot.RightDriveMotor.Run(speed)
}

func RunToPositionDrive(position int, speed int) {
	bot.LeftDriveMotor.RunToPositionAndBrake(position, speed)
	bot.RightDriveMotor.RunToPositionAndBrake(position, speed)
}

func StopDrive(speed int) {
	bot.LeftDriveMotor.Brake()
	bot.RightDriveMotor.Brake()
}

func TurnTankDrive(left int, right int) {
	bot.LeftDriveMotor.Run(left)
	bot.RightDriveMotor.Run(right)
}

func TurnRatioDrive(ratio float64, speed int) {
	bot.LeftDriveMotor.Run(int(ratio * float64(speed)))
	bot.RightDriveMotor.Run(int(1.0 / ratio * float64(speed)))
}