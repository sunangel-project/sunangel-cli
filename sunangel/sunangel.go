package sunangel

import (
	"math"
	"time"

	"github.com/cloudsftp/Sunangel/angle"
)

func AltitudeSunAngleAt(date time.Time, latitude, longitude float64) float64 {
	delta := declinationOfSunAt(date)
	tau := hourAngleOfSunAt(date, longitude)

	argument := math.Cos(delta) * math.Cos(tau) * math.Cos(latitude)
	argument += math.Sin(delta) * math.Sin(latitude)

	result := math.Asin(argument)
	return angle.NormalizeRadians(result)
}

func AzimutSunAngleAt(date time.Time, latitude, longitude float64) float64 {
	delta := declinationOfSunAt(date)
	tau := hourAngleOfSunAt(date, longitude)

	nominator := math.Sin(tau)
	denominator := math.Cos(tau) * math.Sin(latitude)
	denominator -= math.Tan(delta) * math.Cos(latitude)

	argument := nominator / denominator
	result := math.Atan(argument)
	if denominator < 0 {
		result += math.Pi
	}
	return angle.NormalizeRadians(result)
}
