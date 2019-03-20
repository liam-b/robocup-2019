package bot

import (
	"time"
)

const (
	CYCLE_DROP_THRESHOLD = 0.5
	NANO_SECOND = 1000000000.0
)

type Thread struct {
	Target float64
	frequency float64
	alive bool
	running bool
	lastTime int64
	delta float64
	Cycles int64
	Cycle func(float64, int64)
}

func (thread Thread) New() Thread {
	return thread
}

func (thread *Thread) Start() {
	go thread.Run()
}

func (thread *Thread) Run() {
	thread.alive = true
	thread.running = true
	thread.lastTime = time.Now().UnixNano()

	for thread.alive {
		if thread.running {
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
}

func (thread *Thread) Stop() {
	thread.running = false
}

func (thread *Thread) Destroy() {
	thread.alive = false
}