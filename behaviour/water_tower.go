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
      Enter: func() {
        bot.DriveMotorLeft.RunToRelativePositionAndHold(-400, 100)
        bot.DriveMotorRight.RunToRelativePositionAndHold(-400, 100)
      },
      Update: func() {
<<<<<<< HEAD
        if (bot.UltrasonicSensor.Distance() >= 6000 || helper.IsDriveStopped()) {
          state_machine.Transition("water_tower.thing")
=======
        if (bot.UltrasonicSensor.Distance() <= 3300) { //if you are using forward align change to 3500
          // helper.CloseClaw()
          state_machine.Transition("water_tower.backItUp")
>>>>>>> d2f6cee2c70cf4ec1ac92f134a943cc937272642
        } 
        logger.Trace(bot.UltrasonicSensor.Distance())
      },
    })

    state_machine.Add(state_machine.State{
      Name: "water_tower.thing",
      Enter: func() {
<<<<<<< HEAD
        bot.DriveMotorLeft.Hold()
        bot.DriveMotorRight.Hold()
        helper.CloseClaw()
      },
      Update: func() {
        if (helper.IsClawClosed()) {
=======
        bot.DriveMotorLeft.RunToRelativePositionAndHold(-180, 200)
        bot.DriveMotorRight.RunToRelativePositionAndHold(-180, 200)
        helper.CloseClaw()
      },
      Update: func() {
        if (helper.IsDriveStopped()) {
          // state_machine.Transition("water_tower.forwardAlign")
          state_machine.Transition("water_tower.adjust") //water_tower.adjust in case you want to skip forward align
        }
      },
    })

    state_machine.Add(state_machine.State{
      Name: "water_tower.forwardAlign",
      Update: func() {
        if (bot.UltrasonicSensor.Distance() > 3100) {
          bot.DriveMotorRight.RunToRelativePositionAndHold(180, 50)
          bot.DriveMotorLeft.RunToRelativePositionAndHold(180, 50)  
        } else {
>>>>>>> d2f6cee2c70cf4ec1ac92f134a943cc937272642
          state_machine.Transition("water_tower.adjust")
        } 
      },
    })

    // state_machine.Add(state_machine.State{
    //   Name: "water_tower.backItUp",
    //   Enter: func() {
    //     bot.DriveMotorLeft.RunToRelativePositionAndHold(-180, 300)
    //     bot.DriveMotorRight.RunToRelativePositionAndHold(-180, 300)
    //     helper.CloseClaw()
    //   },
    //   Update: func() {
    //     if (helper.IsDriveStopped()) {
    //       state_machine.Transition("water_tower.adjust")
    //     }
    //   },
    // })

    state_machine.Add(state_machine.State{
      Name: "water_tower.adjust",
      Enter: func() {
        bot.DriveMotorRight.RunToRelativePositionAndHold(-180, 160)
        bot.DriveMotorLeft.RunToRelativePositionAndHold(180, 160)        
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
<<<<<<< HEAD
        bot.DriveMotorLeft.RunToRelativePositionAndHold(280, 200)
        bot.DriveMotorRight.RunToRelativePositionAndHold(-15, 200)
=======
        bot.DriveMotorLeft.RunToRelativePositionAndHold(280, 150) //130, 150
        bot.DriveMotorRight.RunToRelativePositionAndHold(-15, 150)
>>>>>>> d2f6cee2c70cf4ec1ac92f134a943cc937272642
      },
      Update: func() {
        if (helper.IsDriveStopped()) {
          state_machine.Transition("follow_line.follow")
        }
      },
    })
  },
}

