package helper

import (
	"github.com/liam-b/robocup-2019/bot"
	"math"
)

const (
	FOLLOW_WHITE_INTENSITY = 44
	FOLLOW_BLACK_INTENSITY = 11
	FOLLOW_EXPONENT = 1.0
)

func LineError() float64 {
	err := NormalisedSensor(bot.ColorSensorLeft.Intensity()) - NormalisedSensor(bot.ColorSensorRight.Intensity())

	if err >= 0 {
		err = math.Pow(err, FOLLOW_EXPONENT)
	} else {
		err = -math.Pow(-err, FOLLOW_EXPONENT)
	}

	return minf(maxf(-1.0, err), 1.0)
}

func NormalisedSensor(value int) float64 {
	raw := ScaledSensor(value)

	normalised := math.Acos(1.0 - (2 * raw)) / 3
	return minf(maxf(0.0, normalised), 1.0)
}

func ScaledSensor(value int) float64 {
	raw := float64(value - FOLLOW_BLACK_INTENSITY) / float64(FOLLOW_WHITE_INTENSITY - FOLLOW_BLACK_INTENSITY)
	return minf(maxf(0.0, raw), 1.0)
}