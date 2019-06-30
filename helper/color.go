package helper

import (
	"github.com/liam-b/robocup-2019/bot"
	"math"
)

const (
	FOLLOW_WHITE_INTENSITY = 44
	FOLLOW_BLACK_INTENSITY = 15
	FOLLOW_MIDDLE_WHITE_INTENSITY = 21
	FOLLOW_MIDDLE_BLACK_INTENSITY = 5
	FOLLOW_EXPONENT = 1.1

	COLOR_BLACK_INTENSITY = 15
	COLOR_SILVER_INTENSITY = 50
	// COLOR_GREEN_INTENSITY_DIFFERENCE = 3
	COLOR_GREEN_INTENSITY_DIFFERENCE = 6
	COLOR_MIDDLE_BLACK_INTENSITY = 35

	COLOR_BLACK = 0
	COLOR_WHITE = 1
	COLOR_SILVER = 2
	COLOR_GREEN = 3
)

func LeftColor() int {
	return colorValue(bot.ColorSensorLeft.RGB())
}

func MiddleColor() int {
	if bot.ColorSensorMiddle.Intensity() < COLOR_MIDDLE_BLACK_INTENSITY {
		return COLOR_BLACK
	}
	return COLOR_WHITE
}

func RightColor() int {
	return colorValue(bot.ColorSensorRight.RGB())
}

func colorValue(red int, green int, blue int) int {
	if green < COLOR_BLACK_INTENSITY {
		return COLOR_BLACK
	} else if green > red + COLOR_GREEN_INTENSITY_DIFFERENCE && green > blue + COLOR_GREEN_INTENSITY_DIFFERENCE {
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
	return NormalisedSensor(green, FOLLOW_WHITE_INTENSITY, FOLLOW_BLACK_INTENSITY)
}

func RightError() float64 {
	_, green, _ := bot.ColorSensorRight.RGB();
	return NormalisedSensor(green, FOLLOW_WHITE_INTENSITY, FOLLOW_BLACK_INTENSITY)
}

func MiddleError() float64 {
	return NormalisedSensor(bot.ColorSensorMiddle.Intensity(), FOLLOW_MIDDLE_WHITE_INTENSITY, FOLLOW_MIDDLE_BLACK_INTENSITY)
}

func NormalisedSensor(value int, whiteIntensity int, blackIntensity int) float64 {
	raw := ScaledSensor(value, whiteIntensity, blackIntensity)

	normalised := math.Acos(1.0 - (2.0 * raw)) / 3.0
	return minf(maxf(0.0, normalised), 1.0)
}

func ScaledSensor(value int, whiteIntensity int, blackIntensity int) float64 {
	raw := float64(value - blackIntensity) / float64(whiteIntensity - blackIntensity)
	return minf(maxf(0.0, raw), 1.0)
}

