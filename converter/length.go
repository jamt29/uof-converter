package converter

import "fmt"

// LengthUnit de medida suportada
type LengthUnit string

// Unidades de longitud soportadas.
const (
	Meter     LengthUnit = "m"
	Kilometer LengthUnit = "km"
	Mile      LengthUnit = "mi"
	Foot      LengthUnit = "ft"
)

var lengthToMeters = map[LengthUnit]float64{
	Meter:     1,
	Kilometer: 1000,
	Mile:      1609.344,
	Foot:      0.3048,
}

// ConvertLength convierte value desde la unidad from hacia la unidad to.
func ConvertLength(value float64, from, to LengthUnit) (float64, error) {
	fromFactor, ok := lengthToMeters[from]
	if !ok {
		return 0, fmt.Errorf("unidad de longitud no soportada: %s", from)
	}
	toFactor, ok := lengthToMeters[to]
	if !ok {
		return 0, fmt.Errorf("unidad de longitud no soportada: %s", to)
	}
	meters := value * fromFactor
	return meters / toFactor, nil
}
