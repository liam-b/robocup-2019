package main

type Program struct {
	looping bool
	start func()
	loop func()
	exit func()
}

func (program Program) new() Program {
	program.looping = true
	return program
}

func (program *Program) init() {
	program.start()
	for program.looping {
		program.loop()
	}
	program.exit()
}

func (program *Program) stop() {
	program.looping = false
}