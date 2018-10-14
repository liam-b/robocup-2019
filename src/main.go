package main

import "fmt"

func main() {
  machine := StateMachine{}.new()
  logger := Logger{machine: &machine}.new()

  machine.add("init")
  machine.add("setup")
  machine.add("setup")

  machine.include(loadStateMachineTemplateJSON("test.json"))

  // machine.link("init").to("setup")
  // machine.link("setup").to("init")

  machine.transition("setup")
  fmt.Println(machine.can("init"))
  fmt.Println(machine.state)
  fmt.Println(machine.is("setup.*"))

  logger.log("hi")
  logger.info("hi")
  logger.warn("hi")
  logger.error("hi")

  // fmt.Println(loadJSON("test.json")["hey"] == nil)

  // fmt.Println(machine.can("stop"))
}
