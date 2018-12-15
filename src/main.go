package main

import "fmt"

var machine StateMachine
var thing int

func main() {
  machine = StateMachine{}.new()
  // logger := Logger{machine: &machine}.new()

  machine.add("init")
  machine.add("loop")

  machine.link("init").to("loop")
  machine.link("loop").to("init")

  // machine.transition("setup")
  // fmt.Println(machine.can("init"))
  // fmt.Println(machine.state)
  // fmt.Println(machine.is("setup.*"))
  //
  // logger.log("hi")
  // logger.info("hi")
  // logger.warn("hi")
  // logger.error("hi")

  thing = 0

  machine.transition("loop")

  loop()
}

func loop() {
  if machine.before("loop") {
    fmt.Println("machine.before(\"loop\")")
    thing = 0
  }
  if machine.is("loop") {
    fmt.Println("machine.is(\"loop\")")
    thing += 1
    if thing >= 10 {
      fmt.Println("machine.transition(\"init\")")
      machine.transition("init")
    }
  }
  if machine.after("loop") {
    fmt.Println("machine.after(\"loop\")")
  }

  // loop()
}

// thing = 0
// machine.transition("state")
//
// if machine.is("state") {
//   thing += 1
//   if thing > 10 {
//     doSomething()
//     machine.transition("other")
//   }
// }