package converter

import "fmt"

type AreaUnit string

const (
	SquareMeter AreaUnit = "m2"
	Hectare     AreaUnit = "ha"
	Acre        AreaUnit = "ac"
	SquareFoot  AreaUnit = "ft2"
)

var areaToSquareMeters = map[AreaUnit]float64{
	SquareMeter: 1,
	Hectare:     10000,
	Acre:        4046.8564224,
	SquareFoot:  0.09290304,
}

func ConvertArea(value float64, from, to AreaUnit) (float64, error) {
	fromFactor, ok := areaToSquareMeters[from]
	if !ok {
		return 0, fmt.Errorf("Unidad de area no soportada: %s", from)
	}
	toFactor, ok := areaToSquareMeters[to]
	if !ok {
		return 0, fmt.Errorf("Unidad de area no soportada: %s", to)
	}
	squareMeters := value * fromFactor
	return squareMeters / toFactor, nil
}
