package behaviour

import (
	"github.com/liam-b/robocup-2019/bot"
	"github.com/liam-b/robocup-2019/helper"
	"github.com/liam-b/robocup-2019/state_machine"
  "github.com/liam-b/robocup-2019/logger"
)


var water_tower = Behaviour{
  Setup: func() {
    state_machine.Add(state_machine.State{
      Name: "water_tower.verify",
      Update: func() {
        if (bot.UltrasonicSensor.Distance() <= 4000) {
          // helper.CloseClaw()
          state_machine.Transition("water_tower.adjust")
        } 
        logger.Trace(bot.UltrasonicSensor.Distance())
      },
    })

    state_machine.Add(state_machine.State{
      Name: "water_tower.adjust",
      Enter: func() {
        helper.CloseClaw()
        bot.DriveMotorRight.RunToRelativePositionAndHold(-150, 300)
        bot.DriveMotorLeft.RunToRelativePositionAndHold(150, 300)        
      },
      Update: func() {
        if (helper.IsDriveHolding()) {
          state_machine.Transition("water_tower.skrrrt")
        }
      },
    })

    state_machine.Add(state_machine.State{
      Name: "water_tower.skrrrt",
      Update: func() {
        if (bot.ColorSensorMiddle.Intensity() >= 10) { // the problem is that it is starting on black, changed to middle sensor try again
          bot.DriveMotorLeft.Run(200)
          bot.DriveMotorRight.Run(290)
        } else {
          state_machine.Transition("follow_line.follow")
        }
      },
    })
  },
}

