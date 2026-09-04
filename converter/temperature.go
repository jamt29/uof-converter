package converter

import "fmt"

// TempUnit representa las escalas de temperatura soportadas
type TempUnit string

// Escalas de temperatura soportadas.
const (
	Celsius    TempUnit = "C"
	Fahrenheit TempUnit = "F"
	Kelvin     TempUnit = "K"
)

// ConvertTemperature convierte value desde la escala from hacia la escala to.
func ConvertTemperature(value float64, from, to TempUnit) (float64, error) {
	celsius, err := toCelsius(value, from)
	if err != nil {
		return 0, err
	}
	return fromCelsius(celsius, to)
}

func toCelsius(value float64, unit TempUnit) (float64, error) {
	switch unit {
	case Celsius:
		return value, nil
	case Fahrenheit:
		return (value - 32) * 5 / 9, nil
	case Kelvin:
		return value - 273.15, nil
	default:
		return 0, fmt.Errorf("temperatura no soportada: %s", unit)
	}
}

func fromCelsius(celsius float64, unit TempUnit) (float64, error) {
	switch unit {
	case Celsius:
		return celsius, nil
	case Fahrenheit:
		return celsius*9/5 + 32, nil
	case Kelvin:
		return celsius + 273.15, nil
	default:
		return 0, fmt.Errorf("temperatura no soportada: %s", unit)
	}
}
