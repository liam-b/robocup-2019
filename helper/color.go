package helper

import (
	"github.com/liam-b/robocup-2019/bot"
	"math"
)

const (
	FOLLOW_WHITE_INTENSITY = 44
	FOLLOW_BLACK_INTENSITY = 15
	FOLLOW_EXPONENT        = 1.1

	COLOR_BLACK_INTENSITY = 15

	COLOR_BLACK = 0
	COLOR_WHITE = 1
	COLOR_GREEN = 2
)

func LeftColor() int {
	return colorValue(bot.ColorSensorLeft.RGB())
}

func MiddleColor() int {
	if bot.ColorSensorMiddle.Intensity() > 35 {
		return COLOR_WHITE
	}
	return COLOR_BLACK
}

func RightColor() int {
	return colorValue(bot.ColorSensorRight.RGB())
}

func colorValue(red int, green int, blue int) int {
	if green < COLOR_BLACK_INTENSITY {
		return COLOR_BLACK
	} else if green > red + 8 && green > blue + 8 {
		return COLOR_GREEN
	}

	return COLOR_WHITE
}

func LineError() float64 {
	err := LeftError() - RightError()

	if err >= 0 {
		err = math.Pow(err, FOLLOW_EXPONENT)
	} else {
		err = -math.Pow(-err, FOLLOW_EXPONENT)
	}

	return minf(maxf(-1.0, err), 1.0)
}

func LeftError() float64 {
	_, green, _ := bot.ColorSensorLeft.RGB();
	return NormalisedSensor(green)
}

func RightError() float64 {
	_, green, _ := bot.ColorSensorRight.RGB();
	return NormalisedSensor(green)
}

func MiddleError() float64 {
	return NormalisedSensor(bot.ColorSensorMiddle.Intensity())
}

func NormalisedSensor(value int) float64 {
	raw := ScaledSensor(value)

	normalised := math.Acos(1.0 - (2 * raw)) / 3
	return minf(maxf(0.0, normalised), 1.0)
}

func ScaledSensor(value int) float64 {
	raw := float64(value-FOLLOW_BLACK_INTENSITY) / float64(FOLLOW_WHITE_INTENSITY-FOLLOW_BLACK_INTENSITY)
	return minf(maxf(0.0, raw), 1.0)
}
