package main

var followLine = Behaviour{
	setup: func() {
		logger.trace("setting up follow_line behaviour")
		machine.add(State{name: "follow_line", transitions: []string{"pause", "green_turn.verify", "water_tower.verify"}})
	},
}