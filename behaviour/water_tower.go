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
        if (bot.UltrasonicSensor.Distance() <= 2800) {
          // helper.CloseClaw()
          state_machine.Transition("water_tower.backItUp")
        } 
        logger.Trace(bot.UltrasonicSensor.Distance())
      },
    })

    state_machine.Add(state_machine.State{
      Name: "water_tower.backItUp",
      Enter: func() {
        bot.DriveMotorLeft.RunToRelativePositionAndHold(-180, 300)
        bot.DriveMotorRight.RunToRelativePositionAndHold(-180, 300)
        helper.CloseClaw()
      },
      Update: func() {
        if (helper.IsDriveStopped()) {
          state_machine.Transition("water_tower.adjust")
        }
      },
    })

    state_machine.Add(state_machine.State{
      Name: "water_tower.adjust",
      Enter: func() {
        bot.DriveMotorRight.RunToRelativePositionAndHold(-180, 260) // 200, 300 & -200, 300
        bot.DriveMotorLeft.RunToRelativePositionAndHold(180, 260)        
      },
      Update: func() {
        if (helper.IsDriveStopped()) {
          state_machine.Transition("water_tower.skrrrt")
        }
      },
    })

    state_machine.Add(state_machine.State{
      Name: "water_tower.skrrrt",
      Update: func() {
        if (bot.ColorSensorMiddle.Intensity() >= 10) { // the problem is that it is starting on black, changed to middle sensor try again
          bot.DriveMotorLeft.Run(200)
          bot.DriveMotorRight.Run(380)
        } else {
          bot.DriveMotorLeft.Brake()
          bot.DriveMotorRight.Brake()
          state_machine.Transition("water_tower.captureLine")
        }
      },
    })

    state_machine.Add(state_machine.State{
      Name: "water_tower.captureLine",
      Enter: func() {
        helper.OpenClaw()
        bot.DriveMotorLeft.RunToRelativePositionAndHold(130, 150)
      },
      Update: func() {
        if (helper.IsDriveStopped()) {
          state_machine.Transition("follow_line.follow")
        }
      },
    })
  },
}

