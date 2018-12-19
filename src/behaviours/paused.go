package main

var pause = Behaviour{
	setup: func() {
		logger.trace("setting up pause behaviour")
		machine.add(State{name: "pause", transitions: []string{"follow_line"}})
	},
}