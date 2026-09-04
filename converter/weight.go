package converter

import "fmt"

// WeightUnit representa las unidades de peso soportado
type WeightUnit string

// Unidades de peso soportadas.
const (
	Kilogram WeightUnit = "kg"
	Gram     WeightUnit = "g"
	Pound    WeightUnit = "lb"
)

var weightToKilograms = map[WeightUnit]float64{
	Kilogram: 1,
	Gram:     0.001,
	Pound:    0.45359237,
}

// ConvertWeight convierte value desde la unidad from hacia la unidad to.
func ConvertWeight(value float64, from, to WeightUnit) (float64, error) {
	fromFactor, ok := weightToKilograms[from]
	if !ok {
		return 0, fmt.Errorf("unidad de peso no soportada: %s", from)
	}
	toFactor, ok := weightToKilograms[to]
	if !ok {
		return 0, fmt.Errorf("unidad de peso no soportada: %s", to)
	}
	kilograms := value * fromFactor
	return kilograms / toFactor, nil
}
