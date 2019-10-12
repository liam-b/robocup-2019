package helper

const (
	// small tiles //
  	// PROPORTIONAL = 580
	// INTEGRAL     = 1.5
	// DERIVATIVE  = 1950
	// BASE_SPEED = 160

	PROPORTIONAL = 550 // 550
	INTEGRAL     = 1.46 //1.46
	DERIVATIVE   = 1950 //1950
	BASE_SPEED       = 205 // 240+ is danger zone

	HARD_TURN_VALUE  = 0.25
	HARD_TURN_SPEED  = 400
	HARD_TURN_OFFSET = 000
	RESET_THRESHOLD  = 0.3 //0.22
)

var (
	lastError = 0.0
	integral  = 0.0
)

func PID() (int, int) {
	currentError := LineError()
	integral += currentError
	derivative := currentError - lastError

	speed := (PROPORTIONAL * currentError) + (INTEGRAL * integral) + (DERIVATIVE * derivative)
	lastError = currentError

	left := min(max(-1000, BASE_SPEED + int(speed)), 1000)
	right := min(max(-1000, BASE_SPEED - int(speed)), 1000)

	// if LeftError() < HARD_TURN_VALUE {
	// 	left = HARD_TURN_OFFSET - HARD_TURN_SPEED
	// 	right = HARD_TURN_OFFSET + HARD_TURN_SPEED
	// }
  //
	// if RightError() < HARD_TURN_VALUE {
	// 	left = HARD_TURN_OFFSET + HARD_TURN_SPEED
	// 	right = HARD_TURN_OFFSET - HARD_TURN_SPEED
	// }
  //
	if MiddleError() > RESET_THRESHOLD {
		integral = 0
	}

	return left, right
}

func ResetPID() {
	lastError = 0.0
	integral = 0.0
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func cap(v int, r int) int {
	if v > r {
		return r
	}
	if v < -r {
		return -r
	}
	return v
}

func maxf(a float64, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func minf(a float64, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func capf(v float64, r float64) float64 {
	if v > r {
		return r
	}
	if v < -r {
		return -r
	}
	return v
}
