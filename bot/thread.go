package bot

import (
	"time"
)

const NANO_SECOND float64 = 1000000000.0

type Thread struct {
	Target float64
	frequency float64
	running bool
	lastTime int64
	delta float64
	Cycles int64
	Cycle func(float64, int64)
}

func (thread Thread) New() Thread {
	return thread
}

func (thread Thread) Start() {
	go thread.Run()
}

func (thread Thread) Run() {
	thread.running = true
	thread.lastTime = time.Now().UnixNano()

	for thread.running {
		now := time.Now().UnixNano()
		thread.delta += (float64)(now - thread.lastTime) / (NANO_SECOND / thread.Target)
		thread.frequency = 1.0 / thread.delta * thread.Target
		thread.lastTime = now
		if thread.delta >= 1.0 {
			thread.Cycle(thread.frequency, thread.Cycles)
			thread.Cycles += 1
			thread.delta = 0.0
		}
	}
}

func (thread Thread) Stop() {
	thread.running = false
}