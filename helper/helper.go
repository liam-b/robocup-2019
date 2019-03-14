package helper

import "github.com/liam-b/robocup-2019/logger"

var helpers = []Helper{
	claw,
	drive,
}

type Helper struct {
	Setup func()
	Cleanup func()
}

func Setup() {
	logger.Trace("running setup for helpers")

	for _, helper := range helpers {
		if helper.Setup != nil {
			helper.Setup()
		}
	}
}

func Cleanup() {
	logger.Trace("running cleanup for helpers")

	for _, helper := range helpers {
		if helper.Cleanup != nil {
			helper.Cleanup()
		}
	}
}