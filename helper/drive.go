package helper

import (
	"github.com/liam-b/robocup-2019/bot"
)

func RunDrive(speed int) {
	bot.DriveMotorLeft.Run(speed)
	bot.DriveMotorRight.Run(speed)
}

func RunToPositionDrive(position int, speed int) {
	bot.DriveMotorLeft.RunToAbsolutePositionAndBrake(position, speed)
	bot.DriveMotorRight.RunToAbsolutePositionAndBrake(position, speed)
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

func IsDriveStopped() bool {
	return bot.DriveMotorLeft.IsStopped() && bot.DriveMotorRight.IsStopped()
}