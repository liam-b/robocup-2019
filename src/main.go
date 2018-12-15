package main

import "os"

var program Program
var machine StateMachine
var logger Logger

func main() {
  program = Program{setup: setup, start: start, loop: loop, exit: exit}.new()
  logger = Logger{level: DebugLevel}.new()
  machine = StateMachine{}.new()

  program.init()
}

func setup() {
  logger.info("setting up states")
}

func start() {
  logger.info("program started")
  machine.add(State{name: "initial", transitions: []string{"loop"},
    enter: func() {
      // logger.debug("entered initial");
    },
    update: func() {
      // logger.debug("updated initial");
    },
    exit: func() {
      // logger.debug("exited initial");
    },
  })

  machine.add(State{name: "loop", transitions: []string{""},
    enter: func() {
      // logger.debug("entered loop");
    },
    update: func() {
      // logger.debug("updated loop");
    },
    exit: func() {
      // logger.debug("exited loop");
    },
  })

  machine.transition("loop")
  machine.update()
}

func loop() {
  logger.trace("looping")
  program.stop()
}

func exit() {
  logger.info("program exited")
  os.Exit(1)
}