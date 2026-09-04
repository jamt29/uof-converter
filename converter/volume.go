package converter

import "fmt"

type VolumeUnit string

const (
	Liter      VolumeUnit = "l"
	Milliliter VolumeUnit = "ml"
	Gallon     VolumeUnit = "gal"
)

var volumeToLiters = map[VolumeUnit]float64{
	Liter:      1,
	Milliliter: 0.001,
	Gallon:     3.785411784,
}

func ConvertVolume(value float64, from, to VolumeUnit) (float64, error) {
	fromFactor, ok := volumeToLiters[from]
	if !ok {
		return 0, fmt.Errorf("Unidad de volumen no soportada: %s", from)
	}
	toFactor, ok := volumeToLiters[to]
	if !ok {
		return 0, fmt.Errorf("Unidad de volumen no soportada: %s", to)
	}
	liters := value * fromFactor
	return liters / toFactor, nil
}
