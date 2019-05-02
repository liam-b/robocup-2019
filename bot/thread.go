package bot

import (
	// "github.com/liam-b/robocup-2019/logger"

	"time"
	// "strconv"
)

const (
	THREAD_CYCLE_DROP_THRESHOLD = 0.5
	NANO_SECOND = 1000000000.0
)

type Thread struct {
	Target float64
	Frequency float64
	Cycle func()
	Cycles int64
	LastCycleTime int64
	
	alive bool
	running bool
	delta float64
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
	thread.LastCycleTime = time.Now().UnixNano()

	for thread.alive {
		if thread.running {
			now := time.Now().UnixNano()
			thread.delta += float64(now - thread.LastCycleTime) / (NANO_SECOND / thread.Target)
			thread.Frequency = 1.0 / thread.delta * thread.Target
			thread.LastCycleTime = now

			if thread.delta >= 1.0 {
				thread.doCycle()
			}
			thread.checkFrameDrop()
		}
	}
}

func (thread *Thread) doCycle() {
	thread.Cycle()
	thread.Cycles += 1
	thread.delta = 0.0
}

func (thread Thread) checkFrameDrop() {
	// dropped := 100 - int(1.0 / thread.delta * 100.0)
	// if dropped > int(THREAD_CYCLE_DROP_THRESHOLD * 100.0) {
	// 	logger.Warn("thread dropping " + strconv.Itoa(dropped) + "% of cycles")
	// }
}

func (thread *Thread) Stop() {
	thread.running = false
}

func (thread *Thread) Destroy() {
	thread.alive = false
}