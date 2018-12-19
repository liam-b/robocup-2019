package main

import "time"

var program Program
var machine StateMachine
var logger Logger

func main() {
  program = Program{start: start, loop: loop, exit: exit}.new()
  logger = Logger{level: TraceLevel}.new()
  machine = StateMachine{}.new()

  program.init()
}

func start() {
  logger.info("setting up states")
  machine.add(State{name: "initial", transitions: []string{"follow_line"}})
  setupBehaviourStates()

  logger.info("running start")
  machine.transition("follow_line")
  setupInterrupt()
}

func loop() {
  machine.update()
  time.Sleep(time.Second)
}

func exit() {
  logger.info("program exited")
}