package behaviour

import (
	"github.com/liam-b/robocup-2019/bot"
	// "github.com/liam-b/robocup-2019/helper"
	"github.com/liam-b/robocup-2019/state_machine"
  // "github.com/liam-b/robocup-2019/logger"
)


var waterTower = Behaviour{
  Setup: func() {
    state_machine.Add(state_machine.State{
      Name: "water_tower.verify",
      Enter: func() {
        bot.DriveMotorLeft.RunToRelativePositionAndCoast(200, 100)
        bot.DriveMotorRight.RunToRelativePositionAndCoast(200, 100)
      },
      Update: func() {
        // if (bot.UltrasonicSensor.Distance() <= 20000) {
        //   state_machine.Transition("waterTower.skrrrt")
        // }
        // logger.Trace(bot.UltrasonicSensor.Distance())
      },
    })

    state_machine.Add(state_machine.State{
      Name: "waterTower.skrrrt",
      Update: func() {
        // bot.DriveMotorRight.ResetPosition()
        // bot.DriveMotorRight.Run()

      },
    })

  },
}
