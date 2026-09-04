package converter

import "fmt"

// SpeedUnit representa las unidades de velocidad soportadas.
type SpeedUnit string

// Unidades de velocidad soportadas.
const (
	KmPerHour SpeedUnit = "km/h"
	Mph       SpeedUnit = "mph"
	Mps       SpeedUnit = "m/s"
)

var speedToMps = map[SpeedUnit]float64{
	KmPerHour: 1.0 / 3.6,
	Mph:       0.44704,
	Mps:       1,
}

// ConvertSpeed convierte value desde la unidad from hacia la unidad to.
func ConvertSpeed(value float64, from, to SpeedUnit) (float64, error) {
	fromFactor, ok := speedToMps[from]
	if !ok {
		return 0, fmt.Errorf("unidad de velocidad no soportada: %s", from)
	}
	toFactor, ok := speedToMps[to]
	if !ok {
		return 0, fmt.Errorf("unidad de velocidad no soportada: %s", to)
	}
	mps := value * fromFactor
	return mps / toFactor, nil
}
